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
(let
	# basic
    _ (print "hello world")
	_ (print list)
	_ (print [1 2 3 (list 4 5 6) (lambda x {x + 3})])
	_ (print (map (list 1 2 3) (lambda x {x + 2})))
	_ (print [1 2 3] $[1 2 3])				# unwrap arguments
	_ (print {1 + 2 - 3 + 4})

	# fibonacci
	fib (lambda n (match {n <= 1}
		true n 								# if n <= 1 then n
		(let								# else p = fib(n-1), q = fib(n-2), p + q
			p (fib {n - 1})
			q (fib {n - 2})
			{p + q}		
		)
	))
	_ (print "fib(20)=" (fib 20))

	# deep recursion
	count (lambda n (match {n <= 0}
		true 0 								# if n <= 0 then 0
		{1 + (count {n - 1})}				# else 1 + count(n-1)
	))
	_ (print "count(2000)=" (count 2000))

	# mutual recursion
    even (lambda n (match {n <= 0}
		true true 							# if n <= 0 then true
			 (odd {n - 1}) 					# else odd(n-1)
	))
	odd (lambda n (match {n <= 0}
		true false 							# if n <= 0 then false
			 (even {n - 1}) 					# else even(n-1)
	))
	_ (print "evens and odds:" [(odd 10) (even 10) (odd 11) (even 11) (odd 12) (even 12)])

	# type system
	_ (print "some types" (let
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
	))

	# weird implementation
	_ (print (let
			f (lambda x (add x y))
			y 2
			(f 3)		# this works since function call uses the current frame
	))

	# nil
	_ (print (list () nil))
	_ (print "type(nil)=" (type nil))
	_ (print "empty()=" ())

	# simple match sanity
	_ (print "match 1==1 ->" (match 1 1 "ok" "no"))
	_ (print "match 1==2 ->" (match 1 2 "ok" "no"))

	# arrow function 
	_ (print "arrow function:" {x y => {x + y}})

	# type check
	is_func (lambda x (match (type x)
		function true
				 false
	))

	_ (print (is_func (lambda x x)) (is_func 1))
	
	nil
)`

func main() {
	testRuntime()
}
func testRuntime() {
	tokens := parser.Tokenize(withTemplate(program))

	r, s := runtime_ext.NewBasicRuntime()

	var e ast.Expr
	var o runtime.Object
	var err error
	ctx := context.Background()
	for len(tokens) > 0 {
		e, tokens, err = parser.Parse(tokens)
		if err != nil {
			panic(err)
		}
		fmt.Println("expr\t", e)
		if err := r.Step(ctx, s, e).Unwrap(&o); err != nil {
			panic(err)
		}
		fmt.Println("output\t", o)
		fmt.Println()
	}
}

// withTemplate - add some common template to the code
func withTemplate(s string) string {
	return fmt.Sprintf(`
(let

# identity - identity function
unit (lambda x x) 

# get - get element from list
get (lambda l i (unit $(slice l (range i (add i 1)))))			# get l[i]
head (lambda l (get l 0))							# get l[0]
rest (lambda l (slice l (range 1 (len l))))			# get l[1:]

# operators
+ add - sub x mul / div %% mod			# short hand for common operator
== eq != ne <= le < lt > gt >= ge

# map
map (lambda l f (match (len l)
	0 []					# if len l == 0 then return empty list
	(let
		first_elem (head l)
		first_elem2 (f first_elem)
		rest_elems (rest l)
		rest_elems2 (map rest_elems f)	# recursive call
		(list first_elem2 $rest_elems2)
	)
))

%s

)`, s)
}
