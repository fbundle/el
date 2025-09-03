package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
	"fmt"
)

func testRuntime() {
	tokens := parser.TokenizeWithListAndInfix(`
		# simple example
		(let
			y 20
			add20 (lambda x (add x y))			# value of y is captured by add20
			y 40								# add20 takes the old value of y not new value of y
			(list (add20 10) y)
		)
		
		# list unwrapping
		(add *(list 1 2 3) 4 *(list 5 6))				# equivalent to (add 1 2 3 4 5 6)
		(add * * (list (list 2 3)) 4 *(list 5 6))		# nested unwrap

		# function call
		((lambda x (add x 1)) 20)


		# list
		(let
			l (list 3 4 5 6 7)
			length (len l)
			sublist (slice l (range 2 (len l)))
			unit (lambda x x)											# unit is the identity function
			get (lambda l i (unit * (slice l (range i (add i 1))))) 	# define get function from unit, slice, range
			second (get l 1)
			
			# implement map -> hence for loop
			map (lambda l f (match (le (len l) 0)
				true (list) 											# if l is empty then return empty list
				(let													# otherwise
					first_elem (get l 0)
					first_value (f first_elem)
					rest (slice l (range 1 (len l)))
					rest_values (map rest f)
					(list first_value *rest_values)
				)
																		
			))

			another (map (range 5 10) (lambda x [x mul 2]))				
			
			(list length sublist second another)
		)

		# does not have deep_recursion here 
		(let
			count (lambda n (match (le n 0)
				true 0 								# if n <= 0 then 0
				(add n (count (sub n 1)))		# else n + count(n-1)
			))
			(count 200)
		)

		# test recursion
		(let
			fib (lambda x (match (le x 1)
				true x 								# if x <= 1 then x
				(let							
					y (fib (sub x 1))				# else y = fib(x-1), z = fib(x-2), y + z
					z (fib (sub x 2))
					(add y z)
				)
			))
			(fib 20)
		)

		# test mutual recursion
		(let
			even (lambda n (match (le n 0)
				true true 							# if n <= 0 then true
				(odd (sub n 1))				# else odd(n-1)
			))
			odd (lambda n (match (le n 0)
				true false 							# if n <= 0 then false
				(even (sub n 1))				# else even(n-1)
			))
			[ (odd 10) (even 10) (odd 11) (even 11) (odd 12) (even 12) ]
		)

		# syntactic sugar for list
		[1 2 3 4 5]

		# rewrite fib
		(let											# { is the same as (let, } is the same as )
			+ add - sub x mul / div % mod			# short hand for common operator
			== eq != ne <= le < lt > gt >= ge

			fib (lambda n (match {n <= 1}			# {n <= 1} is the same as (<= n 1)
				true n 								# if n <= 1 then n
				(let								    # else p = fib(n-1), q = fib(n-2), p + q
					p (fib {n - 1})
					q (fib {n - 2})
					{p + q}
				)
			))
			(fib 20)
		)

		# more infix operator
		(let
			+ add - sub x mul / div % mod			# short hand for common operator
			== eq != ne <= le < lt > gt >= ge
			{1 + 2 - 3 + -4}
		)

		# map filter reduce

		# while loop
		(let
			unit (lambda x x)											# unit is the identity function
			get (lambda l i (unit * (slice l (range i (add i 1))))) 	# define get function from unit, slice, range

			while (lambda cond_func body_func state (
				match (cond_func state)
				false	state
				 		(while cond_func body_func (body_func state))	# does not have TCO here 
					
			))
			
			# sum from 1 to 200 										# sum, n = 0, 200; while n > 0: sum = sum + n; n = n - 1
			sum	0
			n	200
			state (list sum n)
			cond_func (lambda state (gt (get state 1) 0)) 				# keep looping while n > 0
			body_func (lambda state (let
				sum	(get state 0)
				n	(get state 1)
				new_sum (add sum n)
				new_n (sub n 1)
				new_state (list new_sum new_n)
				new_state
			))
			final_state (while  cond_func body_func state)
			final_state
		)

		# if else
		(let
			if (lambda cond true_value false_value (
				match cond
				true  true_value
						false_value
			))

			+ add - sub x mul / div % mod			# short hand for common operator
			== eq != ne <= le < lt > gt >= ge

			x 2
			y 4

			[ (if [x < 3] "x less than 3" "x bigger than 3") (if [y < 3] "y less than 3" "y bigger than 3") ]
		)

		# type system
		(let
			x 1
			y (list 1 2 3)
			z (lambda x {x+y})
			x1 (type x)
			y1 (type y)
			z1 (type z)
			x2 (type (type x))
			x3 (type (type (type x)))
			t (type type)

			[x y z x1 y1 z1 x2 x3 t]
		)
		
		# weird implementation
		(let
			f (lambda x (add x y))
			y 2
			(f 3)		# this works since function call uses the current frame
		)

		# empty expression -> nil
		(list () nil)
	
		# 
		list
	`)

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
