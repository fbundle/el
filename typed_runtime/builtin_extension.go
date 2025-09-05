package runtime

import (
	"context"
	"el/ast"
	"el/sorts"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

func init() {
	Builtin["type.of"] = makeData(typeOfExtension.Module(), BuiltinType)
	Builtin["type.cast"] = makeData(typeCastExtension.Module(), BuiltinType)
	Builtin["type.chain"] = makeData(typeChainExtension.Module(), BuiltinType)
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
			var argList []Object
			if err := r.stepAndUnwrapArgs(ctx, frame, argExprList).Unwrap(&argList); err != nil {
				return resultErr(err)
			}
			return ext.Exec(ctx, argList...)
		},
	}
}

var typeOfExtension = Extension{
	Name: "type.of",
	Man:  "",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		if len(values) != 1 {
			return resultErrStrf("type.of expected 1 argument")
		}
		return resultObj(values[0].Type())
	},
}

var typeCastExtension = Extension{
	Name: "type.cast",
	Man:  "",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		if len(values) != 2 {
			return resultErrStrf("type.cast expected 2 arguments")
		}
		parent, object := values[0], values[1]
		var newObject Object
		if ok := object.Cast(parent).Unwrap(&newObject); !ok {
			return resultErrStrf("cannot cast object %s of type %s into type %s", object, object.Type(), parent)
		}
		return resultObj(newObject)
	},
}

var typeChainExtension = Extension{
	Name: "type.chain",
	Man:  "",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		if len(values) == 0 {
			return resultErrStrf("type.chain expected at least 1 argument")
		}
		sortList := make([]sorts.Sort, 0, len(values))
		for _, value := range values {
			sortList = append(sortList, value.Sort())
		}
		var newSort sorts.Sort
		if ok := sorts.Arrow(sortList...).Unwrap(&newSort); !ok {
			return resultErrStrf("cannot make sort from %s", sortList)
		}
		return resultObj(makeSort(newSort))
	},
}
