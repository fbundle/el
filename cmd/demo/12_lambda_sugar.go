package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Lambda syntactic sugar demonstration
func main() {
	program := `
(let
	_ (print "=== Lambda Syntactic Sugar ===")
	
	# Basic lambda syntax
	_ (print "\n=== Basic Lambda Syntax ===")
	
	# Old syntax: (lambda x {x * 2})
	# New syntax: {x => {x * 2}}
	square_old (lambda x {x * x})
	square_new {x => {x * x}}
	
	_ (print "square_old(5) =" (square_old 5))
	_ (print "square_new(5) =" (square_new 5))
	
	# Multiple parameters
	add_old (lambda a b {a + b})
	add_new {a b => {a + b}}
	
	_ (print "add_old(3, 7) =" (add_old 3 7))
	_ (print "add_new(3, 7) =" (add_new 3 7))
	
	# Complex expressions in body
	complex_old (lambda x y (let
		sum {x + y}
		product {x * y}
		{sum + product}
	))
	complex_new {x y => (let
		sum {x + y}
		product {x * y}
		{sum + product}
	)}
	
	_ (print "complex_old(2, 3) =" (complex_old 2 3))
	_ (print "complex_new(2, 3) =" (complex_new 2 3))
	
	# No parameters
	_ (print "\n=== No Parameters ===")
	constant_old (lambda () 42)
	constant_new {() => 42}
	
	_ (print "constant_old() =" (constant_old))
	_ (print "constant_new() =" (constant_new))
	
	# Single parameter
	_ (print "\n=== Single Parameter ===")
	identity_old (lambda x x)
	identity_new {x => x}
	
	_ (print "identity_old(99) =" (identity_old 99))
	_ (print "identity_new(99) =" (identity_new 99))
	
	# Higher-order functions
	_ (print "\n=== Higher-Order Functions ===")
	
	# Function that returns a lambda
	make_multiplier_old (lambda n (lambda x {x * n}))
	make_multiplier_new (lambda n {x => {x * n}})
	
	mult_by_5_old (make_multiplier_old 5)
	mult_by_5_new (make_multiplier_new 5)
	
	_ (print "mult_by_5_old(7) =" (mult_by_5_old 7))
	_ (print "mult_by_5_new(7) =" (mult_by_5_new 7))
	
	# Function composition
	compose_old (lambda f g (lambda x (f (g x))))
	compose_new {f g => {x => (f (g x))}}
	
	double (lambda x {x * 2})
	increment (lambda x {x + 1})
	
	composed_old (compose_old double increment)
	composed_new (compose_new double increment)
	
	_ (print "composed_old(5) =" (composed_old 5))  # (5+1)*2 = 12
	_ (print "composed_new(5) =" (composed_new 5))  # (5+1)*2 = 12
	
	# List operations with lambda sugar
	_ (print "\n=== List Operations ===")
	
	numbers [1 2 3 4 5]
	
	# Map with lambda sugar
	doubled_old (map numbers (lambda x {x * 2}))
	doubled_new (map numbers {x => {x * 2}})
	
	_ (print "doubled_old:" doubled_old)
	_ (print "doubled_new:" doubled_new)
	
	# Filter with lambda sugar
	is_even_old (lambda x {x % 2 == 0})
	is_even_new {x => {x % 2 == 0}}
	
	evens_old (filter numbers is_even_old)
	evens_new (filter numbers is_even_new)
	
	_ (print "evens_old:" evens_old)
	_ (print "evens_new:" evens_new)
	
	# Reduce with lambda sugar
	sum_old (reduce numbers (lambda acc x {acc + x}) 0)
	sum_new (reduce numbers {acc x => {acc + x}} 0)
	
	_ (print "sum_old:" sum_old)
	_ (print "sum_new:" sum_new)
	
	# Nested lambdas
	_ (print "\n=== Nested Lambdas ===")
	
	# Function that takes a function and applies it twice
	apply_twice_old (lambda f (lambda x (f (f x))))
	apply_twice_new {f => {x => (f (f x))}}
	
	square (lambda x {x * x})
	
	result_old (apply_twice_old square)
	result_new (apply_twice_new square)
	
	_ (print "apply_twice_old(square)(3) =" (result_old 3))  # ((3^2)^2) = 81
	_ (print "apply_twice_new(square)(3) =" (result_new 3))  # ((3^2)^2) = 81
	
	# Currying with lambda sugar
	_ (print "\n=== Currying ===")
	
	curry_add_old (lambda a (lambda b {a + b}))
	curry_add_new {a => {b => {a + b}}}
	
	_ (print "curry_add_old(5)(3) =" ((curry_add_old 5) 3))
	_ (print "curry_add_new(5)(3) =" ((curry_add_new 5) 3))
	
	# Complex nested example
	_ (print "\n=== Complex Nested Example ===")
	
	# Function that creates a function that creates a function
	create_adder_old (lambda base (lambda step (lambda x {base + step * x})))
	create_adder_new {base => {step => {x => {base + step * x}}}}
	
	# Create a function that adds 10 + 2*x
	adder_old (create_adder_old 10)
	adder_new (create_adder_new 10)
	
	final_adder_old (adder_old 2)
	final_adder_new (adder_new 2)
	
	_ (print "final_adder_old(5) =" (final_adder_old 5))  # 10 + 2*5 = 20
	_ (print "final_adder_new(5) =" (final_adder_new 5))  # 10 + 2*5 = 20
	
	# Comparison of syntax verbosity
	_ (print "\n=== Syntax Comparison ===")
	_ (print "Old: (lambda x y (add (mul x 2) (mul y 3)))")
	_ (print "New: {x y => (add (mul x 2) (mul y 3))}")
	
	func_old (lambda x y (add (mul x 2) (mul y 3)))
	func_new {x y => (add (mul x 2) (mul y 3))}
	
	_ (print "func_old(5, 7) =" (func_old 5 7))  # 2*5 + 3*7 = 10 + 21 = 31
	_ (print "func_new(5, 7) =" (func_new 5 7))  # 2*5 + 3*7 = 10 + 21 = 31
	
	nil
)`

	runProgram(program)
}

func runProgram(program string) {
	tokens := parser.Tokenize(runtime_ext.WithTemplate(program))
	r, s := runtime_ext.NewBasicRuntime()

	ctx := context.Background()
	for len(tokens) > 0 {
		var e ast.Expr
		var err error
		e, tokens, err = parser.Parse(tokens)
		if err != nil {
			panic(err)
		}

		var o runtime.Object
		if err := r.Step(ctx, s, e).Unwrap(&o); err != nil {
			panic(err)
		}
	}
}
