package runtime

import (
	"context"
	"el/pkg/el/expr"
	"errors"
	"fmt"
)

func (r *Runtime) LoadConstant(name Name, value Object) *Runtime {
	head := r.Stack.Pop()
	head[name] = value
	r.Stack.Push(head)
	return r
}

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

func makeModuleFromExtension(ext Extension) Module {
	return Module{
		Name: ext.Name,
		Exec: func(ctx context.Context, r *Runtime, e expr.Lambda) (Object, error) {
			var args []Object
			for _, argExpr := range e.Args {
				arg, err := r.Step(ctx, argExpr)
				if err != nil {
					return nil, err
				}
				args = append(args, arg)
			}

			args, err := unwrapArgs(args)
			if err != nil {
				return nil, err
			}

			return ext.Exec(ctx, args...)
		},
		Man: ext.Man,
	}
}

var listExtension = Extension{
	Name: "list",
	Exec: func(ctx context.Context, values ...Object) (Object, error) {
		l := List{}
		for _, v := range values {
			l = append(l, v)
		}
		return l, nil
	},
	Man: "module: (list 1 2 (lambda x (add x 1))) - make a list",
}

var lenExtension = Extension{
	Name: "len",
	Exec: func(ctx context.Context, values ...Object) (Object, error) {
		if len(values) != 1 {
			return nil, fmt.Errorf("len requires 1 argument")
		}
		l, ok := values[0].(List)
		if !ok {
			return nil, fmt.Errorf("len argument must be a list")
		}
		return Int(len(l)), nil
	},
	Man: "module: (len (list 1 2 3)) - get the length of a list",
}

var rangeExtension = Extension{
	Name: "range",
	Exec: func(ctx context.Context, values ...Object) (Object, error) {
		if len(values) != 2 {
			return nil, fmt.Errorf("range requires 2 arguments")
		}
		i, ok := values[0].(Int)
		if !ok {
			return nil, fmt.Errorf("range beg non-list value")
		}
		j, ok := values[1].(Int)
		if !ok {
			return nil, fmt.Errorf("range end non-list value")
		}
		output := make(List, 0)
		for k := i; k < j; k++ {
			output = append(output, k)
		}
		return output, nil
	},
	Man: "module: (range m n) - make a list of integers from m to n-1",
}

var sliceExtension = Extension{
	Name: "slice",
	Exec: func(ctx context.Context, values ...Object) (Object, error) {
		if len(values) != 2 {
			return nil, fmt.Errorf("slice requires 2 arguments")
		}
		l, ok := values[0].(List)
		if !ok {
			return nil, fmt.Errorf("slice first argument not a list")
		}
		i, ok := values[1].(List)
		if !ok {
			return nil, fmt.Errorf("slice first argument not a list")
		}
		output := make(List, 0)
		for _, index := range i {
			if index, ok := index.(Int); ok {
				output = append(output, l[index])
			} else {
				return nil, fmt.Errorf("slice index must be an integer")
			}
		}
		return output, nil
	},
	Man: "module: (get (list 1 2 3) (list 0 2)) - get the 0th and 2nd element of a list",
}

func makeArithExtension(name Name, op func(vs ...Int) (Int, error)) Extension {
	return Extension{
		Name: name,
		Exec: func(ctx context.Context, values ...Object) (Object, error) {
			vs := make([]Int, 0)
			for _, v := range values {
				if v, ok := v.(Int); ok {
					vs = append(vs, v)
				} else {
					return nil, fmt.Errorf("%s argument must be an integer", name)
				}
			}
			return op(vs...)
		},
		Man: fmt.Sprintf("module: (%s 1 2 3)", name),
	}
}

func boolToBool(b bool) Int {
	if b {
		return True
	} else {
		return False
	}
}

var eqExtension = makeArithExtension("eq", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("eq requires 2 arguments")
	}
	return boolToBool(vs[0] == vs[1]), nil
})

var neExtension = makeArithExtension("ne", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("ne requires 2 arguments")
	}
	return boolToBool(vs[0] != vs[1]), nil
})

var ltExtension = makeArithExtension("lt", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("lt requires 2 arguments")
	}
	return boolToBool(vs[0] < vs[1]), nil
})

var leExtension = makeArithExtension("le", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("le requires 2 arguments")
	}
	return boolToBool(vs[0] <= vs[1]), nil
})

var gtExtension = makeArithExtension("gt", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("gt requires 2 arguments")
	}
	return boolToBool(vs[0] > vs[1]), nil
})

var geExtension = makeArithExtension("ge", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("ge requires 2 arguments")
	}
	return boolToBool(vs[0] >= vs[1]), nil
})

var addExtension = makeArithExtension("add", func(vs ...Int) (Int, error) {
	output := Int(0)
	for _, v := range vs {
		output += v
	}
	return output, nil
})

var subExtension = makeArithExtension("sub", func(vs ...Int) (Int, error) {
	if len(vs) == 0 {
		return Int(0), errors.New("sub requires at least 1 argument")
	}
	output := vs[0]
	for i := 1; i < len(vs); i++ {
		v := vs[i]
		output -= v
	}
	return output, nil
})

var mulExtension = makeArithExtension("mul", func(vs ...Int) (Int, error) {
	output := Int(1)
	for _, v := range vs {
		output *= v
	}
	return output, nil
})

var divExtension = makeArithExtension("div", func(vs ...Int) (Int, error) {
	if len(vs) == 0 {
		return Int(0), errors.New("div requires at least 1 argument")
	}
	output := vs[0]
	for i := 1; i < len(vs); i++ {
		v := vs[i]
		output /= v
	}
	return output, nil
})

var modExtension = makeArithExtension("mod", func(vs ...Int) (Int, error) {
	if len(vs) == 0 {
		return Int(0), errors.New("mod requires at least 1 argument")
	}
	output := vs[0]
	for i := 1; i < len(vs); i++ {
		v := vs[i]
		output %= v
	}
	return output, nil
})
