package runtime

import (
	"context"
	"el/pkg/el/expr"
	"errors"
	"fmt"
	"maps"
	"time"
)

const (
	TAIL_CALL_OPTIMIZATION = true
	MAX_STACK_DEPTH        = 1000
)

var NameNotFoundError = func(name Name) error {
	return fmt.Errorf("obj not found %s", name)
}
var InterruptError = errors.New("interrupt")
var TimeoutError = errors.New("timeout")
var StackOverflowError = errors.New("stackoverflow")

type Runtime struct {
	ParseLiteral func(lit string) (Object, error)
	Stack        FrameStack
}

func (r *Runtime) searchOnStack(name Name) (out Object, err error) {
	err = NameNotFoundError(name)
	r.Stack.Iter(func(frame Frame) bool {
		if o, ok := frame[name]; ok {
			out = o
			err = nil
			return false
		}
		return true
	})
	return out, err
}

type stepOptions struct {
	tailCall bool
}

func getOptionsFromContext(ctx context.Context) (*stepOptions, bool) {
	if o, ok := ctx.Value("step_options").(*stepOptions); ok {
		return o, true
	}
	// default option
	return &stepOptions{
		tailCall: false,
	}, false
}

func setOptionsToContext(ctx context.Context, o *stepOptions) context.Context {
	return context.WithValue(ctx, "step_options", o)
}

// Step -
func (r *Runtime) Step(ctx context.Context, e expr.Expr) (Object, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	options, _ := getOptionsFromContext(ctx)

	deadline, ok := ctx.Deadline()
	if ok && time.Now().After(deadline) {
		return nil, TimeoutError
	}
	if r.Stack.Depth() > MAX_STACK_DEPTH {
		return nil, StackOverflowError
	}
	select {
	case <-ctx.Done():
		return nil, InterruptError
	default:
		/*
			the whole language is every simple
				1. parse literal or search on stack
				2. function application: push a new frame, exec the function, pop
				3. builtin module
					a. lambda: capture the current frame and save the implementation
					b. let: push a new frame, exec the function, pop
					c. match: eval and match

			only let and function application push a new frame since
				- let requires local scope to bind new variables
				- function application requires local scope to
					bind parameters and previously captured variables in lambda
		*/

		// TODO - actually function call can be implmented as
		/*
			(let
				f (lambda x y [x + y])
				(let
					x 1
					y 2
					f		// bind local variables x y and exec implementation of f
				)
			)
		*/

		switch e := e.(type) {
		case expr.Name:
			var v Object
			// load literal
			v, err := r.ParseLiteral(string(e))
			if err == nil {
				return v, nil
			}
			// find in stack for variable
			v, err = r.searchOnStack(Name(e))
			if err != nil {
				return nil, err
			}
			return v, nil

		case expr.Lambda:
			getLambda := func(cmd expr.Expr) (Object, error) {
				switch cmd := e.Cmd.(type) {
				case expr.Name:
					return r.searchOnStack(Name(cmd))
				case expr.Lambda:
					return r.Step(ctx, cmd)
				default:
					return nil, fmt.Errorf("lambda: invalid command")
				}
			}
			lambda, err := getLambda(e.Cmd)
			if err != nil {
				return nil, err
			}
			switch lambda := lambda.(type) {
			case Module:
				o, err := lambda.Exec(ctx, r, e)
				if err != nil {
					return nil, err
				}
				return o, nil
			case Lambda:
				// 1. evaluate arguments
				args, err := r.stepMany(ctx, e.Args...)
				if err != nil {
					return nil, err
				}
				args, err = unwrapArgs(args)
				if err != nil {
					return nil, err
				}
				if len(args) != len(lambda.ParamNameList) {
					return nil, fmt.Errorf("lambda: expected %d arguments, got %d", len(lambda.ParamNameList), len(args))
				}
				// 2. make local frame from captured frame and arguments
				localFrame := maps.Clone(lambda.Closure)
				for i, paramName := range lambda.ParamNameList {
					localFrame[paramName] = args[i]
				}
				// 3. push local frame to stack if not tail call
				if options.tailCall {
					head := r.Stack.Pop()
					maps.Copy(head, localFrame)
					r.Stack.Push(head)
				} else {
					r.Stack.Push(localFrame)
				}
				defer func() {
					// 5. pop Closure from FrameStack
					if options.tailCall {
					} else {
						r.Stack.Pop()
					}
				}()

				// 4. exec function
				v, err := r.Step(ctx, lambda.Implementation)
				if err != nil {
					return nil, err
				}
				return v, nil
			default:
				return nil, fmt.Errorf("expression cannot be executed: %s", e.String())
			}
		default:
			return nil, fmt.Errorf("unknown expression type")
		}
	}
}

func (r *Runtime) stepMany(ctx context.Context, eList ...expr.Expr) ([]Object, error) {
	outputs := make([]Object, len(eList))
	for i, e := range eList {
		if i == len(eList)-1 && TAIL_CALL_OPTIMIZATION {
			ctx = setOptionsToContext(ctx, &stepOptions{
				tailCall: true,
			})
		}
		value, err := r.Step(ctx, e)
		if err != nil {
			return nil, err
		}
		outputs[i] = value
	}
	return outputs, nil
}

func (r *Runtime) LoadModule(ms ...Module) *Runtime {
	head := r.Stack.Pop()
	for _, m := range ms {
		head[m.Name] = m
	}
	r.Stack.Push(head)
	return r
}

func unwrapArgs(args []Object) ([]Object, error) {
	var unwrapArgsLoop func(args []Object) ([]Object, bool, error)
	unwrapArgsLoop = func(args []Object) ([]Object, bool, error) {
		unwrapped := false
		unwrappedArgs := make([]Object, 0, len(args))
		for len(args) > 0 {
			head := args[0]
			if _, ok := head.(Unwrap); ok {
				if len(args) <= 1 {
					return unwrappedArgs, unwrapped, errors.New("unwrapping argument empty")
				}
				switch next := args[1].(type) {
				case List:
					unwrappedArgs = append(unwrappedArgs, next...)
					args = args[2:]
					unwrapped = true
				case Unwrap: // nested unwrap
					unwrappedArgs = append(unwrappedArgs, head)
					args = args[1:]
				default:
					return unwrappedArgs, unwrapped, errors.New("unwrapping argument must be a list or an unwrap")
				}
			} else {
				unwrappedArgs = append(unwrappedArgs, head)
				args = args[1:]
			}
		}
		return unwrappedArgs, unwrapped, nil
	}
	var unwrapped bool
	var err error
	for { // keep unwrapping
		args, unwrapped, err = unwrapArgsLoop(args)
		if err != nil {
			return nil, err
		}
		if !unwrapped {
			return args, nil
		}
	}
}
