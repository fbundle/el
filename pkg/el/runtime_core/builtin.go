package runtime_core

import (
	"context"
	"el/pkg/el/expr"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

func NewBuiltinStack() Stack {
	return UpdateHead(Stack{}.Push(Frame{}), func(frame Frame) Frame {
		for _, m := range []Module{letModule, lambdaModule, matchModule} {
			frame = frame.Set(m.Name, m)
		}
		return frame
	})
}

var letModule = Module{
	Name: "let",
	Exec: func(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Result[Value] {
		if len(argList) < 1 || len(argList)%2 != 1 {
			return errValueString("let requires at least 1 arguments and odd number of arguments")
		}

		s = s.Push(Frame{}) // push a new empty frame

		var lvalue expr.Name
		var rvalue Value
		for i := 0; i < len(argList)-1; i += 2 {
			lexpr, rexpr := argList[i], argList[i+1]
			if ok := adt.Cast[expr.Name](lexpr).Unwrap(&lvalue); !ok {
				return errValueString("lvalue must be a Name")
			}
			if err := r.Step(ctx, s, rexpr).Unwrap(&rvalue); err != nil {
				return errValue(err)
			}
			// update stack
			s = UpdateHead(s, func(f Frame) Frame {
				return f.Set(Name(lvalue), rvalue)
			})
		}

		return r.Step(ctx, s, argList[len(argList)-1])
	},
	Man: "Module: (let x 3) - assign value 3 to local variable x",
}

var lambdaModule = Module{
	Name: "lambda",
	Exec: func(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Result[Value] {
		if len(argList) < 1 {
			return errValueString("lambda requires at least 1 arguments")
		}

		paramList := make([]Name, 0, len(argList)-1)
		var lvalue expr.Name
		for i := 0; i < len(argList)-1; i++ {
			lexpr := argList[i]
			if ok := adt.Cast[expr.Name](lexpr).Unwrap(&lvalue); !ok {
				return errValueString("lvalue must be a Name")
			}
			paramList = append(paramList, Name(lvalue))
		}

		body := argList[len(argList)-1]
		closure := s.Peek() // capture only the top of the frame stack
		return value(Lambda{
			ParamList: paramList,
			Body:      body,
			Closure:   closure,
		})
	},
	Man: "Module: (lambda x y (add x y) - declare a function",
}

var matchModule = Module{
	Name: "match",
	Exec: func(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Result[Value] {
		if len(argList) < 2 || len(argList)%2 != 0 {
			return errValueString("match requires at least 2 arguments and even number of arguments")
		}
		var cond Value
		if err := r.Step(ctx, s, argList[0]).Unwrap(&cond); err != nil {
			return errValue(err)
		}

		var comp Value
		for i := 1; i < len(argList)-1; i += 2 {
			lexpr, rexpr := argList[i], argList[i+1]
			if err := r.Step(ctx, s, lexpr).Unwrap(&comp); err != nil {
				return errValue(err)
			}
			if comp == cond {
				return r.Step(ctx, s, rexpr)
			}
		}
		return r.Step(ctx, s, argList[len(argList)-1])
	},
	Man: "Module: (match x 1 2 4 5 6) - match, if x=1 then return 3, if x=4 the return 5, otherwise return 6",
}
