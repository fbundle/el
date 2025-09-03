package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Argument unwrapping demonstration
func main() {
	program := `
(let
	_ (print "=== Basic Unwrapping ===")
	
	# Simple unwrapping
	numbers [1 2 3 4 5]
	_ (print "Original list:" numbers)
	_ (print "Unwrapped:" *numbers)
	
	# Unwrapping in function calls
	_ (print "\n=== Unwrapping in Function Calls ===")
	_ (print "add(*[1 2 3]) =" (add *[1 2 3]))
	_ (print "mul(*[2 3 4]) =" (mul *[2 3 4]))
	_ (print "print(*[\"hello\" \"world\"]) =" (print *["hello" "world"]))
	
	# Mixed arguments with unwrapping
	_ (print "\n=== Mixed Arguments ===")
	_ (print "add(10, *[20 30]) =" (add 10 *[20 30]))
	_ (print "print(\"Numbers:\", *[1 2 3]) =" (print "Numbers:" *[1 2 3]))
	
	# Unwrapping empty lists
	_ (print "\n=== Empty List Unwrapping ===")
	_ (print "add(*[]) =" (add *[]))
	_ (print "print(*[]) =" (print *[]))
	
	# Nested unwrapping
	_ (print "\n=== Nested Unwrapping ===")
	nested [[1 2] [3 4] [5 6]]
	_ (print "Nested list:" nested)
	_ (print "First inner list:" (head nested))
	_ (print "Unwrap first inner list:" *(head nested))
	
	# Unwrapping in list construction
	_ (print "\n=== Unwrapping in List Construction ===")
	list1 [1 2 3]
	list2 [4 5 6]
	_ (print "Concatenate lists:" (list *list1 *list2))
	_ (print "Add element to list:" (list 0 *list1))
	_ (print "Add element after list:" (list *list1 7))
	
	# Unwrapping with different functions
	_ (print "\n=== Unwrapping with Different Functions ===")
	coords [10 20 30]
	_ (print "Coordinates:" coords)
	_ (print "Sum of coordinates:" (add *coords))
	_ (print "Product of coordinates:" (mul *coords))
	_ (print "Max coordinate:" (max *coords))
	
	# Unwrapping in custom functions
	_ (print "\n=== Unwrapping in Custom Functions ===")
	
	# Function that takes variable arguments
	sum_all (lambda args (match (len args)
		0 0
		(let
			first (head args)
			rest (rest args)
			{first + (sum_all rest)}
		)
	))
	
	_ (print "sum_all(*[1 2 3 4 5]) =" (sum_all *[1 2 3 4 5]))
	_ (print "sum_all(*[10 20]) =" (sum_all *[10 20]))
	
	# Function that processes multiple values
	process_values (lambda a b c (list a b c {a + b + c}))
	values [100 200 300]
	_ (print "process_values(*values) =" (process_values *values))
	
	# Unwrapping in higher-order functions
	_ (print "\n=== Unwrapping in Higher-Order Functions ===")
	
	# Apply function to list elements
	apply_to_list (lambda f l (match (len l)
		0 []
		(let
			first (head l)
			rest (rest l)
			first_result (f first)
			rest_results (apply_to_list f rest)
			(list first_result *rest_results)
		)
	))
	
	square (lambda x {x * x})
	numbers [1 2 3 4 5]
	_ (print "apply_to_list(square, numbers) =" (apply_to_list square numbers))
	
	# Unwrapping for function composition
	_ (print "\n=== Unwrapping for Function Composition ===")
	compose (lambda f g (lambda x (f (g x))))
	double (lambda x {x * 2})
	increment (lambda x {x + 1})
	
	# Create a list of functions and apply them
	funcs [double increment square]
	value 3
	
	# Apply all functions in sequence
	apply_sequence (lambda funcs value (match (len funcs)
		0 value
		(let
			first_func (head funcs)
			rest_funcs (rest funcs)
			new_value (first_func value)
			(apply_sequence rest_funcs new_value)
		)
	))
	
	_ (print "apply_sequence([double increment square], 3) =" (apply_sequence funcs value))
	
	# Unwrapping for data transformation
	_ (print "\n=== Unwrapping for Data Transformation ===")
	
	# Transform list of pairs into separate lists
	transform_pairs (lambda pairs (let
		extract_first (lambda pairs (match (len pairs)
			0 []
			(let
				first_pair (head pairs)
				rest_pairs (rest pairs)
				first_element (head first_pair)
				rest_elements (extract_first rest_pairs)
				(list first_element *rest_elements)
			)
		))
		extract_second (lambda pairs (match (len pairs)
			0 []
			(let
				first_pair (head pairs)
				rest_pairs (rest pairs)
				second_element (get first_pair 1)
				rest_elements (extract_second rest_pairs)
				(list second_element *rest_elements)
			)
		))
		(list (extract_first pairs) (extract_second pairs))
	))
	
	pairs [[1 "a"] [2 "b"] [3 "c"] [4 "d"]]
	_ (print "Pairs:" pairs)
	_ (print "Transformed:" (transform_pairs pairs))
	
	# Unwrapping for mathematical operations
	_ (print "\n=== Unwrapping for Mathematical Operations ===")
	
	# Calculate statistics
	calculate_stats (lambda numbers (let
		count (len numbers)
		sum (add *numbers)
		mean {sum / count}
		(list count sum mean)
	))
	
	stats_data [10 20 30 40 50]
	_ (print "Data:" stats_data)
	_ (print "Stats (count, sum, mean):" (calculate_stats stats_data))
	
	# Unwrapping for string operations
	_ (print "\n=== Unwrapping for String Operations ===")
	
	# Join strings (simulated with print)
	join_strings (lambda strings (print *strings))
	words ["Hello" "beautiful" "world"]
	_ (print "Join words:")
	_ (join_strings words)
	
	# Unwrapping in conditional expressions
	_ (print "\n=== Unwrapping in Conditionals ===")
	
	# Check if all numbers in list are positive
	all_positive (lambda numbers (match (len numbers)
		0 true
		(let
			first (head numbers)
			rest (rest numbers)
			(match {first > 0}
				true (all_positive rest)
				false
			)
		)
	))
	
	positive_list [1 2 3 4 5]
	negative_list [1 -2 3 4 5]
	_ (print "all_positive([1 2 3 4 5]) =" (all_positive positive_list))
	_ (print "all_positive([1 -2 3 4 5]) =" (all_positive negative_list))
	
	# Complex unwrapping example
	_ (print "\n=== Complex Unwrapping Example ===")
	
	# Matrix operations (represented as list of lists)
	matrix [[1 2 3] [4 5 6] [7 8 9]]
	
	# Get first row
	first_row (head matrix)
	_ (print "Matrix:" matrix)
	_ (print "First row:" first_row)
	_ (print "Sum of first row:" (add *first_row))
	
	# Get first column (extract first element from each row)
	get_column (lambda matrix col_index (match (len matrix)
		0 []
		(let
			first_row (head matrix)
			rest_matrix (rest matrix)
			element (get first_row col_index)
			rest_elements (get_column rest_matrix col_index)
			(list element *rest_elements)
		)
	))
	
	first_column (get_column matrix 0)
	_ (print "First column:" first_column)
	_ (print "Sum of first column:" (add *first_column))
	
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
