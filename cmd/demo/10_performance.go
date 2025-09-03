package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Performance and benchmarking demonstration
func main() {
	program := `
(let
	_ (print "=== Performance Testing ===")
	
	# Simple timing function
	time_function (lambda name func (let
		start_time 0  # We'll measure in iterations instead
		result (func)
		(list name result)
	))
	
	# Iterative vs Recursive Fibonacci
	_ (print "\n=== Fibonacci Performance ===")
	
	# Recursive Fibonacci (exponential time)
	fib_rec (lambda n (match {n <= 1}
		true n
		{(fib_rec {n - 1}) + (fib_rec {n - 2})}
	))
	
	# Iterative Fibonacci (linear time)
	fib_iter (lambda n (let
		loop (lambda a b count (match {count <= 0}
			true a
			(loop b {a + b} {count - 1})
		))
		(loop 0 1 n)
	))
	
	# Test recursive version (small numbers only)
	_ (print "fib_rec(10) =" (fib_rec 10))
	_ (print "fib_rec(15) =" (fib_rec 15))
	
	# Test iterative version (can handle larger numbers)
	_ (print "fib_iter(10) =" (fib_iter 10))
	_ (print "fib_iter(20) =" (fib_iter 20))
	_ (print "fib_iter(30) =" (fib_iter 30))
	_ (print "fib_iter(40) =" (fib_iter 40))
	
	# List operations performance
	_ (print "\n=== List Operations Performance ===")
	
	# Generate large list
	generate_list (lambda size (let
		loop (lambda result count (match {count <= 0}
			true result
			(loop (list *result count) {count - 1})
		))
		(loop [] size)
	))
	
	small_list (generate_list 10)
	medium_list (generate_list 50)
	
	_ (print "Small list length:" (len small_list))
	_ (print "Medium list length:" (len medium_list))
	
	# Sum list elements
	sum_list (lambda l (let
		loop (lambda l acc (match (len l)
			0 acc
			(let
				first (head l)
				rest (rest l)
				(loop rest {acc + first})
			)
		))
		(loop l 0)
	))
	
	_ (print "Sum of small list:" (sum_list small_list))
	_ (print "Sum of medium list:" (sum_list medium_list))
	
	# Sorting performance
	_ (print "\n=== Sorting Performance ===")
	
	# Bubble sort (O(n^2))
	bubble_sort (lambda l (let
		swap (lambda l i j (let
			val_i (get l i)
			val_j (get l j)
			# Create new list with swapped elements
			replace_at (lambda l i val (let
				loop (lambda result remaining idx (match (len remaining)
					0 result
					(let
						first (head remaining)
						rest (rest remaining)
						new_val (match {idx == i}
							true val
							first
						)
						(loop (list *result new_val) rest {idx + 1})
					)
				))
				(loop [] l 0)
			))
			l1 (replace_at l i val_j)
			l2 (replace_at l1 j val_i)
			l2
		))
		
		# One pass of bubble sort
		bubble_pass (lambda l (let
			loop (lambda l i (match {i >= (len l) - 1}
				true l
				(let
					val_i (get l i)
					val_next (get l {i + 1})
					(match {val_i > val_next}
						true (loop (swap l i {i + 1}) {i + 1})
						(loop l {i + 1})
					)
				)
			))
			(loop l 0)
		))
		
		# Full bubble sort
		sort_loop (lambda l (let
			pass_result (bubble_pass l)
			(match {pass_result == l}
				true l
				(sort_loop pass_result)
			)
		))
		(sort_loop l)
	))
	
	unsorted [5 2 8 1 9 3 7 4 6]
	_ (print "Unsorted:" unsorted)
	_ (print "Bubble sorted:" (bubble_sort unsorted))
	
	# Memory usage simulation
	_ (print "\n=== Memory Usage Simulation ===")
	
	# Create nested structures to test memory
	make_nested (lambda depth (match {depth <= 0}
		true "leaf"
		(let
			left (make_nested {depth - 1})
			right (make_nested {depth - 1})
			(list left right)
		)
	))
	
	shallow_tree (make_nested 3)
	_ (print "Shallow tree created (depth 3)")
	
	# Count nodes in tree
	count_nodes (lambda tree (match (type tree)
		"string" 1
		"list" (match (len tree)
			0 0
			(let
				first (head tree)
				rest (rest tree)
				first_count (count_nodes first)
				rest_count (count_nodes rest)
				{first_count + rest_count}
			)
		)
		0
	))
	
	_ (print "Nodes in shallow tree:" (count_nodes shallow_tree))
	
	# Algorithm complexity demonstration
	_ (print "\n=== Algorithm Complexity ===")
	
	# O(1) - constant time
	constant_time (lambda n 42)
	
	# O(log n) - logarithmic time (simulated with repeated division)
	log_time (lambda n (let
		loop (lambda n count (match {n <= 1}
			true count
			(loop {n / 2} {count + 1})
		))
		(loop n 0)
	))
	
	# O(n) - linear time
	linear_time (lambda n (let
		loop (lambda i sum (match {i > n}
			true sum
			(loop {i + 1} {sum + i})
		))
		(loop 1 0)
	))
	
	# O(n^2) - quadratic time
	quadratic_time (lambda n (let
		outer_loop (lambda i sum (match {i > n}
			true sum
			(let
				inner_loop (lambda j inner_sum (match {j > n}
					true inner_sum
					(inner_loop {j + 1} {inner_sum + 1})
				))
				inner_result (inner_loop 1 0)
				(outer_loop {i + 1} {sum + inner_result})
			)
		))
		(outer_loop 1 0)
	))
	
	test_n 10
	_ (print "Testing with n =" test_n)
	_ (print "O(1) result:" (constant_time test_n))
	_ (print "O(log n) result:" (log_time test_n))
	_ (print "O(n) result:" (linear_time test_n))
	_ (print "O(n^2) result:" (quadratic_time test_n))
	
	# Caching and memoization performance
	_ (print "\n=== Caching Performance ===")
	
	# Simple cache implementation
	make_cache (lambda () (let
		cache []
		get (lambda key (let
			find (lambda cache key (match (len cache)
				0 nil
				(let
					entry (head cache)
					rest (rest cache)
					entry_key (get entry 0)
					entry_value (get entry 1)
					(match {entry_key == key}
						true entry_value
						(find rest key)
					)
				)
			))
			(find cache key)
		))
		set (lambda key value (let
			# Simple implementation - just append
			new_entry (list key value)
			cache (list *cache new_entry)
			value
		))
		(lambda method (match method
			"get" get
			"set" set
			"Error: unknown method"
		))
	))
	
	cache (make_cache)
	
	# Test cache performance
	_ (cache "set" "key1" 100)
	_ (cache "set" "key2" 200)
	_ (print "Cache get key1:" (cache "get" "key1"))
	_ (print "Cache get key2:" (cache "get" "key2"))
	_ (print "Cache get key3:" (cache "get" "key3"))
	
	# String operations performance
	_ (print "\n=== String Operations Performance ===")
	
	# String concatenation (simulated with lists)
	concat_strings (lambda strings (let
		loop (lambda strings result (match (len strings)
			0 result
			(let
				first (head strings)
				rest (rest strings)
				new_result (list *result first)
				(loop rest new_result)
			)
		))
		(loop strings [])
	))
	
	words ["Hello" "beautiful" "world" "of" "programming"]
	_ (print "Concatenated:" (concat_strings words))
	
	# Pattern matching performance
	_ (print "\n=== Pattern Matching Performance ===")
	
	# Complex pattern matching
	complex_match (lambda data (match (type data)
		"int" (match {data < 0}
			true (match {data % 2 == 0}
				true "negative even"
				"negative odd"
			)
			(match {data % 2 == 0}
				true "positive even"
				"positive odd"
			)
		)
		"string" (match (len data)
			0 "empty"
			1 "single char"
			"multi char"
		)
		"list" (match (len data)
			0 "empty list"
			1 "single element"
			"multiple elements"
		)
		"unknown"
	))
	
	test_values [42 -17 "hello" "" [1] [1 2 3] nil]
	_ (print "Complex matching results:")
	_ (print "42:" (complex_match 42))
	_ (print "-17:" (complex_match -17))
	_ (print "\"hello\":" (complex_match "hello"))
	_ (print "\"\" (empty):" (complex_match ""))
	_ (print "[1]:" (complex_match [1]))
	_ (print "[1 2 3]:" (complex_match [1 2 3]))
	_ (print "nil:" (complex_match nil))
	
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
