package runtime_ext

import (
	"context"
	"errors"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var printExtension = Extension{
	Name: "print",
	Man:  "{builtin: (print 1 2 (lambda x (add x 1))) - print}",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		for i, v := range values {
			fmt.Print(v)
			if i < len(values)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		return value(nil)
	},
}

var typeExtension = Extension{
	Name: "type",
	Man:  "{builtin: (print 1 2 (lambda x (add x 1))) - get type in string}",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		return errValue(errors.New("not implemented"))
		/*
			- add Type() into Object interface

			1 						-> "object" : integer
			(lambda x y {x + y}) 	-> "func2" : function of 2 arguments
		*/
	},
}
