package el

import (
	"context"
	"errors"
)

type Extension struct {
	Name NameExpr
	Exec func(ctx context.Context, values ...Object) (Object, error)
	Man  string
}

func (r *Runtime) LoadExtension(es ...Extension) *Runtime {
	for _, e := range es {
		r.LoadModule(makeModuleFromExtension(e))
	}
	return r
}

func makeModuleFromExtension(e Extension) Module {
	return Module{
		Name: e.Name,
		Exec: func(ctx context.Context, r *Runtime, expr LambdaExpr) (Object, error) {
			args, err := r.stepMany(ctx, expr.Args...)
			if err != nil {
				return nil, err
			}
			var unwrappedArgs []Object
			i := 0
			for i < len(args) {
				if _, ok := args[i].(Unwrap); ok {
					if i+1 >= len(args) {
						return nil, errors.New("unwrapping arguments must be a list")
					}
					argsList, ok := args[i+1].(List)
					if !ok {
						return nil, errors.New("unwrapping arguments must be a list")
					}
					for _, elem := range argsList {
						unwrappedArgs = append(unwrappedArgs, elem)
					}
					i += 2
				} else {
					unwrappedArgs = append(unwrappedArgs, args[i])
					i++
				}
			}
			return e.Exec(ctx, unwrappedArgs...)
		},
		Man: e.Man,
	}
}
