package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
)

// Advanced features and complex examples demonstration
func main() {
	program := `
(let
	_ (print "=== Advanced Data Structures ===")
	
	# Tree structure
	make_tree (lambda value left right (list "tree" value left right))
	leaf (lambda value (make_tree value nil nil))
	
	tree (make_tree 1 
		(make_tree 2 (leaf 4) (leaf 5))
		(make_tree 3 (leaf 6) (leaf 7))
	)
	
	_ (print "Tree structure:" tree)
	
	# Tree traversal
	preorder (lambda tree (match (type tree)
		"nil" []
		(let
			value (get tree 1)
			left (get tree 2)
			right (get tree 3)
			left_result (preorder left)
			right_result (preorder right)
			(list value *left_result *right_result)
		)
	))
	
	_ (print "Preorder traversal:" (preorder tree))
	
	# Graph representation (adjacency list)
	_ (print "\n=== Graph Operations ===")
	
	make_graph (lambda edges (let
		# Convert edge list to adjacency list
		build_adjacency (lambda edges (match (len edges)
			0 []
			(let
				edge (head edges)
				rest_edges (rest edges)
				from (get edge 0)
				to (get edge 1)
				rest_adj (build_adjacency rest_edges)
				# Add edge to adjacency list
				(match (len rest_adj)
					0 (list (list from (list to)))
					(let
						first_node (head rest_adj)
						rest_nodes (rest rest_adj)
						(match {from == (get first_node 0)}
							true (list (list from (list to *(get first_node 1))) *rest_nodes)
							(list first_node * (build_adjacency rest_edges))
						)
					)
				)
			)
		))
		(build_adjacency edges)
	))
	
	edges [[1 2] [1 3] [2 4] [3 4] [4 5]]
	graph (make_graph edges)
	_ (print "Graph edges:" edges)
	_ (print "Graph adjacency:" graph)
	
	# Object-oriented programming simulation
	_ (print "\n=== Object-Oriented Programming Simulation ===")
	
	# Create a "class" using closures
	make_point (lambda x y (let
		get_x (lambda () x)
		get_y (lambda () y)
		set_x (lambda new_x (make_point new_x y))
		set_y (lambda new_y (make_point x new_y))
		move (lambda dx dy (make_point {x + dx} {y + dy}))
		distance (lambda other (let
			other_x (other "get_x")
			other_y (other "get_y")
			dx {x - other_x}
			dy {y - other_y}
			{dx * dx + dy * dy}  # squared distance
		))
		# Method dispatcher
		(lambda method (match method
			"get_x" (get_x)
			"get_y" (get_y)
			"set_x" set_x
			"set_y" set_y
			"move" move
			"distance" distance
			"Error: unknown method"
		))
	))
	
	point1 (make_point 3 4)
	point2 (make_point 0 0)
	
	_ (print "Point1 x:" (point1 "get_x"))
	_ (print "Point1 y:" (point1 "get_y"))
	_ (print "Point2 x:" (point2 "get_x"))
	_ (print "Point2 y:" (point2 "get_y"))
	
	point1_moved (point1 "move" 1 1)
	_ (print "Point1 after move(1,1) x:" (point1_moved "get_x"))
	_ (print "Point1 after move(1,1) y:" (point1_moved "get_y"))
	
	# Functional programming patterns
	_ (print "\n=== Functional Programming Patterns ===")
	
	# Monad-like operations
	maybe_bind (lambda maybe f (match (type maybe)
		"nil" nil
		(f maybe)
	))
	
	maybe_map (lambda maybe f (match (type maybe)
		"nil" nil
		(f maybe)
	))
	
	safe_sqrt (lambda x (match {x >= 0}
		true {x * x}  # simplified sqrt
		nil
	))
	
	_ (print "maybe_map(4, safe_sqrt) =" (maybe_map 4 safe_sqrt))
	_ (print "maybe_map(-4, safe_sqrt) =" (maybe_map -4 safe_sqrt))
	
	# Currying and partial application
	_ (print "\n=== Currying and Partial Application ===")
	
	curry3 (lambda f (lambda a (lambda b (lambda c (f a b c)))))
	
	add3 (lambda a b c {a + b + c})
	curried_add3 (curry3 add3)
	
	_ (print "curried_add3(1)(2)(3) =" (((curried_add3 1) 2) 3))
	
	# Partial application
	partial3 (lambda f a (lambda b c (f a b c)))
	add_10 (partial3 add3 10)
	_ (print "add_10(5, 3) =" (add_10 5 3))
	
	# Memoization with closures
	_ (print "\n=== Memoization ===")
	
	memoize (lambda f (let
		cache []
		(lambda x (let
			# Simple linear search for existing result
			find_cached (lambda cache x (match (len cache)
				0 nil
				(let
					entry (head cache)
					key (get entry 0)
					value (get entry 1)
					rest (rest cache)
					(match {key == x}
						true value
						(find_cached rest x)
					)
				)
			))
			cached (find_cached cache x)
			(match cached
				nil (let
					result (f x)
					new_entry (list x result)
					cache (list *cache new_entry)
					result
				)
				cached
			)
		))
	))
	
	# Memoized fibonacci
	fib_memo (memoize (lambda n (match {n <= 1}
		true n
		{(fib_memo {n - 1}) + (fib_memo {n - 2})}
	)))
	
	_ (print "fib_memo(10) =" (fib_memo 10))
	_ (print "fib_memo(15) =" (fib_memo 15))
	_ (print "fib_memo(20) =" (fib_memo 20))
	
	# Lazy evaluation simulation
	_ (print "\n=== Lazy Evaluation Simulation ===")
	
	# Lazy list (infinite sequence)
	make_lazy_list (lambda start step (let
		current start
		(lambda () (let
			value current
			current {current + step}
			value
		))
	))
	
	natural_numbers (make_lazy_list 1 1)
	even_numbers (make_lazy_list 0 2)
	
	# Take first n elements from lazy list
	take (lambda lazy_list n (match {n <= 0}
		true []
		(let
			value (lazy_list)
			rest (take lazy_list {n - 1})
			(list value *rest)
		)
	))
	
	_ (print "First 10 natural numbers:" (take natural_numbers 10))
	_ (print "First 10 even numbers:" (take even_numbers 10))
	
	# Stream processing
	_ (print "\n=== Stream Processing ===")
	
	# Filter a lazy list
	filter_lazy (lambda lazy_list pred (lambda () (let
		value (lazy_list)
		(match (pred value)
			true value
			((filter_lazy lazy_list pred))
		)
	)))
	
	# Take while condition holds
	take_while (lambda lazy_list pred (let
		collect (lambda result (let
			value (lazy_list)
			(match (pred value)
				true (collect (list *result value))
				result
			)
		))
		(collect [])
	))
	
	# Numbers less than 20
	small_numbers (take_while natural_numbers (lambda x {x < 20}))
	_ (print "Numbers < 20:" small_numbers)
	
	# Event-driven programming simulation
	_ (print "\n=== Event-Driven Programming Simulation ===")
	
	# Simple event system
	make_event_system (lambda () (let
		handlers []
		add_handler (lambda event handler (let
			handlers (list *handlers (list event handler))
			handlers
		))
		emit (lambda event data (let
			# Find and call handlers for this event
			call_handlers (lambda handlers event data (match (len handlers)
				0 nil
				(let
					handler_entry (head handlers)
					rest_handlers (rest handlers)
					handler_event (get handler_entry 0)
					handler_func (get handler_entry 1)
					(match {handler_event == event}
						true (handler_func data)
						nil
					)
					(call_handlers rest_handlers event data)
				)
			))
			(call_handlers handlers event data)
		))
		# Return event system interface
		(lambda method (match method
			"add_handler" add_handler
			"emit" emit
			"Error: unknown method"
		))
	))
	
	event_system (make_event_system)
	
	# Add event handlers
	_ (event_system "add_handler" "user_login" (lambda data (print "User logged in:" data)))
	_ (event_system "add_handler" "user_logout" (lambda data (print "User logged out:" data)))
	
	# Emit events
	_ (event_system "emit" "user_login" "alice")
	_ (event_system "emit" "user_logout" "alice")
	
	# State machine
	_ (print "\n=== State Machine ===")
	
	make_state_machine (lambda initial_state transitions (let
		current_state initial_state
		transition (lambda event (let
			# Find transition for current state and event
			find_transition (lambda transitions state event (match (len transitions)
				0 nil
				(let
					transition (head transitions)
					rest_transitions (rest transitions)
					from_state (get transition 0)
					event_name (get transition 1)
					to_state (get transition 2)
					(match (list {from_state == state} {event_name == event})
						[true true] to_state
						(find_transition rest_transitions state event)
					)
				)
			))
			new_state (find_transition transitions current_state event)
			(match new_state
				nil (list "Error: no transition" current_state event)
				(let
					current_state new_state
					(list "OK" current_state)
				)
			)
		))
		get_state (lambda () current_state)
		(lambda method (match method
			"transition" transition
			"get_state" get_state
			"Error: unknown method"
		))
	))
	
	# Define state machine for a simple vending machine
	transitions [
		["idle" "coin" "ready"]
		["ready" "coin" "ready"]
		["ready" "select" "dispensing"]
		["dispensing" "complete" "idle"]
		["idle" "select" "error"]
		["error" "reset" "idle"]
	]
	
	vending_machine (make_state_machine "idle" transitions)
	
	_ (print "Initial state:" (vending_machine "get_state"))
	_ (print "Insert coin:" (vending_machine "transition" "coin"))
	_ (print "State after coin:" (vending_machine "get_state"))
	_ (print "Select item:" (vending_machine "transition" "select"))
	_ (print "State after select:" (vending_machine "get_state"))
	_ (print "Complete:" (vending_machine "transition" "complete"))
	_ (print "Final state:" (vending_machine "get_state"))
	
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
