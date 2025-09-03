package runtime_ext

import (
	"context"
	"el/pkg/el/expr"
	"el/pkg/el/runtime_core"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Name = runtime_core.Name
type Module = runtime_core.Module

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

var listExtension = Extension{
	Name: "list",
	Man:  "Module: (list 1 2 (lambda x (add x 1))) - make a list",
	Exec: func(ctx context.Context, values ...Value) adt.Result[Value] {
		l := List{}
		for _, v := range values {
			l = List{l.Ins(l.Len(), v)}
		}
		return value(l)
	},
}

var lenExtension = Extension{
	Name: "len",
	Man:  "Module: (len (list 1 2 3)) - get the length of a list",
	Exec: func(ctx context.Context, values ...Value) adt.Result[Value] {
		if len(values) != 1 {
			return errValueString("len requires 1 argument")
		}
		var l List
		if ok := adt.Cast[List](values[0]).Unwrap(&l); !ok {
			return errValueString("len argument must be a list")
		}
		return value(Int{l.Len()})
	},
}

var sliceExtension = Extension{
	Name: "slice",
	Man:  "Module: (get (list 1 2 3) (list 0 2)) - get the 0th and 2nd element of a list",
	Exec: func(ctx context.Context, values ...Value) adt.Result[Value] {
		if len(values) != 2 {
			return errValueString("slice requires 2 arguments")
		}
		var l List
		if ok := adt.Cast[List](values[0]).Unwrap(&l); !ok {
			return errValueString("slice first argument must be a list")
		}
		var i List
		if ok := adt.Cast[List](values[1]).Unwrap(&i); !ok {
			return errValueString("slice second argument must be a list of integers")
		}
		output := List{}
		for _, o := range i.Iter {
			var index Int
			if ok := adt.Cast[Int](o).Unwrap(&index); !ok {
				return errValueString("slice second argument must be a list of integers")
			}
			v := l.Get(index.int)
			output = List{output.Ins(output.Len(), v)}
		}
		return value(output)
	},
}
