package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Function definition and closure demonstration
func main() {
	program := `
(let
	_ (print "=== Basic Functions ===")
	
	# Simple functions
	identity (lambda x x)
	double (lambda x {x * 2})
	square (lambda x {x * x})
	
	_ (print "identity(42) =" (identity 42))
	_ (print "double(21) =" (double 21))
	_ (print "square(5) =" (square 5))
	
	# Multi-parameter functions
	_ (print "\n=== Multi-parameter Functions ===")
	add (lambda a b {a + b})
	multiply (lambda a b {a * b})
	max_two (lambda a b (match {a > b} true a b))
	
	_ (print "add(10, 20) =" (add 10 20))
	_ (print "multiply(3, 7) =" (multiply 3 7))
	_ (print "max_two(15, 8) =" (max_two 15 8))
	
	# Variable-arity functions
	_ (print "\n=== Variable-arity Functions ===")
	sum_all (lambda args (match (len args)
		0 0
		(let
			first (head args)
			rest (rest args)
			{first + (sum_all rest)}
		)
	))
	
	_ (print "sum_all([1 2 3 4 5]) =" (sum_all [1 2 3 4 5]))
	_ (print "sum_all([10 20]) =" (sum_all [10 20]))
	_ (print "sum_all([]) =" (sum_all []))
	
	# Higher-order functions
	_ (print "\n=== Higher-order Functions ===")
	apply_twice (lambda f x (f (f x)))
	compose (lambda f g (lambda x (f (g x))))
	
	_ (print "apply_twice(double, 5) =" (apply_twice double 5))
	_ (print "apply_twice(square, 3) =" (apply_twice square 3))
	
	double_then_square (compose square double)
	_ (print "double_then_square(3) =" (double_then_square 3))
	
	# Closures and lexical scoping
	_ (print "\n=== Closures and Lexical Scoping ===")
	
	# Counter with closure
	make_counter (lambda start (lambda (let
		current start
		(lambda (let
			old current
			current {current + 1}
			old
		))
	)))
	
	counter1 (make_counter 0)
	counter2 (make_counter 100)
	
	_ (print "counter1() =" (counter1))
	_ (print "counter1() =" (counter1))
	_ (print "counter1() =" (counter1))
	_ (print "counter2() =" (counter2))
	_ (print "counter2() =" (counter2))
	
	# Accumulator with closure
	make_accumulator (lambda initial (lambda (let
		sum initial
		(lambda value (let
			sum {sum + value}
			sum
		))
	)))
	
	acc (make_accumulator 0)
	_ (print "acc(10) =" (acc 10))
	_ (print "acc(20) =" (acc 20))
	_ (print "acc(30) =" (acc 30))
	
	# Function that returns functions
	_ (print "\n=== Functions Returning Functions ===")
	make_multiplier (lambda n (lambda x {x * n}))
	
	multiply_by_2 (make_multiplier 2)
	multiply_by_3 (make_multiplier 3)
	multiply_by_5 (make_multiplier 5)
	
	_ (print "multiply_by_2(7) =" (multiply_by_2 7))
	_ (print "multiply_by_3(7) =" (multiply_by_3 7))
	_ (print "multiply_by_5(7) =" (multiply_by_5 7))
	
	# Partial application
	_ (print "\n=== Partial Application ===")
	partial (lambda f arg1 (lambda arg2 (f arg1 arg2)))
	
	add_5 (partial add 5)
	multiply_10 (partial multiply 10)
	
	_ (print "add_5(3) =" (add_5 3))
	_ (print "add_5(7) =" (add_5 7))
	_ (print "multiply_10(4) =" (multiply_10 4))
	
	# Currying
	_ (print "\n=== Currying ===")
	curry_add (lambda a (lambda b {a + b}))
	
	_ (print "curry_add(5)(3) =" ((curry_add 5) 3))
	_ (print "curry_add(10)(7) =" ((curry_add 10) 7))
	
	# Function composition
	_ (print "\n=== Function Composition ===")
	compose2 (lambda f g (lambda x (f (g x))))
	compose3 (lambda f g h (lambda x (f (g (h x)))))
	
	# f(x) = x^2, g(x) = x + 1, h(x) = x * 2
	# h(g(f(x))) = h(g(x^2)) = h(x^2 + 1) = 2*(x^2 + 1)
	complex_func (compose3 (lambda x {x * 2}) (lambda x {x + 1}) square)
	_ (print "complex_func(3) =" (complex_func 3))  # 2*(3^2 + 1) = 2*10 = 20
	
	# Memoization (simple version)
	_ (print "\n=== Memoization ===")
	memoize (lambda f (let
		cache []
		(lambda x (let
			# Simple linear search for memoization
			find_result (lambda cache x (match (len cache)
				0 nil
				(let
					first (head cache)
					rest (rest l)
					(match {first == x}
						true first
						(find_result rest x)
					)
				)
			))
			result (find_result cache x)
			(match result
				nil (let
					computed (f x)
					cache (list *cache (list x computed))
					computed
				)
				computed
			)
		))
	))
	
	# Anonymous functions
	_ (print "\n=== Anonymous Functions ===")
	_ (print "Anonymous square:" ((lambda x {x * x}) 6))
	_ (print "Anonymous add:" ((lambda a b {a + b}) 10 15))
	
	# Functions as data
	_ (print "\n=== Functions as Data ===")
	operations [+ - * /]
	_ (print "operations:" operations)
	_ (print "First operation (add):" (head operations))
	_ (print "Using first operation:" ((head operations) 5 3))
	
	# Function that takes function as argument
	apply_to_list (lambda l f (map l f))
	_ (print "apply_to_list([1 2 3], square):" (apply_to_list [1 2 3] square))
	_ (print "apply_to_list([1 2 3], double):" (apply_to_list [1 2 3] double))
	
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
