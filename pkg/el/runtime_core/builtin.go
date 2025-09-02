package runtime_core

import (
	"context"
	"el/pkg/el/expr"
	"errors"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

func NewBuiltinStack() Stack {
	return PeekAndUpdate(Stack{}.Push(Frame{}), func(frame Frame) Frame {
		for _, m := range []Module{letModule, lambdaModule, matchModule} {
			frame = frame.Set(m.Name, m)
		}
		return frame
	})
}

var letModule = Module{
	Name: "let",
	Exec: func(r Runtime, ctx context.Context, s Stack, args []expr.Expr) adt.Option[Value] {
		if len(args) < 1 {
			return errorObjectString("let requires at least 1 arguments")
		}
		if len(args)%2 != 1 {
			return errorObjectString("let requires odd number of arguments")
		}

		s = s.Push(Frame{}) // new empty frame
		for i := 0; i < len(args)-1; i += 2 {
			// for let, rexpr must be executed in sequence
			lexpr, rexpr := args[i], args[i+1]
			lvalue, ok := lexpr.(expr.Name)
			if !ok {
				return errorObjectString("lvalue must be a Name")
			}

			var rvalue Value
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
	Man: "Module: (let x 3) - assign value 3 to local variable x",
}

var lambdaModule = Module{
	Name: "lambda",
	Exec: func(r Runtime, ctx context.Context, s Stack, args []expr.Expr) adt.Option[Value] {
		if len(args) < 1 {
			return errorObjectString("lambda requires at least 1 arguments")
		}
		paramList := make([]Name, 0, len(args)-1)
		for i := 0; i < len(args)-1; i++ {
			lvalue, ok := args[i].(expr.Name)
			if !ok {
				return errorObjectString("lvalue must be a Name")
			}
			paramList = append(paramList, Name(lvalue))
		}
		implementation := args[len(args)-1]
		closure := s.Peek() // capture only top of stack // TODO - capture more but only necessary variables

		return object(Lambda{
			Params:  paramList,
			Body:    implementation,
			Closure: closure,
		})
	},
	Man: "Module: (lambda x y (add x y) - declare a function",
}

var matchModule = Module{
	Name: "match",
	Exec: func(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Option[Value] {
		if len(argList) < 3 {
			return errorObjectString("match requires at least 3 arguments")
		}
		if len(argList)%2 != 0 {
			return errorObjectString("match requires even number of arguments")
		}
		var cond Value
		if err := r.StepOpt(ctx, s, argList[0]).Unwrap(&cond); err != nil {
			return errorObject(err)
		}

		for i := 1; i < len(argList)-1; i += 2 {
			var comp Value
			if err := r.StepOpt(ctx, s, argList[i]).Unwrap(&comp); err != nil {
				return errorObject(err)
			}
			if comp == cond {
				return r.StepOpt(ctx, s, argList[i+1])
			}
		}
		return r.StepOpt(ctx, s, argList[len(argList)-1])
	},
	Man: "Module: (match x 1 2 4 5 6) - match, if x=1 then return 3, if x=4 the return 5, otherwise return 6",
}

func errorObjectString(msg string) adt.Option[Value] {
	return errorObject(errors.New(msg))
}
