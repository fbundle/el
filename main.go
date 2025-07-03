package main

import (
	"context"
	"el/pkg/el"
	"fmt"
)

func testRuntime() {
	tokens := el.TokenizeWithInfixOperator(el.Transpile(`
		// simple example
		(let
			y 20
			add20 (lambda x (add x y))			// value of y is captured by add20
			y 40								// add20 takes the old value of y not new value of y
			(list (add20 10) y)
		)
		
		// list unwrapping
		(add *(list 1 2 3) 4 *(list 5 6))		// equivalent to (add 1 2 3 4 5 6)

		// list
		(let
			l (list 3 4 5 6 7)
			length (len l)
			sublist (slice l (range 2 (len l)))
			unit (lambda x x)							// unit is the identity function
			get (lambda l i (unit * (slice l (range i (add i 1))))) 	// define get function from unit, slice, range
			second (get l 1)
			(list length sublist second)
		)

		// test tco
		(let
			count (lambda n (match (le n 0)
				true 0 								// if n <= 0 then 0
				false (add n (count (sub n 1)))		// else n + count(n-1)
			))
			count 2000
		)

		// test recursion
		(let
			fib (lambda x (match (le x 1)
				true x 								// if x <= 1 then x
				false (let							// _ is the wildcard symbol
					y (fib (sub x 1))				// else y = fib(x-1), z = fib(x-2), y + z
					z (fib (sub x 2))
					(add y z)
				)
			))
			(fib 20)
		)

		// test mutual recursion
		(let
			even (lambda n (match (le n 0)
				true true 							// if n <= 0 then true
				false (odd (sub n 1))				// else odd(n-1)
			))
			odd (lambda n (match (le n 0)
				true false 							// if n <= 0 then false
				false (even (sub n 1))				// else even(n-1)
			))
			[[ (odd 10) (even 10) (odd 11) (even 11) (odd 12) (even 12) ]]
		)

		// syntactic sugar for list
		[[ 1 2 3 4 5 ]]

		// rewrite fib
		{											// { is the same as (let, } is the same as )
			+ add - sub x mul / div % mod			// short hand for common operator
			== eq != ne <= le < lt > gt >= ge

			fib (lambda n (match [n <= 1]			// [n <= 1] is the same as (<= n 1)
				true n 								// if n <= 1 then n
				false {								// else p = fib(n-1), q = fib(n-2), p + q
					p (fib [n - 1])
					q (fib [n - 2])
					[p + q]
				}
			))
			(fib 20)
		}

		// more infix operator
		{
			+ add - sub x mul / div % mod			// short hand for common operator
			== eq != ne <= le < lt > gt >= ge
			[1 + 2 - 3 + -4]
		}
	`))

	r := el.NewBasicRuntime()
	var expr el.Expr
	var err error
	ctx := context.Background()
	for len(tokens) > 0 {
		expr, tokens, err = el.ParseWithInfixOperator(tokens)
		if err != nil {
			panic(err)
		}
		o, err := r.Step(ctx, expr)
		if err != nil {
			panic(err)
		}
		fmt.Println("expr\t", expr)
		fmt.Println("output\t", o)
		fmt.Println()
	}
}

func main() {
	testRuntime()
}
