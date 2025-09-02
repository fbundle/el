package runtime_core

import (
	"context"
	"el/pkg/el/expr"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var letModule = Module{
	Name: "let",
	Exec: func(r Runtime, ctx context.Context, s Stack, e expr.Lambda) adt.Option[Object] {
		if e.Cmd.(expr.Name) != "let" {
			return errorObject(ErrorInternal)
		}
		if len(e.Args) < 1 {
			return errorObjectString("let requires at least 1 arguments")
		}
		if len(e.Args)%2 != 1 {
			return errorObjectString("let requires odd number of arguments")
		}

		s = s.Push(emptyFrame) // new empty frame
		for i := 0; i < len(e.Args)-1; i += 2 {
			// for let, rexpr must be executed in sequence
			lexpr, rexpr := e.Args[i], e.Args[i+1]
			lvalue, ok := lexpr.(expr.Name)
			if !ok {
				return errorObjectString("lvalue must be a name")
			}

			var rvalue Object
			if err := r.StepOpt(ctx, s, rexpr).Unwrap(&rvalue); err != nil {
				return errorObject(err)
			}

			// update stack
			s = PeekAndUpdate(s, func(f Frame) Frame {
				return f.Set(Name(lvalue), rvalue)
			})
		}

		return r.StepOpt(ctx, s, e.Args[len(e.Args)-1])
	},
	Man: "module: (let x 3) - assign value 3 to local variable x",
}

var lambdaModule = Module{
	Name: "lambda",
	Exec: func(r Runtime, ctx context.Context, s Stack, e expr.Lambda) adt.Option[Object] {
		if e.Cmd.(expr.Name) != "lambda" {
			return errorObject(ErrorInternal)
		}
		if len(e.Args) < 1 {
			return errorObjectString("lambda requires at least 1 arguments")
		}
		paramList := make([]Name, 0, len(e.Args)-1)
		for i := 0; i < len(e.Args)-1; i++ {
			lvalue, ok := e.Args[i].(expr.Name)
			if !ok {
				return errorObjectString("lvalue must be a name")
			}
			paramList = append(paramList, Name(lvalue))
		}
		implementation := e.Args[len(e.Args)-1]
		closure := s.Peek() // capture only top of stack

		return object(Lambda{
			Params:  paramList,
			Impl:    implementation,
			Closure: closure,
		})
	},
	Man: "module: (lambda x y (add x y) - declare a function",
}

var matchModule = Module{
	Name: "match",
	Exec: func(r Runtime, ctx context.Context, s Stack, e expr.Lambda) adt.Option[Object] {
		if e.Cmd.(expr.Name) != "match" {
			return errorObject(ErrorInternal)
		}
		if len(e.Args) < 3 {
			return errorObjectString("match requires at least 3 arguments")
		}
		var cond Object
		if err := r.StepOpt(ctx, s, e.Args[0]).Unwrap(&cond); err != nil {
			return errorObject(err)
		}

		i, err := func() (int, error) {
			for i := 1; i < len(e.Args); i += 2 {
				compExpr := e.Args[i]
				var comp Object
				if err := r.StepOpt(ctx, s, compExpr).Unwrap(&comp); err != nil {
					return 0, err
				}
				if _, ok := comp.(Wildcard); ok || comp == cond {
					return i, nil
				}
			}
			return 0, fmt.Errorf("no case matched: %s", e.String())
		}()
		if err != nil {
			return errorObject(err)
		}
		return r.StepOpt(ctx, s, e.Args[i+1])
	},
	Man: "module: (match x 1 2 4 5) - match, if x=1 then return 3, if x=4 the return 5",
}
