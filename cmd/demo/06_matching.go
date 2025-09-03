package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Pattern matching demonstration
func main() {
	program := `
(let
	_ (print "=== Basic Pattern Matching ===")
	
	# Simple value matching
	_ (print "match 1 with 1:" (match 1 1 "yes" "no"))
	_ (print "match 1 with 2:" (match 1 2 "yes" "no"))
	_ (print "match 42 with 42:" (match 42 42 "found" "not found"))
	
	# Multiple patterns
	grade (lambda score (match score
		90 "A"
		80 "B"
		70 "C"
		60 "D"
		"F"
	))
	
	_ (print "\n=== Grade Matching ===")
	_ (print "grade(95) =" (grade 95))
	_ (print "grade(85) =" (grade 85))
	_ (print "grade(75) =" (grade 75))
	_ (print "grade(65) =" (grade 65))
	_ (print "grade(55) =" (grade 55))
	
	# String matching
	greeting (lambda name (match name
		"Alice" "Hello Alice!"
		"Bob" "Hi Bob!"
		"Charlie" "Hey Charlie!"
		"Hello stranger!"
	))
	
	_ (print "\n=== String Matching ===")
	_ (print "greeting(\"Alice\") =" (greeting "Alice"))
	_ (print "greeting(\"Bob\") =" (greeting "Bob"))
	_ (print "greeting(\"David\") =" (greeting "David"))
	
	# Boolean matching
	is_positive (lambda n (match {n > 0}
		true "positive"
		"non-positive"
	))
	
	_ (print "\n=== Boolean Matching ===")
	_ (print "is_positive(5) =" (is_positive 5))
	_ (print "is_positive(0) =" (is_positive 0))
	_ (print "is_positive(-3) =" (is_positive -3))
	
	# List length matching
	list_info (lambda l (match (len l)
		0 "empty"
		1 "single"
		2 "pair"
		"many"
	))
	
	_ (print "\n=== List Length Matching ===")
	_ (print "list_info([]) =" (list_info []))
	_ (print "list_info([1]) =" (list_info [1]))
	_ (print "list_info([1 2]) =" (list_info [1 2]))
	_ (print "list_info([1 2 3 4]) =" (list_info [1 2 3 4]))
	
	# Complex conditional matching
	classify_number (lambda n (match n
		0 "zero"
		(match {n > 0}
			true (match {n % 2 == 0}
				true "positive even"
				"positive odd"
			)
			(match {n % 2 == 0}
				true "negative even"
				"negative odd"
			)
		)
	))
	
	_ (print "\n=== Complex Number Classification ===")
	_ (print "classify_number(0) =" (classify_number 0))
	_ (print "classify_number(4) =" (classify_number 4))
	_ (print "classify_number(7) =" (classify_number 7))
	_ (print "classify_number(-6) =" (classify_number -6))
	_ (print "classify_number(-9) =" (classify_number -9))
	
	# Function type matching
	apply_operation (lambda op a b (match op
		+ {a + b}
		- {a - b}
		* {a * b}
		/ {a / b}
		"unknown operation"
	))
	
	_ (print "\n=== Function Matching ===")
	_ (print "apply_operation(+, 5, 3) =" (apply_operation + 5 3))
	_ (print "apply_operation(-, 10, 4) =" (apply_operation - 10 4))
	_ (print "apply_operation(*, 6, 7) =" (apply_operation * 6 7))
	
	# Nested matching
	process_data (lambda data (match (type data)
		"int" (match {data < 0}
			true "negative integer"
			"positive integer"
		)
		"string" (match (len data)
			0 "empty string"
			"non-empty string"
		)
		"list" (match (len data)
			0 "empty list"
			"non-empty list"
		)
		"unknown type"
	))
	
	_ (print "\n=== Nested Type Matching ===")
	_ (print "process_data(5) =" (process_data 5))
	_ (print "process_data(-3) =" (process_data -3))
	_ (print "process_data(\"\") =" (process_data ""))
	_ (print "process_data(\"hello\") =" (process_data "hello"))
	_ (print "process_data([]) =" (process_data []))
	_ (print "process_data([1 2]) =" (process_data [1 2]))
	
	# Pattern matching with computed values
	find_in_list (lambda l target (match (len l)
		0 false
		(let
			first (head l)
			rest (rest l)
			(match {first == target}
				true true
				(find_in_list rest target)
			)
		)
	))
	
	_ (print "\n=== List Search with Matching ===")
	_ (print "find_in_list([1 2 3 4], 3) =" (find_in_list [1 2 3 4] 3))
	_ (print "find_in_list([1 2 3 4], 5) =" (find_in_list [1 2 3 4] 5))
	_ (print "find_in_list([\"a\" \"b\" \"c\"], \"b\") =" (find_in_list ["a" "b" "c"] "b"))
	
	# Switch-like behavior
	day_of_week (lambda n (match n
		1 "Monday"
		2 "Tuesday"
		3 "Wednesday"
		4 "Thursday"
		5 "Friday"
		6 "Saturday"
		7 "Sunday"
		"Invalid day"
	))
	
	_ (print "\n=== Day of Week Matching ===")
	_ (print "day_of_week(1) =" (day_of_week 1))
	_ (print "day_of_week(3) =" (day_of_week 3))
	_ (print "day_of_week(7) =" (day_of_week 7))
	_ (print "day_of_week(8) =" (day_of_week 8))
	
	# Range matching
	age_group (lambda age (match {age < 13}
		true "child"
		(match {age < 20}
			true "teenager"
			(match {age < 65}
				true "adult"
				"senior"
			)
		)
	))
	
	_ (print "\n=== Age Group Matching ===")
	_ (print "age_group(8) =" (age_group 8))
	_ (print "age_group(16) =" (age_group 16))
	_ (print "age_group(30) =" (age_group 30))
	_ (print "age_group(70) =" (age_group 70))
	
	# Error handling with matching
	safe_divide (lambda a b (match {b == 0}
		true "Error: division by zero"
		{a / b}
	))
	
	_ (print "\n=== Safe Division with Matching ===")
	_ (print "safe_divide(10, 2) =" (safe_divide 10 2))
	_ (print "safe_divide(10, 0) =" (safe_divide 10 0))
	_ (print "safe_divide(15, 3) =" (safe_divide 15 3))
	
	# Complex data structure matching
	analyze_point (lambda point (match (len point)
		2 (let
			x (get point 0)
			y (get point 1)
			(match {x == 0}
				true (match {y == 0}
					true "origin"
					"on y-axis"
				)
				(match {y == 0}
					true "on x-axis"
					"general point"
				)
			)
		)
		"invalid point"
	))
	
	_ (print "\n=== Point Analysis Matching ===")
	_ (print "analyze_point([0 0]) =" (analyze_point [0 0]))
	_ (print "analyze_point([0 5]) =" (analyze_point [0 5]))
	_ (print "analyze_point([3 0]) =" (analyze_point [3 0]))
	_ (print "analyze_point([2 3]) =" (analyze_point [2 3]))
	_ (print "analyze_point([1]) =" (analyze_point [1]))
	
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
