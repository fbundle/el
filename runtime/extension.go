package runtime

import (
	"context"
	"el/ast"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Extension struct {
	Name Name
	Man  string
	Exec func(ctx context.Context, values ...Object) adt.Result[Object]
}

func (ext Extension) Module() Function {
	return Function{
		repr: ext.Man,
		exec: func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object] {
			var args []Object
			if err := r.stepAndUnwrapArgs(ctx, frame, argExprList).Unwrap(&args); err != nil {
				return errValue(err)
			}

			return ext.Exec(ctx, args...)
		},
	}
}
