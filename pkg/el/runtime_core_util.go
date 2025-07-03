package el

import (
	"context"
	"errors"
)

type Extension struct {
	Name Name
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
			unwrappedArgs, err := unwrapArgs(args)
			if err != nil {
				return nil, err
			}
			return e.Exec(ctx, unwrappedArgs...)
		},
		Man: e.Man,
	}
}

func unwrapArgs(args []Object) ([]Object, error) {
	unwrappedArgs := make([]Object, 0, len(args))
	for len(args) > 0 {
		head := args[0]
		if _, ok := head.(Unwrap); ok {
			if len(args) <= 1 {
				return unwrappedArgs, errors.New("unwrapping argument empty")
			}
			next, ok := args[1].(List)
			if !ok {
				return unwrappedArgs, errors.New("unwrapping argument must be a list")
			}
			unwrappedArgs = append(unwrappedArgs, next...)
			args = args[2:]
		} else {
			unwrappedArgs = append(unwrappedArgs, head)
			args = args[1:]
		}
	}
	return unwrappedArgs, nil
}
