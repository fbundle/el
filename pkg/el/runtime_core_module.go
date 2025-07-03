package el

import (
	"context"
	"errors"
	"fmt"
	"maps"
)

var InternalError = errors.New("internal")

var letModule = Module{
	Name: "let",
	Exec: func(ctx context.Context, r *Runtime, expr LambdaExpr) (Object, error) {
		if expr.Cmd.(NameExpr) != "let" {
			return nil, InternalError
		}
		if len(expr.Args) < 1 {
			return nil, fmt.Errorf("let requires at least 1 arguments")
		}
		r.Stack.Push(Frame{})
		defer r.Stack.Pop()

		for i := 0; i < len(expr.Args)-1; i += 2 {
			lvalue, ok := expr.Args[i].(NameExpr)
			if !ok {
				return nil, fmt.Errorf("lvalue must be a name")
			}
			rvalue, err := r.Step(ctx, expr.Args[i+1])
			if err != nil {
				return nil, err
			}
			name := Name(lvalue)
			// update stack
			head := r.Stack.Pop()
			head[name] = rvalue
			r.Stack.Push(head)
		}
		value, err := r.Step(ctx, expr.Args[len(expr.Args)-1])
		if err != nil {
			return nil, err
		}
		return value, nil
	},
	Man: "module: (let x 3) - assign value 3 to local variable x",
}

var lambdaModule = Module{
	Name: "lambda",
	Exec: func(ctx context.Context, r *Runtime, expr LambdaExpr) (Object, error) {
		if expr.Cmd.(NameExpr) != "lambda" {
			return nil, InternalError
		}
		if len(expr.Args) < 1 {
			return nil, fmt.Errorf("lambda requires at least 1 arguments")
		}
		v := Lambda{
			Params:  nil,
			Impl:    nil,
			Closure: nil,
		}
		for i := 0; i < len(expr.Args)-1; i++ {
			lvalue, ok := expr.Args[i].(NameExpr)
			if !ok {
				return nil, fmt.Errorf("lvalue must be a name")
			}
			name := Name(lvalue)
			v.Params = append(v.Params, name)
		}
		v.Impl = expr.Args[len(expr.Args)-1]
		// capture only the top of FrameStack
		head := r.Stack.Pop()
		v.Closure = maps.Clone(head)
		r.Stack.Push(head)
		return v, nil
	},
	Man: "module: (lambda x y (add x y) - declare a function",
}

var matchModule = Module{
	Name: "match",
	Exec: func(ctx context.Context, r *Runtime, expr LambdaExpr) (Object, error) {
		if expr.Cmd.(NameExpr) != "match" {
			return nil, InternalError
		}
		if len(expr.Args) < 2 {
			return nil, fmt.Errorf("match requires at least 2 arguments")
		}

		cond, err := r.Step(ctx, expr.Args[0])
		if err != nil {
			return nil, err
		}
		i, err := func() (int, error) {
			for i := 1; i < len(expr.Args); i += 2 {
				comp, err := r.Step(ctx, expr.Args[i])
				if err != nil {
					return 0, err
				}
				if _, ok := comp.(Wildcard); ok || comp == cond {
					return i, nil
				}
			}
			return 0, fmt.Errorf("no case matched: %s", expr.String())
		}()
		if err != nil {
			return nil, err
		}
		return r.Step(ctx, expr.Args[i+1])
	},
	Man: "module: (match x 1 2 4 5) - match, if x=1 then return 3, if x=4 the return 5",
}
