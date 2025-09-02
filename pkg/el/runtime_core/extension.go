package runtime_core

import (
	"context"
	"el/pkg/el/expr"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Extension struct {
	Name Name
	Exec func(ctx context.Context, values ...Object) adt.Option[Object]
	Man  string
}

func MakeModuleFromExtension(ext Extension) Module {
	return Module{
		Name: ext.Name,
		Man:  ext.Man,
		Exec: func(r Runtime, ctx context.Context, s Stack, e expr.Lambda) adt.Option[Object] {
			args := make([]Object, len(e.Args))
			for i, argExpr := range e.Args {
				if err := r.StepOpt(ctx, s, argExpr).Unwrap(&args[i]); err != nil {
					return errorObject(err)
				}
			}
			if err := r.UnwrapArgsOpt(args).Unwrap(&args); err != nil {
				return errorObject(err)
			}

			return ext.Exec(ctx, args...)
		},
	}
}
