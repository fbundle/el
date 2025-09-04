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

var ErrorTooManyArguments = errors.New("too many arguments")

var BuiltinFrame Frame

func init() {
	frame := Frame{}
	for name, object := range builtinObject {
		frame = frame.Set(name, object)
	}
	BuiltinFrame = frame
}

var builtinObject = map[Name]Object{
	"let":    letFunc,
	"match":  matchFunc,
	"lambda": lambdaFunc,
	"nil":    Nil{},
}

var letFunc = Function{
	repr: "{builtin: (let x 3) - assign value 3 to local variable x}",
	exec: func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object] {
		if len(argExprList) < 1 || len(argExprList)%2 != 1 {
			return errValueString("let requires at least 1 arguments and odd number of arguments")
		}

		var lvalue ast.Name
		var rvalue Object
		for i := 0; i < len(argExprList)-1; i += 2 {
			lexpr, rexpr := argExprList[i], argExprList[i+1]
			if ok := adt.Cast[ast.Name](lexpr).Unwrap(&lvalue); !ok {
				return errValueString(fmt.Sprintf("lvalue must be a Name: %s", lexpr.String()))
			}
			if err := r.Step(ctx, frame, rexpr).Unwrap(&rvalue); err != nil {
				return errValue(err)
			}
			frame = frame.Set(Name(lvalue), rvalue)
		}
		return r.Step(ctx, frame, argExprList[len(argExprList)-1])
	},
}

var matchFunc = Function{
	repr: "{builtin: (match x 1 2 3 4 5) - match, if x=1 then return 2, if x=3 the return 4, otherwise return 5",
	exec: func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object] {
		if len(argExprList) < 2 || len(argExprList)%2 != 0 {
			return errValueString("match requires at least 2 arguments and even number of arguments")
		}
		var cond Object
		if err := r.Step(ctx, frame, argExprList[0]).Unwrap(&cond); err != nil {
			return errValue(err)
		}

		var finalRexpr = argExprList[len(argExprList)-1]
		for i := 1; i < len(argExprList)-1; i += 2 {
			lexpr, rexpr := argExprList[i], argExprList[i+1]
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
	repr: "{builtin: (lambda x y (add x y) - declare a function}",
	exec: func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object] {
		if len(argExprList) < 1 {
			return errValueString("lambda requires at least 1 arguments")
		}

		paramList := make([]Name, 0, len(argExprList)-1)
		var lvalue ast.Name
		for i := 0; i < len(argExprList)-1; i++ {
			lexpr := argExprList[i]
			if ok := adt.Cast[ast.Name](lexpr).Unwrap(&lvalue); !ok {
				return errValueString(fmt.Sprintf("lvalue must be a Name: %s", lexpr.String()))
			}
			paramList = append(paramList, Name(lvalue))
		}

		body := argExprList[len(argExprList)-1]
		closure := frame
		for _, name := range paramList {
			closure = closure.Del(name) // remove all the parameters from the local
		}

		return value(makeFunction(paramList, body, closure))
	},
}

func makeFunction(paramList []Name, body ast.Expr, closure Frame) Function {
	return Function{
		repr: makeLambdaRepr(paramList, body, closure),
		exec: makeLambdaExec(paramList, body, closure),
	}
}

func makeLambdaRepr(paramList []Name, body ast.Expr, closure Frame) string {
	paramNameList := make([]string, 0, len(paramList))
	for _, name := range paramList {
		paramNameList = append(paramNameList, string(name))
	}
	return fmt.Sprintf("{closure{...}; %s => %s}", strings.Join(paramNameList, " "), body.String())
}

func makeLambdaExec(paramList []Name, body ast.Expr, closure Frame) Exec {
	return func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object] {
		/*
			for recursive function, the name of that function is in `frame`
		*/

		// 1. evaluate arguments
		var argList []Object
		if err := r.stepAndUnwrapArgs(ctx, frame, argExprList).Unwrap(&argList); err != nil {
			return errValue(err)
		}
		// 2. add params to closure
		zip(paramList, argList, func(param Name, arg Object) bool {
			closure = closure.Set(param, arg)
			return true
		})

		if len(argList) > len(paramList) {
			// 3. too many arguments
			return errValue(ErrorTooManyArguments)
		} else if len(argList) == len(paramList) {
			// 3. add environment frame into closure and make call
			for k, v := range frame.Iter {
				if _, ok := closure.Get(k); !ok {
					closure = closure.Set(k, v)
				}
			}
			var o Object
			if err := r.Step(ctx, closure, body).Unwrap(&o); err != nil {
				return errValue(err)
			}
			return value(o)
		} else {
			// 3. currying
			return value(makeFunction(paramList[len(argList):], body, closure))
		}
	}
}

func zip[T1 any, T2 any](l1 []T1, l2 []T2, yield func(T1, T2) bool) {
	length := min(len(l1), len(l2))
	for i := 0; i < length; i++ {
		if ok := yield(l1[i], l2[i]); !ok {
			return
		}
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
