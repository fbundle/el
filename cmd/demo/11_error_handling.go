package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Error handling and edge cases demonstration
func main() {
	program := `
(let
	_ (print "=== Error Handling and Edge Cases ===")
	
	# Safe division
	safe_divide (lambda a b (match {b == 0}
		true "Error: division by zero"
		{a / b}
	))
	
	_ (print "safe_divide(10, 2) =" (safe_divide 10 2))
	_ (print "safe_divide(10, 0) =" (safe_divide 10 0))
	_ (print "safe_divide(15, 3) =" (safe_divide 15 3))
	
	# Safe list access
	safe_get (lambda l i (match (type l)
		"list" (match (type i)
			"int" (match {i >= 0}
				true (match {i < (len l)}
					true (get l i)
					"Error: index out of bounds"
				)
				"Error: negative index"
			)
			"Error: index must be integer"
		)
		"Error: not a list"
	))
	
	_ (print "\n=== Safe List Access ===")
	_ (print "safe_get([1 2 3], 1) =" (safe_get [1 2 3] 1))
	_ (print "safe_get([1 2 3], 5) =" (safe_get [1 2 3] 5))
	_ (print "safe_get([1 2 3], -1) =" (safe_get [1 2 3] -1))
	_ (print "safe_get(42, 1) =" (safe_get 42 1))
	_ (print "safe_get([1 2 3], \"hello\") =" (safe_get [1 2 3] "hello"))
	
	# Safe head operation
	safe_head (lambda l (match (type l)
		"list" (match (len l)
			0 "Error: empty list"
			(head l)
		)
		"Error: not a list"
	))
	
	_ (print "\n=== Safe Head Operation ===")
	_ (print "safe_head([1 2 3]) =" (safe_head [1 2 3]))
	_ (print "safe_head([]) =" (safe_head []))
	_ (print "safe_head(42) =" (safe_head 42))
	
	# Type validation
	validate_type (lambda value expected_type (match {type value == expected_type}
		true value
		"Error: type mismatch"
	))
	
	_ (print "\n=== Type Validation ===")
	_ (print "validate_type(42, \"int\") =" (validate_type 42 "int"))
	_ (print "validate_type(\"hello\", \"string\") =" (validate_type "hello" "string"))
	_ (print "validate_type(42, \"string\") =" (validate_type 42 "string"))
	_ (print "validate_type([1 2 3], \"list\") =" (validate_type [1 2 3] "list"))
	
	# Error propagation
	_ (print "\n=== Error Propagation ===")
	
	# Function that might fail
	risky_operation (lambda x (match {x < 0}
		true "Error: negative input"
		{x * x}
	))
	
	# Function that uses risky operation
	use_risky (lambda x (let
		result (risky_operation x)
		(match (type result)
			"string" result  # Error case
			{result + 10}    # Success case
		)
	))
	
	_ (print "use_risky(5) =" (use_risky 5))
	_ (print "use_risky(-3) =" (use_risky -3))
	
	# Try-catch simulation
	_ (print "\n=== Try-Catch Simulation ===")
	
	try_catch (lambda try_func catch_func (let
		result (try_func)
		(match (type result)
			"string" (catch_func result)  # Error case
			result                        # Success case
		)
	))
	
	# Example usage
	dangerous_divide (lambda a b (match {b == 0}
		true "Division by zero"
		{a / b}
	))
	
	handle_error (lambda error_msg (list "Caught error:" error_msg))
	
	_ (print "try_catch((lambda () (dangerous_divide 10 2)), handle_error) =" 
		(try_catch (lambda () (dangerous_divide 10 2)) handle_error))
	_ (print "try_catch((lambda () (dangerous_divide 10 0)), handle_error) =" 
		(try_catch (lambda () (dangerous_divide 10 0)) handle_error))
	
	# Input validation
	_ (print "\n=== Input Validation ===")
	
	validate_positive (lambda x (match (type x)
		"int" (match {x > 0}
			true x
			"Error: must be positive"
		)
		"Error: must be integer"
	))
	
	validate_range (lambda x min max (match (type x)
		"int" (match {x >= min}
			true (match {x <= max}
				true x
				"Error: too large"
			)
			"Error: too small"
		)
		"Error: must be integer"
	))
	
	_ (print "validate_positive(5) =" (validate_positive 5))
	_ (print "validate_positive(-3) =" (validate_positive -3))
	_ (print "validate_positive(\"hello\") =" (validate_positive "hello"))
	_ (print "validate_range(5, 1, 10) =" (validate_range 5 1 10))
	_ (print "validate_range(15, 1, 10) =" (validate_range 15 1 10))
	_ (print "validate_range(0, 1, 10) =" (validate_range 0 1 10))
	
	# Edge cases for arithmetic
	_ (print "\n=== Arithmetic Edge Cases ===")
	
	# Overflow simulation (using large numbers)
	large_add (lambda a b (match {a > 1000}
		true "Error: overflow"
		(match {b > 1000}
			true "Error: overflow"
			{a + b}
		)
	))
	
	_ (print "large_add(500, 300) =" (large_add 500 300))
	_ (print "large_add(1500, 300) =" (large_add 1500 300))
	_ (print "large_add(500, 1500) =" (large_add 500 1500))
	
	# Edge cases for list operations
	_ (print "\n=== List Edge Cases ===")
	
	# Safe list operations
	safe_slice (lambda l indices (match (type l)
		"list" (match (type indices)
			"list" (let
				# Validate all indices
				validate_indices (lambda indices (match (len indices)
					0 true
					(let
						first (head indices)
						rest (rest indices)
						(match (type first)
							"int" (match {first >= 0}
								true (match {first < (len l)}
									true (validate_indices rest)
									false
								)
								false
							)
							false
						)
					)
				))
				(match (validate_indices indices)
					true (slice l indices)
					"Error: invalid indices"
				)
			)
			"Error: indices must be list"
		)
		"Error: not a list"
	))
	
	_ (print "safe_slice([1 2 3 4 5], [0 2 4]) =" (safe_slice [1 2 3 4 5] [0 2 4]))
	_ (print "safe_slice([1 2 3 4 5], [0 10]) =" (safe_slice [1 2 3 4 5] [0 10]))
	_ (print "safe_slice([1 2 3 4 5], [0 -1]) =" (safe_slice [1 2 3 4 5] [0 -1]))
	_ (print "safe_slice(42, [0 1]) =" (safe_slice 42 [0 1]))
	_ (print "safe_slice([1 2 3], \"hello\") =" (safe_slice [1 2 3] "hello"))
	
	# Edge cases for functions
	_ (print "\n=== Function Edge Cases ===")
	
	# Function that handles wrong number of arguments
	flexible_add (lambda args (match (len args)
		0 "Error: no arguments"
		1 (get args 0)
		2 (let
			a (get args 0)
			b (get args 1)
			{a + b}
		)
		"Error: too many arguments"
	))
	
	_ (print "flexible_add([]) =" (flexible_add []))
	_ (print "flexible_add([5]) =" (flexible_add [5]))
	_ (print "flexible_add([3 7]) =" (flexible_add [3 7]))
	_ (print "flexible_add([1 2 3]) =" (flexible_add [1 2 3]))
	
	# Edge cases for matching
	_ (print "\n=== Matching Edge Cases ===")
	
	# Match with no default case
	strict_match (lambda x (match x
		1 "one"
		2 "two"
		3 "three"
		"Error: unexpected value"
	))
	
	_ (print "strict_match(1) =" (strict_match 1))
	_ (print "strict_match(2) =" (strict_match 2))
	_ (print "strict_match(5) =" (strict_match 5))
	_ (print "strict_match(\"hello\") =" (strict_match "hello"))
	
	# Edge cases for unwrapping
	_ (print "\n=== Unwrapping Edge Cases ===")
	
	# Safe unwrapping
	safe_unwrap (lambda l (match (type l)
		"list" (match (len l)
			0 "Error: empty list"
			* l
		)
		"Error: not a list"
	))
	
	_ (print "safe_unwrap([1 2 3]) =" (safe_unwrap [1 2 3]))
	_ (print "safe_unwrap([]) =" (safe_unwrap []))
	_ (print "safe_unwrap(42) =" (safe_unwrap 42))
	
	# Edge cases for recursion
	_ (print "\n=== Recursion Edge Cases ===")
	
	# Recursive function with depth limit
	limited_factorial (lambda n (let
		helper (lambda n depth (match {depth > 100}
			true "Error: recursion too deep"
			(match {n <= 1}
				true 1
				(helper {n - 1} {depth + 1})
			)
		))
		(helper n 0)
	))
	
	_ (print "limited_factorial(5) =" (limited_factorial 5))
	_ (print "limited_factorial(10) =" (limited_factorial 10))
	
	# Edge cases for type system
	_ (print "\n=== Type System Edge Cases ===")
	
	# Type checking edge cases
	check_type_hierarchy (lambda x (let
		t1 (type x)
		t2 (type t1)
		t3 (type t2)
		(list t1 t2 t3)
	))
	
	_ (print "check_type_hierarchy(42) =" (check_type_hierarchy 42))
	_ (print "check_type_hierarchy(\"hello\") =" (check_type_hierarchy "hello"))
	_ (print "check_type_hierarchy(nil) =" (check_type_hierarchy nil))
	
	# Error recovery
	_ (print "\n=== Error Recovery ===")
	
	# Function that tries multiple approaches
	robust_operation (lambda x (let
		# Try first approach
		approach1 (match (type x)
			"int" {x * 2}
			"Error: not integer"
		)
		# If first approach fails, try second
		approach2 (match (type approach1)
			"string" (match (type x)
				"string" (list x x)
				"Error: cannot handle type"
			)
			approach1
		)
		approach2
	))
	
	_ (print "robust_operation(5) =" (robust_operation 5))
	_ (print "robust_operation(\"hello\") =" (robust_operation "hello"))
	_ (print "robust_operation([1 2 3]) =" (robust_operation [1 2 3]))
	
	# Comprehensive error handling example
	_ (print "\n=== Comprehensive Error Handling ===")
	
	# Calculator with full error handling
	calculator (lambda op a b (let
		# Validate inputs
		valid_a (match (type a)
			"int" a
			"Error: first argument must be integer"
		)
		valid_b (match (type b)
			"int" b
			"Error: second argument must be integer"
		)
		# Check if validation passed
		(match (type valid_a)
			"string" valid_a
			(match (type valid_b)
				"string" valid_b
				# Perform operation
				(match op
					+ {valid_a + valid_b}
					- {valid_a - valid_b}
					* {valid_a * valid_b}
					/ (match {valid_b == 0}
						true "Error: division by zero"
						{valid_a / valid_b}
					)
					"Error: unknown operation"
				)
			)
		)
	))
	
	_ (print "calculator(+, 5, 3) =" (calculator + 5 3))
	_ (print "calculator(-, 10, 4) =" (calculator - 10 4))
	_ (print "calculator(*, 6, 7) =" (calculator * 6 7))
	_ (print "calculator(/, 15, 3) =" (calculator / 15 3))
	_ (print "calculator(/, 10, 0) =" (calculator / 10 0))
	_ (print "calculator(+, 5, \"hello\") =" (calculator + 5 "hello"))
	_ (print "calculator(\"unknown\", 5, 3) =" (calculator "unknown" 5 3))
	
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
