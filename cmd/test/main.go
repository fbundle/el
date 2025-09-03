package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
	"fmt"
)

var program = `
	(map (list 1 2 3) (lambda x (add x 2)))
`

func testRuntime() {
	tokens := parser.TokenizeWithListAndInfix(runtime_ext.WithTemplate(program))

	r, s := runtime_ext.NewBasicRuntime()

	var e ast.Expr
	var o runtime.Object
	var err error
	ctx := context.Background()
	for len(tokens) > 0 {
		e, tokens, err = parser.ParseWithInfix(tokens)
		if err != nil {
			panic(err)
		}
		fmt.Println("expr\t", e)
		if err := r.Step(ctx, s, e).Unwrap(&o); err != nil {
			panic(err)
		}
		fmt.Println("output\t", o.String())
		fmt.Println()
	}
}

func main() {
	testRuntime()
}
