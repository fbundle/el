package main

import (
	"context"
	"el/pkg/el"
	"fmt"
)

func testRuntime() {
	tokens := el.Tokenize(`
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

		// test tail recursion
		(let
			count (lambda n (match (le n 0)
				true 0
				false (add 1 (count (sub n 1)))
			))
			(list (count 2000) (count -1000))
		)

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
			get (lambda l i (unit * (slice l (range i (add i 1))))) 	// define get function from unit, slice, range
			second (get l 1)
			(list length sublist second)
		)
		// some tests from chatgpt
		(let f (lambda n (match (eq n 0) true 1 false (mul n (f (sub n 1)))))
			 (f 5))

		(let even (lambda x (match x 0 1 _ (odd (sub x 1))))
			 odd  (lambda x (match x 0 0 _ (even (sub x 1))))
			 (even 10))
		
		// inplace operator
		(let
			+ add
			- sub
			x mul
			/ div
			% mod
			[1 + 2 - 3 + -4]
		)
	`)

	r := el.NewBasicRuntime()
	var expr el.Expr
	var err error
	ctx := context.Background()
	for len(tokens) > 0 {
		expr, tokens, err = el.ParseWithInplaceOperator(tokens)
		if err != nil {
			panic(err)
		}
		o, err := r.Step(ctx, expr)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s -> %s\n", expr.String(), o.String())
	}
}

func main() {
	testRuntime()

}
