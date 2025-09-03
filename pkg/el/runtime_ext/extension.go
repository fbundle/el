package runtime_ext

import (
	"context"
	"el/pkg/el/expr"
	"el/pkg/el/runtime"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Name = runtime.Name
type Module = runtime.Module

type Extension struct {
	Name Name
	Exec func(ctx context.Context, values ...Value) adt.Result[Value]
	Man  string
}

func (ext Extension) Module() Module {
	return Module{
		Man: ext.Man,
		Exec: func(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Result[Value] {
			args := make([]Value, len(argList))
			for i, argExpr := range argList {
				if err := r.Step(ctx, s, argExpr).Unwrap(&args[i]); err != nil {
					return errValue(err)
				}
			}
			if err := r.PostProcessArgsOpt(args).Unwrap(&args); err != nil {
				return errValue(err)
			}

			return ext.Exec(ctx, args...)
		},
	}
}
