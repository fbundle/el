package runtime

import (
	"context"
	"el/pkg/el/expr"
	"el/pkg/el/obj"
	"errors"
	"fmt"
)

func (r *Runtime) LoadConstant(name obj.Name, value obj.Object) *Runtime {
	head := r.Stack.Pop()
	head[name] = value
	r.Stack.Push(head)
	return r
}

type Extension struct {
	Name obj.Name
	Exec func(ctx context.Context, values ...obj.Object) (obj.Object, error)
	Man  string
}

func (r *Runtime) LoadExtension(es ...Extension) *Runtime {
	for _, e := range es {
		r.LoadModule(makeModuleFromExtension(e))
	}
	return r
}

func makeModuleFromExtension(ext Extension) obj.Module[Runtime] {
	return obj.Module[Runtime]{
		Name: ext.Name,
		Exec: func(ctx context.Context, r *Runtime, e expr.Lambda) (obj.Object, error) {
			args, err := r.stepMany(ctx, e.Args...)
			if err != nil {
				return nil, err
			}
			args, err = unwrapArgs(args)
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
	Exec: func(ctx context.Context, values ...obj.Object) (obj.Object, error) {
		l := obj.List{}
		for _, v := range values {
			l = append(l, v)
		}
		return l, nil
	},
	Man: "module: (list 1 2 (lambda x (add x 1))) - make a list",
}

var lenExtension = Extension{
	Name: "len",
	Exec: func(ctx context.Context, values ...obj.Object) (obj.Object, error) {
		if len(values) != 1 {
			return nil, fmt.Errorf("len requires 1 argument")
		}
		l, ok := values[0].(obj.List)
		if !ok {
			return nil, fmt.Errorf("len argument must be a list")
		}
		return obj.Int(len(l)), nil
	},
	Man: "module: (len (list 1 2 3)) - get the length of a list",
}

var rangeExtension = Extension{
	Name: "range",
	Exec: func(ctx context.Context, values ...obj.Object) (obj.Object, error) {
		if len(values) != 2 {
			return nil, fmt.Errorf("range requires 2 arguments")
		}
		i, ok := values[0].(obj.Int)
		if !ok {
			return nil, fmt.Errorf("range beg non-list value")
		}
		j, ok := values[1].(obj.Int)
		if !ok {
			return nil, fmt.Errorf("range end non-list value")
		}
		output := make(obj.List, 0)
		for k := i; k < j; k++ {
			output = append(output, k)
		}
		return output, nil
	},
	Man: "module: (range m n) - make a list of integers from m to n-1",
}

var sliceExtension = Extension{
	Name: "slice",
	Exec: func(ctx context.Context, values ...obj.Object) (obj.Object, error) {
		if len(values) != 2 {
			return nil, fmt.Errorf("slice requires 2 arguments")
		}
		l, ok := values[0].(obj.List)
		if !ok {
			return nil, fmt.Errorf("slice first argument not a list")
		}
		i, ok := values[1].(obj.List)
		if !ok {
			return nil, fmt.Errorf("slice first argument not a list")
		}
		output := make(obj.List, 0)
		for _, index := range i {
			if index, ok := index.(obj.Int); ok {
				output = append(output, l[index])
			} else {
				return nil, fmt.Errorf("slice index must be an integer")
			}
		}
		return output, nil
	},
	Man: "module: (get (list 1 2 3) (list 0 2)) - get the 0th and 2nd element of a list",
}

func makeArithExtension(name obj.Name, op func(vs ...obj.Int) (obj.Int, error)) Extension {
	return Extension{
		Name: name,
		Exec: func(ctx context.Context, values ...obj.Object) (obj.Object, error) {
			vs := make([]obj.Int, 0)
			for _, v := range values {
				if v, ok := v.(obj.Int); ok {
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

func boolToBool(b bool) obj.Int {
	if b {
		return obj.True
	} else {
		return obj.False
	}
}

var eqExtension = makeArithExtension("eq", func(vs ...obj.Int) (obj.Int, error) {
	if len(vs) != 2 {
		return obj.False, errors.New("eq requires 2 arguments")
	}
	return boolToBool(vs[0] == vs[1]), nil
})

var neExtension = makeArithExtension("ne", func(vs ...obj.Int) (obj.Int, error) {
	if len(vs) != 2 {
		return obj.False, errors.New("ne requires 2 arguments")
	}
	return boolToBool(vs[0] != vs[1]), nil
})

var ltExtension = makeArithExtension("lt", func(vs ...obj.Int) (obj.Int, error) {
	if len(vs) != 2 {
		return obj.False, errors.New("lt requires 2 arguments")
	}
	return boolToBool(vs[0] < vs[1]), nil
})

var leExtension = makeArithExtension("le", func(vs ...obj.Int) (obj.Int, error) {
	if len(vs) != 2 {
		return obj.False, errors.New("le requires 2 arguments")
	}
	return boolToBool(vs[0] <= vs[1]), nil
})

var gtExtension = makeArithExtension("gt", func(vs ...obj.Int) (obj.Int, error) {
	if len(vs) != 2 {
		return obj.False, errors.New("gt requires 2 arguments")
	}
	return boolToBool(vs[0] > vs[1]), nil
})

var geExtension = makeArithExtension("ge", func(vs ...obj.Int) (obj.Int, error) {
	if len(vs) != 2 {
		return obj.False, errors.New("ge requires 2 arguments")
	}
	return boolToBool(vs[0] >= vs[1]), nil
})

var addExtension = makeArithExtension("add", func(vs ...obj.Int) (obj.Int, error) {
	output := obj.Int(0)
	for _, v := range vs {
		output += v
	}
	return output, nil
})

var subExtension = makeArithExtension("sub", func(vs ...obj.Int) (obj.Int, error) {
	if len(vs) == 0 {
		return obj.Int(0), errors.New("sub requires at least 1 argument")
	}
	output := vs[0]
	for i := 1; i < len(vs); i++ {
		v := vs[i]
		output -= v
	}
	return output, nil
})

var mulExtension = makeArithExtension("mul", func(vs ...obj.Int) (obj.Int, error) {
	output := obj.Int(1)
	for _, v := range vs {
		output *= v
	}
	return output, nil
})

var divExtension = makeArithExtension("div", func(vs ...obj.Int) (obj.Int, error) {
	if len(vs) == 0 {
		return obj.Int(0), errors.New("div requires at least 1 argument")
	}
	output := vs[0]
	for i := 1; i < len(vs); i++ {
		v := vs[i]
		output /= v
	}
	return output, nil
})

var modExtension = makeArithExtension("mod", func(vs ...obj.Int) (obj.Int, error) {
	if len(vs) == 0 {
		return obj.Int(0), errors.New("mod requires at least 1 argument")
	}
	output := vs[0]
	for i := 1; i < len(vs); i++ {
		v := vs[i]
		output %= v
	}
	return output, nil
})
