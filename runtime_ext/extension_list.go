package runtime_ext

import (
	"context"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var listExtension = Extension{
	Name: "list",
	Man:  "[builtin: (list 1 2 (lambda x (add x 1))) - make a list]",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		l := List{}
		for _, v := range values {
			l = List{l.Ins(l.Len(), v)}
		}
		return resultTypedData(l)
	},
}

var lenExtension = Extension{
	Name: "len",
	Man:  "[builtin: (len (list 1 2 3)) - get the length of a list]",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		if len(values) != 1 {
			return resultErrStrf("len requires 1 argument")
		}
		l, ok := values[0].Data().(List)
		if !ok {
			return resultErrStrf("len argument must be a list")
		}
		return resultTypedData(Int{l.Len()})
	},
}

var sliceExtension = Extension{
	Name: "slice",
	Man:  "[builtin: (get (list 1 2 3) (list 0 2)) - get the 0th and 2nd element of a list]",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		if len(values) != 2 {
			return resultErrStrf("slice requires 2 arguments")
		}
		l, ok := values[0].Data().(List)
		if !ok {
			return resultErrStrf("slice first argument must be a list")
		}
		i, ok := values[1].Data().(List)
		if !ok {
			return resultErrStrf("slice second argument must be a list of integers")
		}
		output := List{}
		for _, o := range i.Iter {
			index, ok := o.Data().(Int)
			if !ok {
				return resultErrStrf("slice second argument must be a list of integers")
			}
			v := l.Get(index.Val)
			output = List{output.Ins(output.Len(), v)}
		}
		return resultTypedData(output)
	},
}

var rangeExtension = Extension{
	Name: "range",
	Man:  "[builtin: (range m n) - make a list of integers from m to n-1]",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		if len(values) != 2 {
			return resultErrStrf("range requires 2 arguments")
		}
		i, ok := values[0].Data().(Int)
		if !ok {
			return resultErrStrf("range beg must be an integer")
		}
		j, ok := values[1].Data().(Int)
		if !ok {
			return resultErrStrf("range end must be an integer")
		}
		output := List{}
		for k := i.Val; k < j.Val; k++ {
			output = List{output.Ins(output.Len(), makeTypedData(Int{k}))}
		}
		return resultTypedData(output)
	},
}
