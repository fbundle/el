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

type Exec = func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object]

type FuncData struct {
	exec Exec
	repr string
}

func (f FuncData) String() string {
	return f.repr
}

var letFunc = FuncData{
	repr: "{builtin: (let x 3 4) - assign value 3 to local variable x then return 4}",
	exec: func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object] {
		if len(argExprList) == 0 || len(argExprList)%2 != 1 {
			return resultErrStr("let requires at least 1 arguments and odd number of arguments")
		}

		lastExpr := argExprList[len(argExprList)-1]
		var lExprList, rExprList []ast.Expr
		for i := 0; i < len(argExprList)-1; i += 2 {
			lExprList = append(lExprList, argExprList[i])
			rExprList = append(rExprList, argExprList[i+1])
		}

		for lexpr, rexpr := range zip(lExprList, rExprList) {
			lvalue, ok := lexpr.(ast.Name)
			if !ok {
				return resultErrStr(fmt.Sprintf("lvalue must be a Name: %s", lexpr.String()))
			}
			var rvalue Object
			if err := r.Step(ctx, frame, rexpr).Unwrap(&rvalue); err != nil {
				return resultErr(err)
			}
			frame = frame.Set(Name(lvalue), rvalue)
		}
		return r.Step(ctx, frame, lastExpr)
	},
}

var matchFunc = FuncData{
	repr: "{builtin: (match x 1 2 3 4 5) - match, if x=1 then return 2, if x=3 the return 4, otherwise return 5",
	exec: func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object] {
		if len(argExprList) < 2 || len(argExprList)%2 != 0 {
			return resultErrStr("match requires at least 2 arguments and even number of arguments")
		}
		condExpr := argExprList[0]
		lastExpr := argExprList[len(argExprList)-1]
		var lExprList, rExprList []ast.Expr
		for i := 1; i < len(argExprList)-1; i += 2 {
			lExprList = append(lExprList, argExprList[i])
			rExprList = append(rExprList, argExprList[i+1])
		}
		var cond Object
		if err := r.Step(ctx, frame, condExpr).Unwrap(&cond); err != nil {
			return resultErr(err)
		}

		for lexpr, rexpr := range zip(lExprList, rExprList) {
			var comp Object
			if err := r.Step(ctx, frame, lexpr).Unwrap(&comp); err != nil {
				return resultErr(err)
			}
			var isEqual bool
			if err := equal(cond.Data(), comp.Data()).Unwrap(&isEqual); err != nil {
				return resultErr(err)
			}
			if isEqual {
				lastExpr = rexpr
				break
			}
		}
		return r.Step(ctx, frame, lastExpr)
	},
}

var lambdaFunc = FuncData{
	repr: "{builtin: (lambda x y (add x y) - declare a function}",
	exec: func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object] {
		if len(argExprList) < 1 {
			return resultErrStr("lambda requires at least 1 arguments")
		}

		lastExpr := argExprList[len(argExprList)-1]
		paramExprList := argExprList[:len(argExprList)-1]

		paramList := make([]Name, 0, len(argExprList)-1)
		for _, paramExpr := range paramExprList {
			lvalue, ok := paramExpr.(ast.Name)
			if !ok {
				return resultErrStr(fmt.Sprintf("lvalue must be a Name: %s", paramExpr.String()))
			}
			paramList = append(paramList, Name(lvalue))
		}

		closure := frame
		for _, name := range paramList {
			closure = closure.Del(name) // remove all the parameters from the local
		}

		return resultObj(makeFunction(paramList, lastExpr, closure))
	},
}

func makeFunction(paramList []Name, body ast.Expr, closure Frame) Object {
	funcData := FuncData{
		repr: makeLambdaRepr(paramList, body, closure),
		exec: makeLambdaExec(paramList, body, closure),
	}
	funcType := makeWeakestType(len(paramList))
	return makeData(funcData, funcType)
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
			return resultErr(err)
		}
		// 2. add params to closure
		for param, arg := range zip(paramList, argList) {
			closure = closure.Set(param, arg)
		}

		if len(argList) >= len(paramList) {
			// 3. add environment frame into closure and make call
			for k, v := range frame.Iter {
				if _, ok := closure.Get(k); !ok {
					closure = closure.Set(k, v)
				}
			}
			var o Object
			if err := r.Step(ctx, closure, body).Unwrap(&o); err != nil {
				return resultErr(err)
			}
			return resultObj(o)
		} else {
			// 3. currying
			return resultObj(makeFunction(paramList[len(argList):], body, closure))
		}
	}
}

func zip[T1 any, T2 any](l1 []T1, l2 []T2) func(yield func(T1, T2) bool) {
	return func(yield func(T1, T2) bool) {
		length := min(len(l1), len(l2))
		for i := 0; i < length; i++ {
			if ok := yield(l1[i], l2[i]); !ok {
				return
			}
		}
	}
}

func equal(o1 any, o2 any) adt.Result[bool] {
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
