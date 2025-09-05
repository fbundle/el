package runtime_ext

import (
	"context"
	"el/ast"
	"el/runtime"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var namesFunc = runtime.FuncData{
	Repr: "{builtin: (names) - get all names}",
	Exec: func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object] {
		if len(argExprList) > 0 {
			return resultErrStrf("names takes no arguments")
		}
		l := List{}

		for name, _ := range frame.Iter {
			nameStr := String{Val: string(name)}
			l = List{l.PushBack(makeTypedData(nameStr))}
		}
		return resultTypedData(l)
	},
}
