package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
	"fmt"
	"time"
)

func testSimpleTCO(n int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Simple TCO Test")
	fmt.Println("===============")

	// Simple tail recursive function
	tokens := parser.TokenizeWithInfixOperator(fmt.Sprintf(`
		(let
			# Simple tail recursive counter
			count (lambda n acc (
				match (le n 0)
				true acc
				(count (sub n 1) (add acc 1))  # ← This should be a tail call
			))
			result (count %d 0)
			result
		)
	`, n))

	r, s := runtime_ext.NewBasicRuntime()

	start := time.Now()

	var e ast.Expr
	var o runtime.Object
	var err error

	for len(tokens) > 0 {
		e, tokens, err = parser.ParseWithInfixOperator(tokens)
		if err != nil {
			panic(err)
		}

		if err := r.Step(ctx, s, e).Unwrap(&o); err != nil {
			panic(err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("Counted to %d: %s (took %v)\n", n, o.String(), duration)
}

func main() {
	testSimpleTCO(10000)
}
