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
		if expr.Cmd != "let" {
			return nil, InternalError
		}
		if len(expr.Args) < 1 {
			return nil, fmt.Errorf("let requires at least 1 arguments")
		}
		r.Stack = r.Stack.Push(Frame{})

		for i := 0; i < len(expr.Args)-1; i += 2 {
			lvalue, ok := expr.Args[i].(NameExpr)
			if !ok {
				return nil, fmt.Errorf("lvalue must be a name")
			}
			name := string(lvalue)
			rvalue, err := r.Step(ctx, expr.Args[i+1])
			if err != nil {
				return nil, err
			}
			// update stack
			_, frame := r.Stack.Pop()
			frame[name] = rvalue
		}
		value, err := r.Step(ctx, expr.Args[len(expr.Args)-1])
		if err != nil {
			return nil, err
		}

		r.Stack, _ = r.Stack.Pop()
		return value, nil
	},
	Man: "module: (let x 3) - assign value 3 to local variable x",
}

var lambdaModule = Module{
	Name: "lambda",
	Exec: func(ctx context.Context, r *Runtime, expr LambdaExpr) (Object, error) {
		if expr.Cmd != "lambda" {
			return nil, InternalError
		}
		if len(expr.Args) < 1 {
			return nil, fmt.Errorf("lambda requires at least 1 arguments")
		}
		v := Lambda{
			Params: nil,
			Impl:   nil,
			Frame:  nil,
		}
		for i := 0; i < len(expr.Args)-1; i++ {
			lvalue, ok := expr.Args[i].(NameExpr)
			if !ok {
				return nil, fmt.Errorf("lvalue must be a name")
			}
			paramName := string(lvalue)
			v.Params = append(v.Params, paramName)
		}
		v.Impl = expr.Args[len(expr.Args)-1]
		// capture only the top of Stack
		_, frame := r.Stack.Pop()
		v.Frame = maps.Clone(frame)
		return v, nil
	},
	Man: "module: (lambda x y (add x y) - declare a function",
}

var matchModule = Module{
	Name: "match",
	Exec: func(ctx context.Context, r *Runtime, expr LambdaExpr) (Object, error) {
		if expr.Cmd != "match" {
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

var addExtension = Extension{
	Name: "add",
	Exec: func(ctx context.Context, values ...Object) (Object, error) {
		if len(values) < 1 {
			return nil, fmt.Errorf("add requires at least 1 arguments")
		}
		var sum Int = 0
		for i := 0; i < len(values); i++ {
			v, ok := values[i].(Int)
			if !ok {
				return nil, fmt.Errorf("adding non-integer values")
			}
			sum += v
		}
		return sum, nil
	},
	Man: "module: (add 1 (add 2 3) 3) - exec a sequence of expressions and return the sum",
}

var signExtension = Extension{
	Name: "sign",
	Exec: func(ctx context.Context, values ...Object) (Object, error) {
		if len(values) < 1 {
			return nil, fmt.Errorf("sign requires at least 1 arguments")
		}
		v, ok := values[len(values)-1].(Int)
		if !ok {
			return nil, fmt.Errorf("sign non-integer value")
		}
		switch {
		case v > 0:
			return Int(+1), nil
		case v < 0:
			return Int(-1), nil
		default:
			return Int(0), nil
		}
	},
	Man: "module: (sign 3) - exec an expression and return the sign",
}

var listExtension = Extension{
	Name: "list",
	Exec: func(ctx context.Context, values ...Object) (Object, error) {
		l := List{}
		for _, v := range values {
			l = append(l, v)
		}
		return l, nil
	},
	Man: "module: (list 1 2 (lambda x (add x 1))) - make a list",
}

var getExtension = Extension{
	Name: "get",
	Exec: func(ctx context.Context, values ...Object) (Object, error) {
		if len(values) < 2 {
			return nil, fmt.Errorf("get requires at least 2 arguments")
		}
		l, ok := values[0].(List)
		if !ok {
			return nil, fmt.Errorf("get non-list value")
		}
		i, ok := values[1].(Int)
		if !ok {
			return nil, fmt.Errorf("get non-integer index")
		}
		return l[i], nil
	},
	Man: "module: (get (list 1 2 3) 1) - get the i-th element of a list",
}
