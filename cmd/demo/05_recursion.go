package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Recursion and mutual recursion demonstration
func main() {
	program := `
(let
	_ (print "=== Basic Recursion ===")
	
	# Factorial
	factorial (lambda n (match {n <= 1}
		true 1
		{n * (factorial {n - 1})}
	))
	
	_ (print "factorial(0) =" (factorial 0))
	_ (print "factorial(1) =" (factorial 1))
	_ (print "factorial(5) =" (factorial 5))
	_ (print "factorial(10) =" (factorial 10))
	
	# Fibonacci
	fibonacci (lambda n (match {n <= 1}
		true n
		(let
			p (fibonacci {n - 1})
			q (fibonacci {n - 2})
			{p + q}
		)
	))
	
	_ (print "\n=== Fibonacci Sequence ===")
	_ (print "fibonacci(0) =" (fibonacci 0))
	_ (print "fibonacci(1) =" (fibonacci 1))
	_ (print "fibonacci(2) =" (fibonacci 2))
	_ (print "fibonacci(3) =" (fibonacci 3))
	_ (print "fibonacci(4) =" (fibonacci 4))
	_ (print "fibonacci(5) =" (fibonacci 5))
	_ (print "fibonacci(10) =" (fibonacci 10))
	
	# Sum of digits
	sum_digits (lambda n (match {n < 10}
		true n
		(let
			last_digit {n % 10}
			remaining {n / 10}
			{last_digit + (sum_digits remaining)}
		)
	))
	
	_ (print "\n=== Sum of Digits ===")
	_ (print "sum_digits(123) =" (sum_digits 123))
	_ (print "sum_digits(456) =" (sum_digits 456))
	_ (print "sum_digits(999) =" (sum_digits 999))
	
	# Greatest Common Divisor
	gcd (lambda a b (match {b == 0}
		true a
		(gcd b {a % b})
	))
	
	_ (print "\n=== Greatest Common Divisor ===")
	_ (print "gcd(48, 18) =" (gcd 48 18))
	_ (print "gcd(100, 25) =" (gcd 100 25))
	_ (print "gcd(17, 13) =" (gcd 17 13))
	
	# Mutual recursion examples
	_ (print "\n=== Mutual Recursion ===")
	
	# Even/Odd
	even (lambda n (match {n <= 0}
		true true
		(odd {n - 1})
	))
	odd (lambda n (match {n <= 0}
		true false
		(even {n - 1})
	))
	
	_ (print "even(0) =" (even 0))
	_ (print "even(1) =" (even 1))
	_ (print "even(2) =" (even 2))
	_ (print "even(3) =" (even 3))
	_ (print "even(10) =" (even 10))
	_ (print "odd(0) =" (odd 0))
	_ (print "odd(1) =" (odd 1))
	_ (print "odd(2) =" (odd 2))
	_ (print "odd(3) =" (odd 3))
	_ (print "odd(10) =" (odd 10))
	
	# Ackermann function
	ackermann (lambda m n (match {m == 0}
		true {n + 1}
		(match {n == 0}
			true (ackermann {m - 1} 1)
			(ackermann {m - 1} (ackermann m {n - 1}))
		)
	))
	
	_ (print "\n=== Ackermann Function ===")
	_ (print "ackermann(0, 0) =" (ackermann 0 0))
	_ (print "ackermann(0, 1) =" (ackermann 0 1))
	_ (print "ackermann(1, 0) =" (ackermann 1 0))
	_ (print "ackermann(1, 1) =" (ackermann 1 1))
	_ (print "ackermann(2, 1) =" (ackermann 2 1))
	_ (print "ackermann(3, 1) =" (ackermann 3 1))
	
	# Tree recursion - counting paths
	count_paths (lambda m n (match {m == 1}
		true 1
		(match {n == 1}
			true 1
			{(count_paths {m - 1} n) + (count_paths m {n - 1})}
		)
	))
	
	_ (print "\n=== Count Paths (Tree Recursion) ===")
	_ (print "count_paths(1, 1) =" (count_paths 1 1))
	_ (print "count_paths(2, 2) =" (count_paths 2 2))
	_ (print "count_paths(3, 3) =" (count_paths 3 3))
	_ (print "count_paths(4, 4) =" (count_paths 4 4))
	
	# Recursive list operations
	_ (print "\n=== Recursive List Operations ===")
	
	# Length of list (recursive)
	length_rec (lambda l (match (len l)
		0 0
		{1 + (length_rec (rest l))}
	))
	
	_ (print "length_rec([1 2 3 4 5]) =" (length_rec [1 2 3 4 5]))
	_ (print "length_rec([]) =" (length_rec []))
	
	# Sum of list (recursive)
	sum_rec (lambda l (match (len l)
		0 0
		(let
			first (head l)
			rest (rest l)
			{first + (sum_rec rest)}
		)
	))
	
	_ (print "sum_rec([1 2 3 4 5]) =" (sum_rec [1 2 3 4 5]))
	_ (print "sum_rec([10 20 30]) =" (sum_rec [10 20 30]))
	
	# Reverse list (recursive)
	reverse_rec (lambda l (match (len l)
		0 []
		(let
			first (head l)
			rest (rest l)
			rest_reversed (reverse_rec rest)
			(list *rest_reversed first)
		)
	))
	
	_ (print "reverse_rec([1 2 3 4]) =" (reverse_rec [1 2 3 4]))
	_ (print "reverse_rec([\"a\" \"b\" \"c\"]) =" (reverse_rec ["a" "b" "c"]))
	
	# Flatten nested lists
	flatten (lambda l (match (len l)
		0 []
		(let
			first (head l)
			rest (rest l)
			rest_flattened (flatten rest)
			(match (type first)
				"list" (list *first *rest_flattened)
				(list first *rest_flattened)
			)
		)
	))
	
	_ (print "flatten([[1 2] [3 4] [5 6]]) =" (flatten [[1 2] [3 4] [5 6]]))
	_ (print "flatten([1 [2 3] 4 [5 [6 7]]]) =" (flatten [1 [2 3] 4 [5 [6 7]]]))
	
	# Binary search (recursive)
	binary_search (lambda l target (let
		search (lambda l target low high (match {low > high}
			true -1
			(let
				mid {(low + high) / 2}
				mid_val (get l mid)
				(match {mid_val == target}
					true mid
					(match {mid_val < target}
						true (search l target {mid + 1} high)
						(search l target low {mid - 1})
					)
				)
			)
		))
		(search l target 0 {len l - 1})
	))
	
	sorted_list [1 3 5 7 9 11 13 15 17 19]
	_ (print "binary_search([1 3 5 7 9 11 13 15 17 19], 7) =" (binary_search sorted_list 7))
	_ (print "binary_search([1 3 5 7 9 11 13 15 17 19], 10) =" (binary_search sorted_list 10))
	
	# Tower of Hanoi
	hanoi (lambda n from to aux (match {n == 1}
		true (print "Move disk from" from "to" to)
		(let
			_ (hanoi {n - 1} from aux to)
			_ (print "Move disk from" from "to" to)
			(hanoi {n - 1} aux to from)
		)
	))
	
	_ (print "\n=== Tower of Hanoi (3 disks) ===")
	_ (hanoi 3 "A" "C" "B")
	
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
