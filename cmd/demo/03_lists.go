package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// List operations and manipulation demonstration
func main() {
	program := `
(let
	_ (print "=== List Creation ===")
	
	# Creating lists
	empty_list []
	numbers [1 2 3 4 5]
	strings ["hello" "world" "el"]
	mixed [1 "hello" 3.14 true]
	nested [[1 2] [3 4] [5 6]]
	
	_ (print "Empty list:" empty_list)
	_ (print "Numbers:" numbers)
	_ (print "Strings:" strings)
	_ (print "Mixed:" mixed)
	_ (print "Nested:" nested)
	
	# List length
	_ (print "\n=== List Length ===")
	_ (print "len([]) =" (len []))
	_ (print "len([1 2 3]) =" (len [1 2 3]))
	_ (print "len([\"a\" \"b\" \"c\" \"d\"]) =" (len ["a" "b" "c" "d"]))
	
	# List access
	_ (print "\n=== List Access ===")
	_ (print "head([1 2 3]) =" (head [1 2 3]))
	_ (print "rest([1 2 3]) =" (rest [1 2 3]))
	_ (print "get([10 20 30], 1) =" (get [10 20 30] 1))
	
	# List slicing
	_ (print "\n=== List Slicing ===")
	_ (print "slice([1 2 3 4 5], [0 2 4]) =" (slice [1 2 3 4 5] [0 2 4]))
	_ (print "slice([\"a\" \"b\" \"c\" \"d\"], [1 3]) =" (slice ["a" "b" "c" "d"] [1 3]))
	
	# Range generation
	_ (print "\n=== Range Generation ===")
	_ (print "range(0, 5) =" (range 0 5))
	_ (print "range(1, 6) =" (range 1 6))
	_ (print "range(10, 15) =" (range 10 15))
	
	# List concatenation (using unwrapping)
	_ (print "\n=== List Concatenation ===")
	list1 [1 2 3]
	list2 [4 5 6]
	_ (print "list1:" list1)
	_ (print "list2:" list2)
	_ (print "concat:" (list *list1 *list2))
	
	# List mapping
	_ (print "\n=== List Mapping ===")
	_ (print "map([1 2 3], square) =" (map [1 2 3] square))
	_ (print "map([1 2 3 4], (lambda x {x * 2})) =" (map [1 2 3 4] (lambda x {x * 2})))
	_ (print "map([\"a\" \"b\" \"c\"], (lambda s (list s s))) =" (map ["a" "b" "c"] (lambda s (list s s))))
	
	# List filtering (using map and conditional logic)
	_ (print "\n=== List Filtering ===")
	filter (lambda l pred (match (len l)
		0 []
		(let
			first (head l)
			rest (rest l)
			rest_filtered (filter rest pred)
			(match (pred first)
				true (list first *rest_filtered)
				rest_filtered
			)
		)
	))
	
	is_even (lambda x {x % 2 == 0})
	_ (print "filter([1 2 3 4 5 6], is_even) =" (filter [1 2 3 4 5 6] is_even))
	
	# List reduction
	_ (print "\n=== List Reduction ===")
	reduce (lambda l f init (match (len l)
		0 init
		(let
			first (head l)
			rest (rest l)
			new_init (f init first)
			(reduce rest f new_init)
		)
	))
	
	sum_list (lambda l (reduce l + 0))
	product_list (lambda l (reduce l * 1))
	_ (print "sum([1 2 3 4 5]) =" (sum_list [1 2 3 4 5]))
	_ (print "product([1 2 3 4]) =" (product_list [1 2 3 4]))
	
	# List searching
	_ (print "\n=== List Searching ===")
	contains (lambda l item (match (len l)
		0 false
		(let
			first (head l)
			rest (rest l)
			(match {first == item}
				true true
				(contains rest item)
			)
		)
	))
	
	_ (print "contains([1 2 3], 2) =" (contains [1 2 3] 2))
	_ (print "contains([1 2 3], 5) =" (contains [1 2 3] 5))
	
	# List reversal
	_ (print "\n=== List Reversal ===")
	reverse (lambda l (match (len l)
		0 []
		(let
			first (head l)
			rest (rest l)
			rest_reversed (reverse rest)
			(list *rest_reversed first)
		)
	))
	
	_ (print "reverse([1 2 3 4]) =" (reverse [1 2 3 4]))
	_ (print "reverse([\"a\" \"b\" \"c\"]) =" (reverse ["a" "b" "c"]))
	
	# List sorting (bubble sort)
	_ (print "\n=== List Sorting ===")
	insert (lambda item l (match (len l)
		0 (list item)
		(let
			first (head l)
			rest (rest l)
			(match {item <= first}
				true (list item *l)
				(list first * (insert item rest))
			)
		)
	))
	
	sort (lambda l (match (len l)
		0 []
		(let
			first (head l)
			rest (rest l)
			rest_sorted (sort rest)
			(insert first rest_sorted)
		)
	))
	
	_ (print "sort([3 1 4 1 5]) =" (sort [3 1 4 1 5]))
	_ (print "sort([10 5 8 2 9]) =" (sort [10 5 8 2 9]))
	
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
