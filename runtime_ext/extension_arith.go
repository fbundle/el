package runtime_ext

import (
	"context"
	runtime "el/runtime"
	"errors"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Name = runtime.Name

type Extension = runtime.Extension

func makeArithExtension(name string, f func(...Int) (Int, error)) Extension {
	return Extension{
		Name: Name(name),
		Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
			vs := make([]Int, len(values))
			for i, val := range values {
				v, ok := val.Data().(Int)
				if !ok {
					return resultErrStrf(fmt.Sprintf("%s argument must be an integer", name))
				}
				vs[i] = v
			}
			output, err := f(vs...)
			if err != nil {
				return resultErr(err)
			}
			return resultTypedData(output)
		},
		Man: fmt.Sprintf("{arithmetic_extension: %s}", name),
	}
}

var True = Int{1}
var False = Int{0}

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
	return boolToBool(vs[0].Val == vs[1].Val), nil
})

var neExtension = makeArithExtension("ne", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("ne requires 2 arguments")
	}
	return boolToBool(vs[0].Val != vs[1].Val), nil
})

var ltExtension = makeArithExtension("lt", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("lt requires 2 arguments")
	}
	return boolToBool(vs[0].Val < vs[1].Val), nil
})

var leExtension = makeArithExtension("le", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("le requires 2 arguments")
	}
	return boolToBool(vs[0].Val <= vs[1].Val), nil
})

var gtExtension = makeArithExtension("gt", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("gt requires 2 arguments")
	}
	return boolToBool(vs[0].Val > vs[1].Val), nil
})

var geExtension = makeArithExtension("ge", func(vs ...Int) (Int, error) {
	if len(vs) != 2 {
		return False, errors.New("ge requires 2 arguments")
	}
	return boolToBool(vs[0].Val >= vs[1].Val), nil
})

var addExtension = makeArithExtension("add", func(vs ...Int) (Int, error) {
	output := Int{0}
	for _, v := range vs {
		output.Val += v.Val
	}
	return output, nil
})

var subExtension = makeArithExtension("sub", func(vs ...Int) (Int, error) {
	if len(vs) == 0 {
		return Int{}, errors.New("sub requires at least 1 argument")
	}
	output := vs[0]
	for i := 1; i < len(vs); i++ {
		v := vs[i]
		output.Val -= v.Val
	}
	return output, nil
})

var mulExtension = makeArithExtension("mul", func(vs ...Int) (Int, error) {
	output := Int{1}
	for _, v := range vs {
		output.Val *= v.Val
	}
	return output, nil
})

var divExtension = makeArithExtension("div", func(vs ...Int) (Int, error) {
	if len(vs) == 0 {
		return Int{}, errors.New("div requires at least 1 argument")
	}
	output := vs[0]
	for i := 1; i < len(vs); i++ {
		v := vs[i]
		output.Val /= v.Val
	}
	return output, nil
})

var modExtension = makeArithExtension("mod", func(vs ...Int) (Int, error) {
	if len(vs) == 0 {
		return Int{}, errors.New("mod requires at least 1 argument")
	}
	output := vs[0]
	for i := 1; i < len(vs); i++ {
		v := vs[i]
		output.Val %= v.Val
	}
	return output, nil
})
