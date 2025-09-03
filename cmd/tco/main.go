package main

import (
	"context"
	"el/ast"
	"el/runtime"
	"el/runtime_ext"
	"fmt"
	"time"
)

func testSimpleTCO() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Simple TCO Test")
	fmt.Println("===============")

	// Simple tail recursive function
	tokens := ast.TokenizeWithInfixOperator(`
		(let
			# Simple tail recursive counter
			count (lambda n acc (
				match (le n 0)
				true acc
				(count (sub n 1) (add acc 1))  # ← This should be a tail call
			))
			result (count 1000 0)
			result
		)
	`)

	r, s := runtime_ext.NewBasicRuntime()

	start := time.Now()

	var e ast.Expr
	var o runtime.Value
	var err error

	for len(tokens) > 0 {
		e, tokens, err = ast.ParseWithInfixOperator(tokens)
		if err != nil {
			panic(err)
		}

		if err := r.Step(ctx, s, e).Unwrap(&o); err != nil {
			panic(err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("Counted to 1000: %s (took %v)\n", o.String(), duration)
}

func main() {
	testSimpleTCO()
}
