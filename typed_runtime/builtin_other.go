package runtime

import (
	"context"
	"el/ast"
	"el/sorts"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var Builtin = map[Name]Object{
	sorts.Unit:     NilType,
	sorts.Any:      AnyType,
	"builtin_type": BuiltinType,
	"let":          makeData(letFunc, BuiltinType),
	"match":        makeData(matchFunc, BuiltinType),
	"lambda":       makeData(lambdaFunc, BuiltinType),
}

type Extension struct {
	Name Name
	Man  string
	Exec func(ctx context.Context, values ...Object) adt.Result[Object]
}

func (ext Extension) Module() FuncData {
	return FuncData{
		repr: ext.Man,
		exec: func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object] {
			var args []Object
			if err := r.stepAndUnwrapArgs(ctx, frame, argExprList).Unwrap(&args); err != nil {
				return resultErr(err)
			}

			return ext.Exec(ctx, args...)
		},
	}
}

// extra types
