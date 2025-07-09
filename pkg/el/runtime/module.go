package runtime

import (
	"context"
	"el/pkg/el/expr"
	"errors"
	"fmt"
	"maps"
)

var InternalError = errors.New("ordmap")

var letModule = Module{
	Name: "let",
	Exec: func(ctx context.Context, r *Runtime, e expr.Lambda) (Object, error) {
		if e.Cmd.(expr.Name) != "let" {
			return nil, InternalError
		}
		if len(e.Args) < 1 {
			return nil, fmt.Errorf("let requires at least 1 arguments")
		}
		if getTailCall(ctx) {
		} else {
			r.Stack.Push(Frame{})
			defer r.Stack.Pop()
		}

		for i := 0; i < len(e.Args)-1; i += 2 {
			lvalue, ok := e.Args[i].(expr.Name)
			if !ok {
				return nil, fmt.Errorf("lvalue must be a name")
			}
			rvalue, err := r.Step(ctx, e.Args[i+1])
			if err != nil {
				return nil, err
			}
			name := Name(lvalue)
			// update stack
			head := r.Stack.Pop()
			head[name] = rvalue
			r.Stack.Push(head)
		}
		value, err := r.Step(ctx, e.Args[len(e.Args)-1])
		if err != nil {
			return nil, err
		}
		return value, nil
	},
	Man: "module: (let x 3) - assign value 3 to local variable x",
}

var lambdaModule = Module{
	Name: "lambda",
	Exec: func(ctx context.Context, r *Runtime, e expr.Lambda) (Object, error) {
		if e.Cmd.(expr.Name) != "lambda" {
			return nil, InternalError
		}
		if len(e.Args) < 1 {
			return nil, fmt.Errorf("lambda requires at least 1 arguments")
		}
		paramNameList := make([]Name, 0, len(e.Args)-1)
		for i := 0; i < len(e.Args)-1; i++ {
			lvalue, ok := e.Args[i].(expr.Name)
			if !ok {
				return nil, fmt.Errorf("lvalue must be a name")
			}
			name := Name(lvalue)
			paramNameList = append(paramNameList, name)
		}
		implementation := e.Args[len(e.Args)-1]

		// capture only the top of frame stack
		head := r.Stack.Pop()
		closure := maps.Clone(head)
		r.Stack.Push(head)

		return Lambda{
			ParamNameList:  paramNameList,
			Implementation: implementation,
			Closure:        closure,
		}, nil
	},
	Man: "module: (lambda x y (add x y) - declare a function",
}

var matchModule = Module{
	Name: "match",
	Exec: func(ctx context.Context, r *Runtime, e expr.Lambda) (Object, error) {
		if e.Cmd.(expr.Name) != "match" {
			return nil, InternalError
		}
		if len(e.Args) < 2 {
			return nil, fmt.Errorf("match requires at least 2 arguments")
		}
		cond, err := r.Step(ctx, e.Args[0])
		if err != nil {
			return nil, err
		}
		i, err := func() (int, error) {
			for i := 1; i < len(e.Args); i += 2 {
				comp, err := r.Step(ctx, e.Args[i])
				if err != nil {
					return 0, err
				}
				if _, ok := comp.(Wildcard); ok || comp == cond {
					return i, nil
				}
			}
			return 0, fmt.Errorf("no case matched: %s", e.String())
		}()
		if err != nil {
			return nil, err
		}
		return r.Step(ctx, e.Args[i+1])
	},
	Man: "module: (match x 1 2 4 5) - match, if x=1 then return 3, if x=4 the return 5",
}
