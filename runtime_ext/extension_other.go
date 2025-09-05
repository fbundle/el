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
		return resultObj(nil)
	},
}

var typeExtension = Extension{
	Name: "type",
	Man:  "{builtin: (print 1 2 (lambda x (add x 1))) - get type in string}",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		return resultErr(errors.New("not implemented"))
		/*
			(type int int bool) or {int -> int -> bool}  			# make type
			(typeof 1) 												# int
			(typeof (lambda x y {x + y})) 							# {any -> any -> int}
			(let
				f {(lambda x y {x + y}) : {int -> int -> int} }		# decorate anything with type
				(typeof f)											# {int -> int -> int}
			)
			(f 1 2)													# type check before execution

		*/
	},
}
