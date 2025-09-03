package runtime

import (
	"context"
	"el/ast"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var BuiltinStack Stack

func init() {
	s := Stack{}.Push(Frame{})
	s = updateHead(s, func(frame Frame) Frame {
		frame = frame.Set("type", typeFunc)
		frame = frame.Set("let", letFunc)
		frame = frame.Set("lambda", lambdaFunc)
		frame = frame.Set("match", matchFunc)
		return frame
	})
	BuiltinStack = s
}

var typeFunc = Function{
	repr: "[module: (type (list 1 2 3)) - get the type of an arbitrary object]",
	exec: func(r Runtime, ctx context.Context, s Stack, argList []ast.Expr) adt.Result[Object] {
		if len(argList) != 1 {
			return errValueString("type requires 1 argument")
		}
		var v Object
		if err := r.Step(ctx, s, argList[0]).Unwrap(&v); err != nil {
			return errValue(err)
		}
		return value(v.Type())
	},
}

var letFunc = Function{
	repr: "[module: (let x 3) - assign value 3 to local variable x]",
	exec: func(r Runtime, ctx context.Context, s Stack, argList []ast.Expr) adt.Result[Object] {
		if len(argList) < 1 || len(argList)%2 != 1 {
			return errValueString("let requires at least 1 arguments and odd number of arguments")
		}

		s = s.Push(Frame{}) // push a new empty frame

		var lvalue ast.Leaf
		var rvalue Object
		for i := 0; i < len(argList)-1; i += 2 {
			lexpr, rexpr := argList[i], argList[i+1]
			if ok := adt.Cast[ast.Leaf](lexpr).Unwrap(&lvalue); !ok {
				return errValueString("lvalue must be a Leaf")
			}
			if err := r.Step(ctx, s, rexpr).Unwrap(&rvalue); err != nil {
				return errValue(err)
			}
			// update stack
			s = updateHead(s, func(f Frame) Frame {
				return f.Set(Name(lvalue), rvalue)
			})
		}
		return r.Step(ctx, s, argList[len(argList)-1])
	},
}

var matchFunc = Function{
	repr: "[module: (match x 1 2 4 5 6) - match, if x=1 then return 3, if x=4 the return 5, otherwise return 6]",
	exec: func(r Runtime, ctx context.Context, s Stack, argList []ast.Expr) adt.Result[Object] {
		if len(argList) < 2 || len(argList)%2 != 0 {
			return errValueString("match requires at least 2 arguments and even number of arguments")
		}
		var cond Object
		if err := r.Step(ctx, s, argList[0]).Unwrap(&cond); err != nil {
			return errValue(err)
		}

		var finalRexpr ast.Expr = argList[len(argList)-1]
		var comp Object
		for i := 1; i < len(argList)-1; i += 2 {
			lexpr, rexpr := argList[i], argList[i+1]
			if err := r.Step(ctx, s, lexpr).Unwrap(&comp); err != nil {
				return errValue(err)
			}
			if comp == cond {
				finalRexpr = rexpr
				break
			}
		}
		return r.Step(ctx, s, finalRexpr)
	},
}

var lambdaFunc = Function{
	repr: "[module: (lambda x y (add x y) - declare a function]",
	exec: func(r Runtime, ctx context.Context, s Stack, argList []ast.Expr) adt.Result[Object] {
		if len(argList) < 1 {
			return errValueString("lambda requires at least 1 arguments")
		}

		paramList := make([]Name, 0, len(argList)-1)
		var lvalue ast.Leaf
		for i := 0; i < len(argList)-1; i++ {
			lexpr := argList[i]
			if ok := adt.Cast[ast.Leaf](lexpr).Unwrap(&lvalue); !ok {
				return errValueString("lvalue must be a Leaf")
			}
			paramList = append(paramList, Name(lvalue))
		}

		body := argList[len(argList)-1]
		local := s.Peek()
		for _, name := range paramList {
			local = local.Del(name) // remove all the parameters from the local
		}

		return value(Function{
			repr: makeLambdaRepr(paramList, body, local),
			exec: makeLambdaExec(paramList, body, local),
		})
	},
}

func makeLambdaRepr(paramList []Name, body ast.Expr, local Frame) string {
	repr := fmt.Sprintf("(%s;", local.Repr())
	for _, param := range paramList {
		repr += string(param) + " "
	}
	repr += "-> " + body.String()
	repr += ")"
	return repr
}

func makeLambdaExec(paramList []Name, body ast.Expr, local Frame) Exec {
	return func(r Runtime, ctx context.Context, s Stack, argList []ast.Expr) adt.Result[Object] {
		/*
			for recursive function, the name of that function is in s Stack
		*/
		// 0. sanity check
		if len(argList) < len(paramList) {
			errValue(ErrorNotEnoughArguments)
		}
		// 1. evaluate arguments
		var args []Object
		if err := r.stepAndUnwrapArgs(ctx, s, argList).Unwrap(&args); err != nil {
			return errValue(err)
		}
		// 2. make call stack
		for i := 0; i < len(paramList); i++ {
			param, arg := paramList[i], args[i]
			local = local.Set(param, arg)
		}

		callStack := s.Push(local)
		// 3. make call with new stack - signal tailcall to children
		var o Object
		if err := r.Step(ctx, callStack, body).Unwrap(&o); err != nil {
			return errValue(err)
		}
		return value(o)
	}
}
