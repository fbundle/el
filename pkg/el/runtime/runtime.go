package runtime

import (
	"context"
	"el/pkg/el/expr"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	MAX_STACK_DEPTH = 1000
)

var NameNotFoundError = func(name Name) error {
	return fmt.Errorf("object not found %s", name)
}
var InterruptError = func(err error) error {
	return fmt.Errorf("interrupted: %s", err)
}
var TimeoutError = func(err error) error {
	return fmt.Errorf("timeout: %s", err)
}
var StackOverflowError = errors.New("stack overflow")
var LengthMismatchError = errors.New("length mismatch")

type Runtime struct {
	Stack FrameStack
}

func (r *Runtime) searchOnStack(name Name) (Object, error) {
	for _, frame := range r.Stack.Iter {
		if o, ok := frame.Get(name); ok {
			return o, nil
		}
	}
	return nil, NameNotFoundError(name)
}

// Step -
func (r *Runtime) Step(ctx context.Context, e expr.Expr) (Object, error) {
	deadline, ok := ctx.Deadline()
	if ok && time.Now().After(deadline) {
		return nil, TimeoutError(ctx.Err())
	}
	select {
	case <-ctx.Done():
		return nil, InterruptError(ctx.Err())
	default:
	}
	if r.Stack.Depth() > MAX_STACK_DEPTH {
		return nil, StackOverflowError
	}
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

		hence tailcall optimization will apply for lambda and let
	*/

	switch e := e.(type) {
	case expr.Name:
		var v Object
		// load literal
		v, err := parseLiteral(string(e))
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
		lambda, err := r.Step(ctx, e.Cmd)
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
			localFrame := lambda.Closure
			for i := 0; i < len(lambda.ParamNameList); i++ {
				paramName := lambda.ParamNameList[i]
				arg := args[i]
				localFrame = localFrame.Set(paramName, arg)
			}
			// 3. push local frame to stack if not tailcall

			r.Stack = r.Stack.Push(localFrame)
			defer func() {
				r.Stack = r.Stack.Pop()
			}() // 5. pop local frame from frame stack

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

func (r *Runtime) stepMany(ctx context.Context, es ...expr.Expr) ([]Object, error) {
	var outputs []Object
	for _, e := range es {
		out, err := r.Step(ctx, e)
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, out)
	}
	return outputs, nil
}

func (r *Runtime) LoadModule(ms ...Module) *Runtime {
	for _, m := range ms {
		r.Stack = updateHead(r.Stack, func(frame Frame) Frame {
			return frame.Set(m.Name, m)
		})
	}
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

func parseLiteral(lit string) (Object, error) {
	if len(lit) == 0 {
		return nil, errors.New("empty literal")
	}
	if lit == "_" {
		return Wildcard{}, nil
	}
	if lit == "*" {
		return Unwrap{}, nil
	}
	if lit[0] == '"' && lit[len(lit)-1] == '"' {
		str := ""
		if err := json.Unmarshal([]byte(lit), &str); err != nil {
			return nil, err
		}
		strList := List{}
		for _, ch := range []rune(str) {
			strList = append(strList, Int(ch))
		}
		return strList, nil
	}
	i, err := strconv.Atoi(lit)
	return Int(i), err
}
