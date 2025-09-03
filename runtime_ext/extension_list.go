package runtime_ext

import (
	"context"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var listExtension = Extension{
	Name: "list",
	Man:  "module: (list 1 2 (lambda x (add x 1))) - make a list",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		l := List{}
		for _, v := range values {
			l = List{l.Ins(l.Len(), v)}
		}
		return value(l)
	},
}

var lenExtension = Extension{
	Name: "len",
	Man:  "module: (len (list 1 2 3)) - get the length of a list",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
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
	Man:  "module: (get (list 1 2 3) (list 0 2)) - get the 0th and 2nd element of a list",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
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

var rangeExtension = Extension{
	Name: "range",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		if len(values) != 2 {
			return errValueString("range requires 2 arguments")
		}
		var i, j Int
		if ok := adt.Cast[Int](values[0]).Unwrap(&i); !ok {
			return errValueString("range beg must be an integer")
		}
		if ok := adt.Cast[Int](values[1]).Unwrap(&j); !ok {
			return errValueString("range end must be an integer")
		}
		output := List{}
		for k := i.int; k < j.int; k++ {
			output = List{output.Ins(output.Len(), Int{k})}
		}
		return value(output)
	},
	Man: "module: (range m n) - make a list of integers from m to n-1",
}
