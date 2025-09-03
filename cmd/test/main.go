package main

import (
	"context"
	"el/pkg/el/expr"
	"el/pkg/el/runtime_core"
	"el/pkg/el/runtime_ext"
	"fmt"
)

func testRuntime() {
	tokens := expr.TokenizeWithInfixOperator(`
		(let
			x (list 1 2 3)
			y (list 4 5 6)
			(list *x *y)
		)
	`)

	r, s := runtime_ext.NewBasicRuntime()

	var e expr.Expr
	var o runtime_core.Value
	var err error
	ctx := context.Background()
	for len(tokens) > 0 {
		e, tokens, err = expr.ParseWithInfixOperator(tokens)
		if err != nil {
			panic(err)
		}
		fmt.Println("expr\t", e)
		if err := r.Step(ctx, s, e).Unwrap(&o); err != nil {
			panic(err)
		}
		fmt.Println("output\t", o.MustString())
		fmt.Println()
	}
}

func main() {
	testRuntime()
}
