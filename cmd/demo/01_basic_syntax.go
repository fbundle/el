package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Basic syntax demonstration - shows fundamental language constructs
func main() {
	program := `
(let
	# Basic literals
	_ (print "=== Basic Literals ===")
	_ (print "Number:" 42)
	_ (print "String:" "hello world")
	_ (print "List:" [1 2 3])
	_ (print "Nil:" nil)
	_ (print "Empty list:" [])
	
	# Variable binding
	_ (print "\n=== Variable Binding ===")
	x 42
	y "hello"
	z [1 2 3]
	_ (print "x =" x)
	_ (print "y =" y)
	_ (print "z =" z)
	
	# Nested let expressions
	_ (print "\n=== Nested Let ===")
	result (let
		inner 100
		outer 200
		{inner + outer}
	)
	_ (print "Nested calculation:" result)
	
	# Function definition and call
	_ (print "\n=== Functions ===")
	square (lambda x {x * x})
	_ (print "square(5) =" (square 5))
	_ (print "square(10) =" (square 10))
	
	# Multiple parameters
	add_three (lambda a b c {a + b + c})
	_ (print "add_three(1, 2, 3) =" (add_three 1 2 3))
	
	# Infix expressions
	_ (print "\n=== Infix Expressions ===")
	_ (print "1 + 2 * 3 =" {1 + 2 * 3})
	_ (print "(1 + 2) * 3 =" {(1 + 2) * 3})
	_ (print "10 - 3 - 2 =" {10 - 3 - 2})
	_ (print "2 * 3 * 4 =" {2 * 3 * 4})
	
	# Comments
	_ (print "\n=== Comments Work ===")
	# This is a comment
	commented 99  # This is also a comment
	_ (print "Commented value:" commented)
	
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
