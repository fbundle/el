package runtime_ext

import (
	"context"
	"el/pkg/el/expr"
	"el/pkg/el/runtime_core"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var letModule = Module{
	name: "let",
	exec: func(r runtime_core.Runtime, ctx context.Context, s runtime_core.Stack, args []expr.Expr) adt.Option[runtime_core.Object] {
		if len(args) < 1 {
			return errorObjectString("let requires at least 1 arguments")
		}
		if len(args)%2 != 1 {
			return errorObjectString("let requires odd number of arguments")
		}

		s = s.Push(runtime_core.EmptyFrame) // new empty frame
		for i := 0; i < len(args)-1; i += 2 {
			// for let, rexpr must be executed in sequence
			lexpr, rexpr := args[i], args[i+1]
			lvalue, ok := lexpr.(expr.Name)
			if !ok {
				return errorObjectString("lvalue must be a name")
			}

			var rvalue runtime_core.Object
			if err := r.StepOpt(ctx, s, rexpr).Unwrap(&rvalue); err != nil {
				return errorObject(err)
			}

			// update stack
			s = runtime_core.PeekAndUpdate(s, func(f runtime_core.Frame) runtime_core.Frame {
				return f.Set(runtime_core.Name(lvalue), rvalue)
			})
		}

		return r.StepOpt(ctx, s, args[len(args)-1])
	},
	man: "Module: (let x 3) - assign value 3 to local variable x",
}

var lambdaModule = Module{
	name: "lambda",
	exec: func(r runtime_core.Runtime, ctx context.Context, s runtime_core.Stack, args []expr.Expr) adt.Option[runtime_core.Object] {
		if len(args) < 1 {
			return errorObjectString("lambda requires at least 1 arguments")
		}
		paramList := make([]runtime_core.Name, 0, len(args)-1)
		for i := 0; i < len(args)-1; i++ {
			lvalue, ok := args[i].(expr.Name)
			if !ok {
				return errorObjectString("lvalue must be a name")
			}
			paramList = append(paramList, runtime_core.Name(lvalue))
		}
		implementation := args[len(args)-1]
		closure := s.Peek() // capture only top of stack // TODO - capture more but only necessary variables

		return object(runtime_core.Lambda{
			Params:  paramList,
			Impl:    implementation,
			Closure: closure,
		})
	},
	man: "Module: (lambda x y (add x y) - declare a function",
}

var matchModule = Module{
	name: "match",
	exec: func(r runtime_core.Runtime, ctx context.Context, s runtime_core.Stack, argList []expr.Expr) adt.Option[runtime_core.Object] {
		if len(argList) < 3 {
			return errorObjectString("match requires at least 3 arguments")
		}
		if len(argList)%2 != 0 {
			return errorObjectString("match requires even number of arguments")
		}
		var cond runtime_core.Object
		if err := r.StepOpt(ctx, s, argList[0]).Unwrap(&cond); err != nil {
			return errorObject(err)
		}

		for i := 1; i < len(argList)-1; i += 2 {
			var comp runtime_core.Object
			if err := r.StepOpt(ctx, s, argList[i]).Unwrap(&comp); err != nil {
				return errorObject(err)
			}
			if comp == cond {
				return r.StepOpt(ctx, s, argList[i+1])
			}
		}
		return r.StepOpt(ctx, s, argList[len(argList)-1])
	},
	man: "Module: (match x 1 2 4 5 6) - match, if x=1 then return 3, if x=4 the return 5, otherwise return 6",
}
