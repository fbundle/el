package runtime_ext

import (
	"context"
	"el/pkg/el/expr"
	"el/pkg/el/runtime_core"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Object = runtime_core.Object
type Name = runtime_core.Name
type Extension struct {
	Name Name
	Exec func(ctx context.Context, values ...Object) adt.Option[Object]
	Man  string
}

func (ext Extension) Module() Module {
	return Module{
		man: ext.Man,
		exec: func(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Option[Object] {
			args := make([]Object, len(argList))
			for i, argExpr := range argList {
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

var listExtension = Extension{
	Name: "list",
	Man:  "Module: (list 1 2 (lambda x (add x 1))) - make a list",
	Exec: func(ctx context.Context, values ...Object) adt.Option[Object] {
		l := List{}
		for _, v := range values {
			l = List{l.Ins(l.Len(), v)}
		}
		return adt.Some[Object](l)
	},
}

var lenExtension = Extension{
	Name: "len",
	Man:  "Module: (len (list 1 2 3)) - get the length of a list",
	Exec: func(ctx context.Context, values ...Object) adt.Option[Object] {
		if len(values) != 1 {
			return errorObjectString("len requires 1 argument")
		}
		l, ok := values[0].(List)
		if !ok {
			return errorObjectString("len argument must be a list")
		}
		return object(Int{l.Len()})
	},
}

var sliceExtension = Extension{
	Name: "slice",
	Man:  "Module: (get (list 1 2 3) (list 0 2)) - get the 0th and 2nd element of a list",
	Exec: func(ctx context.Context, values ...Object) adt.Option[Object] {
		if len(values) != 2 {
			return errorObjectString("slice requires 2 arguments")
		}
		l, ok := values[0].(List)
		if !ok {
			return errorObjectString("slice first argument not a list")
		}
		i, ok := values[1].(List)
		if !ok {
			return errorObjectString("slice first argument not a list")
		}
		output := List{}
		for _, index := range i.Iter {
			if index, ok := index.(Int); ok {
				v := l.Get(index.int)
				output = List{output.Ins(output.Len(), v)}
			} else {
				return errorObjectString("slice index must be an integer")
			}
		}
		return object(output)
	},
}
