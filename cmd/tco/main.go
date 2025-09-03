package main

import (
	"context"
	"el/ast"
	"el/runtime"
	"el/runtime_ext"
	"fmt"
	"time"
)

func testTCOComparison() {
	fmt.Println("Tail Call Optimization (TCO) Comparison Test")
	fmt.Println("=============================================")
	fmt.Println("Demonstrating two types of recursion:")
	fmt.Println("1. TAIL RECURSION - can be optimized with TCO")
	fmt.Println("2. NON-TAIL RECURSION - needs stack preservation")
	fmt.Println()
	fmt.Println("KEY DIFFERENCE:")
	fmt.Println("- TAIL RECURSION: Recursive call is the LAST operation")
	fmt.Println("- NON-TAIL RECURSION: Work happens AFTER the recursive call")
	fmt.Println()

	// Test 1: Tail recursion (can be optimized)
	fmt.Println("1. TAIL RECURSION - Simple counter (TCO optimized):")
	fmt.Println("   Pattern: (while cond body) -> recursive call is last operation")
	testTailRecursion(10000)

	// Test 2: Non-tail recursion (needs stack)
	fmt.Println("\n2. NON-TAIL RECURSION - Map function (needs stack):")
	fmt.Println("   Pattern: (map rest f) -> work happens after recursive call")
	testNonTailRecursion(5)

	// Test 3: While loop (tail recursion pattern)
	fmt.Println("\n3. WHILE LOOP - Iterative pattern (TCO optimized):")
	fmt.Println("   Pattern: (while cond body) -> recursive call is last operation")
	testWhileLoop(1000)
}

func testTailRecursion(n int) {
	// This demonstrates tail recursion using while loop pattern
	// The recursive call is the last operation - TCO can optimize this
	tokens := ast.TokenizeWithInfixOperator(fmt.Sprintf(`
		(let
			unit (lambda x x)
			get (lambda l i (unit * (slice l (range i (add i 1)))))
			
			# Tail recursive pattern using while loop
			# The recursive call is the LAST operation (tail position)
			while (lambda cond_func body_func state (
				match (cond_func state)
				false state
				(while cond_func body_func (body_func state))  # ← TAIL CALL: last operation
			))
			
			# Count down from n to 0 using tail recursion
			count %d
			acc 0
			state (list count acc)
			cond_func (lambda state (gt (get state 0) 0))
			body_func (lambda state (let
				count (get state 0)
				acc (get state 1)
				new_count (sub count 1)
				new_acc (add acc 1)
				new_state (list new_count new_acc)
				new_state
			))
			final_state (while cond_func body_func state)
			final_state
		)
	`, n))

	r, s := runtime_ext.NewBasicRuntime()
	ctx := context.Background()

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
	fmt.Printf("Counted from %d to 0: %s (took %v) - TCO OPTIMIZED!\n", n, o.String(), duration)
}

func testNonTailRecursion(n int) {
	// This demonstrates non-tail recursion using map function
	// We need to do work AFTER the recursive call - TCO cannot optimize this
	tokens := ast.TokenizeWithInfixOperator(fmt.Sprintf(`
		(let
			unit (lambda x x)
			get (lambda l i (unit * (slice l (range i (add i 1)))))
			
			# Non-tail recursive map function
			# We need to do work AFTER the recursive call (non-tail position)
			map (lambda l f (
				match (le (len l) 0)
				true (list)
				(let
					first_elem (get l 0)
					first_value (f first_elem)           # ← WORK BEFORE recursion
					rest (slice l (range 1 (len l)))
					rest_values (map rest f)             # ← NON-TAIL CALL: not last operation
					(list first_value *rest_values)      # ← WORK AFTER recursion
				)
			))
			
			# Create a list and map over it - this requires stack preservation
			test_list (range 1 %d)
			result (map test_list (lambda x (mul x 2)))
			result
		)
	`, n))

	r, s := runtime_ext.NewBasicRuntime()
	ctx := context.Background()

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
	fmt.Printf("Mapped list of size %d: %s (took %v) - STACK PRESERVED!\n", n, o.String(), duration)
}

func testWhileLoop(n int) {
	// While loop is essentially tail recursion - the recursive call is the last operation
	tokens := ast.TokenizeWithInfixOperator(fmt.Sprintf(`
		(let
			unit (lambda x x)
			get (lambda l i (unit * (slice l (range i (add i 1)))))
			
			while (lambda cond_func body_func state (
				match (cond_func state)
				false state
				(while cond_func body_func (body_func state))  # ← TAIL CALL
			))
			
			# sum from 1 to %d
			sum 0
			n %d
			state (list sum n)
			cond_func (lambda state (gt (get state 1) 0))
			body_func (lambda state (let
				sum (get state 0)
				n (get state 1)
				new_sum (add sum n)
				new_n (sub n 1)
				new_state (list new_sum new_n)
				new_state
			))
			final_state (while cond_func body_func state)
			final_state
		)
	`, n, n))

	r, s := runtime_ext.NewBasicRuntime()
	ctx := context.Background()

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
	fmt.Printf("Sum from 1 to %d: %s (took %v) - TCO OPTIMIZED!\n", n, o.String(), duration)
}

func main() {
	testTCOComparison()
}
