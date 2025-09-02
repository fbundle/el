package runtime_core

import (
	"context"
	"el/pkg/el/expr"
	"errors"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var ErrorInternal = errors.New("internal")

var letModule = Module{
	Name: "let",
	Exec: func(r Runtime, ctx context.Context, s Stack, args []expr.Expr) adt.Option[Object] {
		if len(args) < 1 {
			return errorObjectString("let requires at least 1 arguments")
		}
		if len(args)%2 != 1 {
			return errorObjectString("let requires odd number of arguments")
		}

		s = s.Push(emptyFrame) // new empty frame
		for i := 0; i < len(args)-1; i += 2 {
			// for let, rexpr must be executed in sequence
			lexpr, rexpr := args[i], args[i+1]
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

		return r.StepOpt(ctx, s, args[len(args)-1])
	},
	Man: "module: (let x 3) - assign value 3 to local variable x",
}

var lambdaModule = Module{
	Name: "lambda",
	Exec: func(r Runtime, ctx context.Context, s Stack, args []expr.Expr) adt.Option[Object] {
		if len(args) < 1 {
			return errorObjectString("lambda requires at least 1 arguments")
		}
		paramList := make([]Name, 0, len(args)-1)
		for i := 0; i < len(args)-1; i++ {
			lvalue, ok := args[i].(expr.Name)
			if !ok {
				return errorObjectString("lvalue must be a name")
			}
			paramList = append(paramList, Name(lvalue))
		}
		implementation := args[len(args)-1]
		closure := s.Peek() // capture only top of stack // TODO - capture more but only necessary variables

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
	Exec: func(r Runtime, ctx context.Context, s Stack, args []expr.Expr) adt.Option[Object] {
		if len(args) < 3 {
			return errorObjectString("match requires at least 3 arguments")
		}
		var cond Object
		if err := r.StepOpt(ctx, s, args[0]).Unwrap(&cond); err != nil {
			return errorObject(err)
		}

		i, err := func() (int, error) {
			for i := 1; i < len(args); i += 2 {
				compExpr := args[i]
				var comp Object
				if err := r.StepOpt(ctx, s, compExpr).Unwrap(&comp); err != nil {
					return 0, err
				}
				if _, ok := comp.(Wildcard); ok || comp == cond {
					return i, nil
				}
				// TODO - drop Wildcard, the last arg will be wildcard
			}
			return 0, fmt.Errorf("no case matched: %s", args)
		}()
		if err != nil {
			return errorObject(err)
		}
		return r.StepOpt(ctx, s, args[i+1])
	},
	Man: "module: (match x 1 2 4 5) - match, if x=1 then return 3, if x=4 the return 5",
}
