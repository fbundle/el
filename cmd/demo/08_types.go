package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Type system and introspection demonstration
func main() {
	program := `
(let
	_ (print "=== Basic Type Introspection ===")
	
	# Basic types
	_ (print "type(42) =" (type 42))
	_ (print "type(\"hello\") =" (type "hello"))
	_ (print "type([1 2 3]) =" (type [1 2 3]))
	_ (print "type(nil) =" (type nil))
	_ (print "type((lambda x x)) =" (type (lambda x x)))
	
	# Type of type
	_ (print "\n=== Type of Type ===")
	_ (print "type(type) =" (type type))
	_ (print "type(type(42)) =" (type (type 42)))
	_ (print "type(type(type(42))) =" (type (type (type 42))))
	
	# Type checking functions
	_ (print "\n=== Type Checking Functions ===")
	
	is_int (lambda x {type x == "int"})
	is_string (lambda x {type x == "string"})
	is_list (lambda x {type x == "list"})
	is_function (lambda x {type x == "function"})
	is_nil (lambda x {type x == "nil"})
	
	_ (print "is_int(42) =" (is_int 42))
	_ (print "is_int(\"hello\") =" (is_int "hello"))
	_ (print "is_string(\"world\") =" (is_string "world"))
	_ (print "is_string(123) =" (is_string 123))
	_ (print "is_list([1 2 3]) =" (is_list [1 2 3]))
	_ (print "is_list(42) =" (is_list 42))
	_ (print "is_function((lambda x x)) =" (is_function (lambda x x)))
	_ (print "is_nil(nil) =" (is_nil nil))
	_ (print "is_nil(0) =" (is_nil 0))
	
	# Type-safe operations
	_ (print "\n=== Type-Safe Operations ===")
	
	safe_add (lambda a b (match (list (type a) (type b))
		["int" "int"] {a + b}
		"Error: both arguments must be integers"
	))
	
	_ (print "safe_add(5, 3) =" (safe_add 5 3))
	_ (print "safe_add(5, \"hello\") =" (safe_add 5 "hello"))
	_ (print "safe_add(\"a\", \"b\") =" (safe_add "a" "b"))
	
	# Type conversion (simulated)
	_ (print "\n=== Type Conversion ===")
	
	# Convert to string representation
	to_string (lambda x (match (type x)
		"int" (match x
			0 "zero"
			1 "one"
			2 "two"
			3 "three"
			4 "four"
			5 "five"
			"many"
		)
		"string" x
		"list" "a list"
		"function" "a function"
		"nil" "nothing"
		"unknown"
	))
	
	_ (print "to_string(0) =" (to_string 0))
	_ (print "to_string(3) =" (to_string 3))
	_ (print "to_string(10) =" (to_string 10))
	_ (print "to_string(\"hello\") =" (to_string "hello"))
	_ (print "to_string([1 2 3]) =" (to_string [1 2 3]))
	_ (print "to_string(nil) =" (to_string nil))
	
	# Type-based dispatch
	_ (print "\n=== Type-Based Dispatch ===")
	
	process (lambda x (match (type x)
		"int" (match {x < 0}
			true "negative integer"
			"positive integer"
		)
		"string" (match (len x)
			0 "empty string"
			"non-empty string"
		)
		"list" (match (len x)
			0 "empty list"
			"non-empty list"
		)
		"function" "a function"
		"nil" "nothing"
		"unknown type"
	))
	
	_ (print "process(-5) =" (process -5))
	_ (print "process(10) =" (process 10))
	_ (print "process(\"\") =" (process ""))
	_ (print "process(\"hello\") =" (process "hello"))
	_ (print "process([]) =" (process []))
	_ (print "process([1 2]) =" (process [1 2]))
	_ (print "process((lambda x x)) =" (process (lambda x x)))
	_ (print "process(nil) =" (process nil))
	
	# Type hierarchy simulation
	_ (print "\n=== Type Hierarchy Simulation ===")
	
	# Simulate type hierarchy with nested types
	get_type_level (lambda x (let
		level_0 (type x)
		level_1 (type level_0)
		level_2 (type level_1)
		(list level_0 level_1 level_2)
	))
	
	_ (print "get_type_level(42) =" (get_type_level 42))
	_ (print "get_type_level(\"hello\") =" (get_type_level "hello"))
	_ (print "get_type_level([1 2 3]) =" (get_type_level [1 2 3]))
	
	# Type validation
	_ (print "\n=== Type Validation ===")
	
	validate_int (lambda x (match (type x)
		"int" true
		false
	))
	
	validate_list (lambda x (match (type x)
		"list" true
		false
	))
	
	validate_function (lambda x (match (type x)
		"function" true
		false
	))
	
	_ (print "validate_int(42) =" (validate_int 42))
	_ (print "validate_int(\"hello\") =" (validate_int "hello"))
	_ (print "validate_list([1 2 3]) =" (validate_list [1 2 3]))
	_ (print "validate_list(42) =" (validate_list 42))
	_ (print "validate_function((lambda x x)) =" (validate_function (lambda x x)))
	_ (print "validate_function(42) =" (validate_function 42))
	
	# Type-safe list operations
	_ (print "\n=== Type-Safe List Operations ===")
	
	safe_head (lambda l (match (type l)
		"list" (match (len l)
			0 "Error: empty list"
			(head l)
		)
		"Error: not a list"
	))
	
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
	
	_ (print "safe_head([1 2 3]) =" (safe_head [1 2 3]))
	_ (print "safe_head([]) =" (safe_head []))
	_ (print "safe_head(42) =" (safe_head 42))
	_ (print "safe_get([10 20 30], 1) =" (safe_get [10 20 30] 1))
	_ (print "safe_get([10 20 30], 5) =" (safe_get [10 20 30] 5))
	_ (print "safe_get([10 20 30], -1) =" (safe_get [10 20 30] -1))
	_ (print "safe_get(42, 1) =" (safe_get 42 1))
	
	# Type introspection for debugging
	_ (print "\n=== Type Introspection for Debugging ===")
	
	debug_info (lambda x (let
		x_type (type x)
		x_type_type (type x_type)
		info (match x_type
			"int" (list "integer" x)
			"string" (list "string" (len x))
			"list" (list "list" (len x))
			"function" (list "function" "closure")
			"nil" (list "nil" "nothing")
			(list "unknown" x_type)
		)
		(list x info)
	))
	
	_ (print "debug_info(42) =" (debug_info 42))
	_ (print "debug_info(\"hello\") =" (debug_info "hello"))
	_ (print "debug_info([1 2 3]) =" (debug_info [1 2 3]))
	_ (print "debug_info((lambda x x)) =" (debug_info (lambda x x)))
	_ (print "debug_info(nil) =" (debug_info nil))
	
	# Type-based function selection
	_ (print "\n=== Type-Based Function Selection ===")
	
	# Different implementations based on type
	generic_add (lambda a b (match (list (type a) (type b))
		["int" "int"] {a + b}
		["string" "string"] (list a b)  # Concatenate strings as list
		["list" "list"] (list *a *b)    # Concatenate lists
		"Error: incompatible types"
	))
	
	_ (print "generic_add(5, 3) =" (generic_add 5 3))
	_ (print "generic_add(\"hello\", \"world\") =" (generic_add "hello" "world"))
	_ (print "generic_add([1 2], [3 4]) =" (generic_add [1 2] [3 4]))
	_ (print "generic_add(5, \"hello\") =" (generic_add 5 "hello"))
	
	# Type checking in recursive functions
	_ (print "\n=== Type Checking in Recursive Functions ===")
	
	# Count elements of specific type in a list
	count_type (lambda l target_type (match (type l)
		"list" (match (len l)
			0 0
			(let
				first (head l)
				rest (rest l)
				first_count (match {type first == target_type}
					true 1
					0
				)
				rest_count (count_type rest target_type)
				{first_count + rest_count}
			)
		)
		0
	))
	
	mixed_list [1 "hello" 2 "world" 3 "test" 4]
	_ (print "mixed_list:" mixed_list)
	_ (print "count_type(mixed_list, \"int\") =" (count_type mixed_list "int"))
	_ (print "count_type(mixed_list, \"string\") =" (count_type mixed_list "string"))
	_ (print "count_type(mixed_list, \"list\") =" (count_type mixed_list "list"))
	
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
