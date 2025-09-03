package runtime

import (
	"context"
	"el/ast"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var BuiltinFrame Frame

func init() {
	frame := Frame{}
	frame = frame.Set("type", typeFunc)
	frame = frame.Set("let", letFunc)
	frame = frame.Set("lambda", lambdaFunc)
	frame = frame.Set("match", matchFunc)
	frame = frame.Set("nil", Nil{})
	BuiltinFrame = frame
}

var typeFunc = Function{
	repr: "[module: (type (list 1 2 3)) - get the type of an arbitrary object]",
	exec: func(r Runtime, ctx context.Context, frame Frame, argList []ast.Expr) adt.Result[Object] {
		if len(argList) != 1 {
			return errValueString("type requires 1 argument")
		}
		var v Object
		if err := r.Step(ctx, frame, argList[0]).Unwrap(&v); err != nil {
			return errValue(err)
		}
		return value(v.Type())
	},
}

var letFunc = Function{
	repr: "[module: (let x 3) - assign value 3 to local variable x]",
	exec: func(r Runtime, ctx context.Context, frame Frame, argList []ast.Expr) adt.Result[Object] {
		if len(argList) < 1 || len(argList)%2 != 1 {
			return errValueString("let requires at least 1 arguments and odd number of arguments")
		}

		var lvalue ast.Name
		var rvalue Object
		for i := 0; i < len(argList)-1; i += 2 {
			lexpr, rexpr := argList[i], argList[i+1]
			if ok := adt.Cast[ast.Name](lexpr).Unwrap(&lvalue); !ok {
				return errValueString("lvalue must be a Name")
			}
			if err := r.Step(ctx, frame, rexpr).Unwrap(&rvalue); err != nil {
				return errValue(err)
			}
			frame = frame.Set(Name(lvalue), rvalue)
		}
		return r.Step(ctx, frame, argList[len(argList)-1])
	},
}

var matchFunc = Function{
	repr: "[module: (match x 1 2 4 5 6) - match, if x=1 then return 3, if x=4 the return 5, otherwise return 6]",
	exec: func(r Runtime, ctx context.Context, frame Frame, argList []ast.Expr) adt.Result[Object] {
		if len(argList) < 2 || len(argList)%2 != 0 {
			return errValueString("match requires at least 2 arguments and even number of arguments")
		}
		var cond Object
		if err := r.Step(ctx, frame, argList[0]).Unwrap(&cond); err != nil {
			return errValue(err)
		}

		var finalRexpr = argList[len(argList)-1]
		for i := 1; i < len(argList)-1; i += 2 {
			lexpr, rexpr := argList[i], argList[i+1]
			var comp Object
			if err := r.Step(ctx, frame, lexpr).Unwrap(&comp); err != nil {
				return errValue(err)
			}
			var isEqual bool
			if err := equal(cond, comp).Unwrap(&isEqual); err != nil {
				return errValue(err)
			}
			if isEqual {
				finalRexpr = rexpr
				break
			}
		}
		return r.Step(ctx, frame, finalRexpr)
	},
}

var lambdaFunc = Function{
	repr: "[module: (lambda x y (add x y) - declare a function]",
	exec: func(r Runtime, ctx context.Context, frame Frame, argList []ast.Expr) adt.Result[Object] {
		if len(argList) < 1 {
			return errValueString("lambda requires at least 1 arguments")
		}

		paramList := make([]Name, 0, len(argList)-1)
		var lvalue ast.Name
		for i := 0; i < len(argList)-1; i++ {
			lexpr := argList[i]
			if ok := adt.Cast[ast.Name](lexpr).Unwrap(&lvalue); !ok {
				return errValueString("lvalue must be a Name")
			}
			paramList = append(paramList, Name(lvalue))
		}

		body := argList[len(argList)-1]
		closure := frame
		for _, name := range paramList {
			closure = closure.Del(name) // remove all the parameters from the local
		}

		return value(Function{
			repr: makeLambdaRepr(paramList, body, closure),
			exec: makeLambdaExec(paramList, body, closure),
		})
	},
}

func makeLambdaRepr(paramList []Name, body ast.Expr, closure Frame) string {
	paramNameList := make([]string, 0, len(paramList))
	for _, name := range paramList {
		paramNameList = append(paramNameList, string(name))
	}
	return fmt.Sprintf("(closure{...}; %s -> %s)", strings.Join(paramNameList, " "), body.String())
}

func makeLambdaExec(paramList []Name, body ast.Expr, closure Frame) Exec {
	return func(r Runtime, ctx context.Context, frame Frame, argList []ast.Expr) adt.Result[Object] {
		/*
			for recursive function, the name of that function is in `frame`
		*/
		// 0. sanity check
		if len(argList) < len(paramList) {
			return errValue(ErrorNotEnoughArguments)
		}
		// 1. evaluate arguments
		var args []Object
		if err := r.stepAndUnwrapArgs(ctx, frame, argList).Unwrap(&args); err != nil {
			return errValue(err)
		}
		// 2. make the call frame
		// for non-recursive function, callFrame = closure + params
		// for recursive function, callFrame = frame + closure + params
		callFrame := frame
		for k, v := range closure.Iter {
			callFrame = callFrame.Set(k, v)
		}
		for i := 0; i < len(paramList); i++ {
			param, arg := paramList[i], args[i]
			callFrame = callFrame.Set(param, arg)
		}
		// 3. make call with new stack - signal tailcall to children
		var o Object
		if err := r.Step(ctx, callFrame, body).Unwrap(&o); err != nil {
			return errValue(err)
		}
		return value(o)
	}
}

func equal(o1 Object, o2 Object) adt.Result[bool] {
	t1 := reflect.TypeOf(o1)
	t2 := reflect.TypeOf(o2)
	if t1 == nil || t2 == nil || !t1.Comparable() || !t2.Comparable() {
		return adt.Err[bool](errors.New("match comparison: non-comparable object"))
	}
	if t1 != t2 {
		return adt.Err[bool](errors.New("match comparison: different types"))
	}
	return adt.Ok(o1 == o2)
}
