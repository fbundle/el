package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Arithmetic operations demonstration
func main() {
	program := `
(let
	_ (print "=== Arithmetic Operations ===")
	
	# Basic arithmetic
	_ (print "Addition: 1 + 2 + 3 =" (add 1 2 3))
	_ (print "Subtraction: 10 - 3 - 2 =" (sub 10 3 2))
	_ (print "Multiplication: 2 * 3 * 4 =" (mul 2 3 4))
	_ (print "Division: 20 / 4 / 2 =" (div 20 4 2))
	_ (print "Modulo: 17 % 5 =" (mod 17 5))
	
	# Using infix syntax
	_ (print "\n=== Infix Arithmetic ===")
	_ (print "1 + 2 + 3 =" {1 + 2 + 3})
	_ (print "10 - 3 - 2 =" {10 - 3 - 2})
	_ (print "2 * 3 * 4 =" {2 * 3 * 4})
	_ (print "20 / 4 / 2 =" {20 / 4 / 2})
	_ (print "17 % 5 =" {17 % 5})
	
	# Complex expressions
	_ (print "\n=== Complex Expressions ===")
	_ (print "(1 + 2) * (3 + 4) =" {(1 + 2) * (3 + 4)})
	_ (print "2 * 3 + 4 * 5 =" {2 * 3 + 4 * 5})
	_ (print "2 * (3 + 4) * 5 =" {2 * (3 + 4) * 5})
	
	# Negative numbers
	_ (print "\n=== Negative Numbers ===")
	_ (print "-5 =" -5)
	_ (print "5 + (-3) =" {5 + (-3)})
	_ (print "5 - 3 =" {5 - 3})
	_ (print "(-5) * 2 =" {(-5) * 2})
	
	# Comparison operations
	_ (print "\n=== Comparison Operations ===")
	_ (print "5 == 5 =" (eq 5 5))
	_ (print "5 == 3 =" (eq 5 3))
	_ (print "5 != 3 =" (ne 5 3))
	_ (print "5 < 10 =" (lt 5 10))
	_ (print "5 <= 5 =" (le 5 5))
	_ (print "10 > 5 =" (gt 10 5))
	_ (print "5 >= 5 =" (ge 5 5))
	
	# Using infix comparison
	_ (print "\n=== Infix Comparison ===")
	_ (print "5 == 5 =" {5 == 5})
	_ (print "5 == 3 =" {5 == 3})
	_ (print "5 != 3 =" {5 != 3})
	_ (print "5 < 10 =" {5 < 10})
	_ (print "5 <= 5 =" {5 <= 5})
	_ (print "10 > 5 =" {10 > 5})
	_ (print "5 >= 5 =" {5 >= 5})
	
	# Arithmetic functions
	_ (print "\n=== Arithmetic Functions ===")
	abs (lambda x (match {x < 0} true {-x} x))
	_ (print "abs(-5) =" (abs -5))
	_ (print "abs(5) =" (abs 5))
	
	max (lambda a b (match {a > b} true a b))
	_ (print "max(10, 20) =" (max 10 20))
	_ (print "max(-10, -20) =" (max -10 -20))
	
	min (lambda a b (match {a < b} true a b))
	_ (print "min(10, 20) =" (min 10 20))
	_ (print "min(-10, -20) =" (min -10 -20))
	
	# Power function (using repeated multiplication)
	power (lambda base exp (match {exp <= 0}
		true 1
		{base * (power base {exp - 1})}
	))
	_ (print "2^3 =" (power 2 3))
	_ (print "3^4 =" (power 3 4))
	_ (print "5^0 =" (power 5 0))
	
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
