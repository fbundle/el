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
			fib (lambda x (match (sign (add x -2))
				-1 x 								// if (x-2) < 0 then x
				_ (let								// _ is the wildcard symbol
					y (fib (add x -1))				// else y = fib(x-1), z = fib(x-2), y + z
					z (fib (add x -2))
					(add y z)
				)
			))
			(fib 20)
		)

		// test tail recursion
		(let
			count (lambda n (match n
				0 0
				_ (add 1 (count (add n -1)))
			))
			count 2000
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
	`)

	r := el.NewBasicRuntime()
	var expr el.Expr
	var err error
	ctx := context.Background()
	for len(tokens) > 0 {
		expr, tokens, err = el.Parse(tokens)
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
