package el

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"time"
)

const (
	TAIL_CALL_OPTIMIZATION = true
	MAX_STACK_DEPTH        = 1000
)

var NameNotFoundError = func(name string) error {
	return fmt.Errorf("object not found %s", name)
}
var InterruptError = errors.New("interrupt")
var TimeoutError = errors.New("timeout")
var StackOverflowError = errors.New("stackoverflow")

type Runtime struct {
	ParseLiteral func(lit string) (Object, error)
	Stack        Stack
}

func (r *Runtime) LoadModule(ms ...Module) *Runtime {
	_, frame := r.Stack.Pop()
	for _, m := range ms {
		frame[m.Name] = m
	}
	return r
}

func (r *Runtime) LoadConstant(name string, value Object) *Runtime {
	_, frame := r.Stack.Pop()
	frame[name] = value
	return r
}

func (r *Runtime) searchOnStack(name string) (Object, error) {
	var stack Stack = r.Stack
	var frame Frame
	for stack.Depth() > 0 {
		stack, frame = stack.Pop()
		if o, ok := frame[name]; ok {
			return o, nil
		}
	}
	return nil, NameNotFoundError(name)
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
func (r *Runtime) Step(ctx context.Context, expr Expr) (Object, error) {
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
		switch expr := expr.(type) {
		case NameExpr:
			var v Object
			// load literal
			v, err := r.ParseLiteral(string(expr))
			if err == nil {
				return v, nil
			}
			// find in stack for variable
			v, err = r.searchOnStack(string(expr))
			if err != nil {
				return nil, err
			}
			return v, nil

		case LambdaExpr:
			lambda, err := r.searchOnStack(expr.Cmd)
			if err != nil {
				return nil, err
			}
			switch lambda := lambda.(type) {
			case Module:
				o, err := lambda.Exec(ctx, r, expr)
				if err != nil {
					return nil, err
				}
				return o, nil
			case Lambda:
				// 1. evaluate arguments
				args, err := r.stepMany(ctx, expr.Args...)
				if err != nil {
					return nil, err
				}
				// 2. make local frame from captured frame and arguments
				localFrame := maps.Clone(lambda.Closure)
				for i, paramName := range lambda.Params {
					localFrame[paramName] = args[i]
				}
				// 3. push local frame to stack if not tail call
				if options.tailCall {
					_, frame := r.Stack.Pop()
					maps.Copy(frame, localFrame)
				} else {
					r.Stack = r.Stack.Push(localFrame)
				}

				// 4. exec function
				v, err := r.Step(ctx, lambda.Impl)
				if err != nil {
					return nil, err
				}
				// 5. pop Closure from Stack
				if options.tailCall {
				} else {
					r.Stack, _ = r.Stack.Pop()
				}
				return v, nil
			default:
				return nil, fmt.Errorf("expression cannot be executed: %s", expr.String())
			}
		default:
			return nil, fmt.Errorf("unknown expression type")
		}
	}
}

func (r *Runtime) stepMany(ctx context.Context, exprList ...Expr) ([]Object, error) {
	outputs := make([]Object, len(exprList))
	for i, expr := range exprList {
		if i == len(exprList)-1 && TAIL_CALL_OPTIMIZATION {
			ctx = setOptionsToContext(ctx, &stepOptions{
				tailCall: true,
			})
		}
		value, err := r.Step(ctx, expr)
		if err != nil {
			return nil, err
		}
		outputs[i] = value
	}
	return outputs, nil
}
