package runtime_ext

import (
	"context"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var printExtension = Extension{
	Name: "print",
	Man:  "{builtin: (print 1 2 (lambda x (add x 1))) - print}",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		for i, v := range values {
			// fmt.Printf("{%s : %s}", v, v.Type())
			fmt.Print(v)
			if i < len(values)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		return resultObj(nil)
	},
}

var inspectExtension = Extension{
	Name: "inspect",
	Man:  "{builtin: (inspect 1 2 (lambda x (add x 1))) - print object with type}",
	Exec: func(ctx context.Context, values ...Object) adt.Result[Object] {
		if len(values) < 1 {
			return resultErrStrf("inspect requires a welcoming message")
		}
		msgObj := values[0]
		fmt.Print(msgObj)
		for i := 1; i < len(values); i++ {
			v := values[i]
			fmt.Printf("{%s : %s}", v, v.Type())
			if i < len(values)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		return resultObj(nil)
	},
}
