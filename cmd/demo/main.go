package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/demo/main.go <demo-name>")
		fmt.Println("Available demos:")
		fmt.Println("  hello       - Hello World")
		fmt.Println("  arithmetic  - Basic arithmetic operations")
		fmt.Println("  functions   - Function definitions and calls")
		fmt.Println("  lists       - List operations")
		fmt.Println("  recursion   - Recursive functions")
		fmt.Println("  closures    - Closure examples")
		fmt.Println("  matching    - Pattern matching")
		fmt.Println("  types       - Type system examples")
		fmt.Println("  advanced    - Advanced features")
		fmt.Println("  all         - Run all demos")
		return
	}

	demoName := os.Args[1]

	switch demoName {
	case "hello":
		runDemo("Hello World", helloWorldDemo)
	case "arithmetic":
		runDemo("Arithmetic Operations", arithmeticDemo)
	case "functions":
		runDemo("Functions", functionsDemo)
	case "lists":
		runDemo("List Operations", listsDemo)
	case "recursion":
		runDemo("Recursion", recursionDemo)
	case "closures":
		runDemo("Closures", closuresDemo)
	case "matching":
		runDemo("Pattern Matching", matchingDemo)
	case "types":
		runDemo("Type System", typesDemo)
	case "advanced":
		runDemo("Advanced Features", advancedDemo)
	case "all":
		runAllDemos()
	default:
		fmt.Printf("Unknown demo: %s\n", demoName)
	}
}

func runDemo(title string, demoFunc func()) {
	fmt.Printf("\n=== %s ===\n", title)
	fmt.Println()
	demoFunc()
	fmt.Println()
}

func runAllDemos() {
	demos := []struct {
		title string
		fn    func()
	}{
		{"Hello World", helloWorldDemo},
		{"Arithmetic Operations", arithmeticDemo},
		{"Functions", functionsDemo},
		{"List Operations", listsDemo},
		{"Recursion", recursionDemo},
		{"Closures", closuresDemo},
		{"Pattern Matching", matchingDemo},
		{"Type System", typesDemo},
		{"Advanced Features", advancedDemo},
	}

	for _, demo := range demos {
		runDemo(demo.title, demo.fn)
	}
}

func executeProgram(program string) {
	tokens := parser.Tokenize(runtime_ext.WithTemplate(program))
	r, s := runtime_ext.NewBasicRuntime()

	var e ast.Expr
	var o runtime.Object
	var err error
	ctx := context.Background()

	for len(tokens) > 0 {
		e, tokens, err = parser.Parse(tokens)
		if err != nil {
			fmt.Printf("Parse error: %v\n", err)
			return
		}

		if err := r.Step(ctx, s, e).Unwrap(&o); err != nil {
			fmt.Printf("Runtime error: %v\n", err)
			return
		}

		if o != nil {
			fmt.Printf("Output: %v\n", o)
		}
	}
}

func helloWorldDemo() {
	program := `
(let
    _ (print "Hello, World!")
    _ (print "Welcome to El programming language!")
    nil
)`
	executeProgram(program)
}

func arithmeticDemo() {
	program := `
(let
    # Basic arithmetic
    _ (print "=== Basic Arithmetic ===")
    _ (print "1 + 2 =" {1 + 2})
    _ (print "10 - 3 =" {10 - 3})
    _ (print "4 * 5 =" {4 * 5})
    _ (print "20 / 4 =" {20 / 4})
    _ (print "17 % 5 =" (mod 17 5))
    
    # Complex expressions
    _ (print "\n=== Complex Expressions ===")
    _ (print "2 + 3 * 4 =" {2 + 3 * 4})
    _ (print "(2 + 3) * 4 =" {(2 + 3) * 4})
    _ (print "10 - 2 * 3 + 1 =" {10 - 2 * 3 + 1})
    
    # Comparisons
    _ (print "\n=== Comparisons ===")
    _ (print "5 == 5:" (eq 5 5))
    _ (print "5 != 3:" (ne 5 3))
    _ (print "3 < 5:" (lt 3 5))
    _ (print "5 <= 5:" (le 5 5))
    _ (print "7 > 3:" (gt 7 3))
    _ (print "4 >= 4:" (ge 4 4))
    
    nil
)`
	executeProgram(program)
}

func functionsDemo() {
	program := `
(let
    # Simple function
    _ (print "=== Simple Functions ===")
    square (lambda x {x * x})
    _ (print "square(5) =" (square 5))
    _ (print "square(7) =" (square 7))
    
    # Multi-parameter function
    add (lambda x y {x + y})
    _ (print "add(3, 4) =" (add 3 4))
    
    # Function composition
    _ (print "\n=== Function Composition ===")
    double (lambda x {x * 2})
    increment (lambda x {x + 1})
    _ (print "double(5) =" (double 5))
    _ (print "increment(5) =" (increment 5))
    
    # Arrow function syntax
    _ (print "\n=== Arrow Functions ===")
    multiply {x y => {x * y}}
    _ (print "multiply(3, 4) =" (multiply 3 4))
    
    # Higher-order function
    _ (print "\n=== Higher-Order Functions ===")
    apply-twice (lambda f x (f (f x)))
    _ (print "apply-twice increment 5 =" (apply-twice increment 5))
    _ (print "apply-twice double 3 =" (apply-twice double 3))
    
    nil
)`
	executeProgram(program)
}

func listsDemo() {
	program := `
(let
    # Basic list operations
    _ (print "=== Basic List Operations ===")
    numbers [1 2 3 4 5]
    _ (print "numbers:" numbers)
    _ (print "length:" (len numbers))
    
    # List creation
    _ (print "\n=== List Creation ===")
    mixed [1 "hello" true 42]
    _ (print "mixed list:" mixed)
    
    # Range function
    _ (print "\n=== Range Function ===")
    range1 (range 0 5)
    range2 (range 5 10)
    _ (print "range(0, 5):" range1)
    _ (print "range(5, 10):" range2)
    
    # List slicing
    _ (print "\n=== List Slicing ===")
    data [10 20 30 40 50 60 70 80 90]
    _ (print "original:" data)
    _ (print "first 3:" (slice data [0 1 2]))
    _ (print "last 3:" (slice data [6 7 8]))
    _ (print "every other:" (slice data [0 2 4 6 8]))
    
    # Nested lists
    _ (print "\n=== Nested Lists ===")
    matrix [[1 2 3] [4 5 6] [7 8 9]]
    _ (print "matrix:" matrix)
    _ (print "first row:" (slice matrix [0]))
    
    nil
)`
	executeProgram(program)
}

func recursionDemo() {
	program := `
(let
    # Factorial
    _ (print "=== Factorial ===")
    factorial (lambda n (match {n <= 1}
        true 1
        {n * factorial {n - 1}}
    ))
    _ (print "factorial(5) =" (factorial 5))
    _ (print "factorial(6) =" (factorial 6))
    
    # Fibonacci
    _ (print "\n=== Fibonacci ===")
    fib (lambda n (match {n <= 1}
        true n
        {fib {n - 1} + fib {n - 2}}
    ))
    _ (print "fib(0) =" (fib 0))
    _ (print "fib(1) =" (fib 1))
    _ (print "fib(5) =" (fib 5))
    _ (print "fib(8) =" (fib 8))
    
    # Sum of list
    _ (print "\n=== Sum of List ===")
    sum-list (lambda lst (match (len lst)
        0 0
        {*slice lst [0] + sum-list (slice lst (range 1 (len lst)))}
    ))
    numbers [1 2 3 4 5]
    _ (print "sum of" numbers "=" (sum-list numbers))
    
    # Count down
    _ (print "\n=== Count Down ===")
    countdown (lambda n (match {n <= 0}
        true []
        [n *countdown {n - 1}]
    ))
    _ (print "countdown(5):" (countdown 5))
    
    nil
)`
	executeProgram(program)
}

func closuresDemo() {
	program := `
(let
    # Simple closure
    _ (print "=== Simple Closure ===")
    make-counter (lambda start (lambda {start + 1}))
    counter1 (make-counter 10)
    counter2 (make-counter 100)
    _ (print "counter1:" (counter1))
    _ (print "counter1:" (counter1))
    _ (print "counter2:" (counter2))
    
    # Adder closure
    _ (print "\n=== Adder Closure ===")
    make-adder (lambda n (lambda x {x + n}))
    add-5 (make-adder 5)
    add-10 (make-adder 10)
    _ (print "add-5(3) =" (add-5 3))
    _ (print "add-5(7) =" (add-5 7))
    _ (print "add-10(3) =" (add-10 3))
    
    # Multiplier closure
    _ (print "\n=== Multiplier Closure ===")
    make-multiplier (lambda n (lambda x {x * n}))
    double (make-multiplier 2)
    triple (make-multiplier 3)
    _ (print "double(5) =" (double 5))
    _ (print "triple(4) =" (triple 4))
    
    # Accumulator closure
    _ (print "\n=== Accumulator Closure ===")
    make-accumulator (lambda initial (lambda x {initial + x}))
    acc (make-accumulator 0)
    _ (print "acc(5):" (acc 5))
    _ (print "acc(3):" (acc 3))
    _ (print "acc(10):" (acc 10))
    
    nil
)`
	executeProgram(program)
}

func matchingDemo() {
	program := `
(let
    # Simple pattern matching
    _ (print "=== Simple Pattern Matching ===")
    classify (lambda n (match n
        0 "zero"
        1 "one"
        2 "two"
        "many"
    ))
    _ (print "classify(0):" (classify 0))
    _ (print "classify(1):" (classify 1))
    _ (print "classify(2):" (classify 2))
    _ (print "classify(5):" (classify 5))
    
    # Boolean pattern matching
    _ (print "\n=== Boolean Pattern Matching ===")
    sign (lambda n (match {n < 0}
        true "negative"
        (match {n > 0}
            true "positive"
            "zero"
        )
    ))
    _ (print "sign(-5):" (sign -5))
    _ (print "sign(0):" (sign 0))
    _ (print "sign(5):" (sign 5))
    
    # Grade classification
    _ (print "\n=== Grade Classification ===")
    grade (lambda score (match {score >= 90}
        true "A"
        (match {score >= 80}
            true "B"
            (match {score >= 70}
                true "C"
                (match {score >= 60}
                    true "D"
                    "F"
                )
            )
        )
    ))
    _ (print "grade(95):" (grade 95))
    _ (print "grade(85):" (grade 85))
    _ (print "grade(75):" (grade 75))
    _ (print "grade(65):" (grade 65))
    _ (print "grade(55):" (grade 55))
    
    # List length pattern matching
    _ (print "\n=== List Length Pattern Matching ===")
    describe-list (lambda lst (match (len lst)
        0 "empty"
        1 "single"
        2 "pair"
        "many"
    ))
    _ (print "describe-list([]):" (describe-list []))
    _ (print "describe-list([1]):" (describe-list [1]))
    _ (print "describe-list([1 2]):" (describe-list [1 2]))
    _ (print "describe-list([1 2 3]):" (describe-list [1 2 3]))
    
    nil
)`
	executeProgram(program)
}

func typesDemo() {
	program := `
(let
    # Type introspection
    _ (print "=== Type Introspection ===")
    x 42
    y "hello"
    z [1 2 3]
    w true
    v nil
    
    _ (print "type of 42:" (type x))
    _ (print "type of 'hello':" (type y))
    _ (print "type of [1 2 3]:" (type z))
    _ (print "type of true:" (type w))
    _ (print "type of nil:" (type v))
    
    # Type of types
    _ (print "\n=== Type of Types ===")
    _ (print "type of (type 42):" (type (type x)))
    _ (print "type of (type (type 42)):" (type (type (type x))))
    
    # Function types
    _ (print "\n=== Function Types ===")
    add-func add
    lambda-func (lambda x {x + 1})
    _ (print "type of add function:" (type add-func))
    _ (print "type of lambda function:" (type lambda-func))
    
    # Type checking function
    _ (print "\n=== Type Checking Function ===")
    is-int (lambda x (eq (type x) "int"))
    is-string (lambda x (eq (type x) "string"))
    is-list (lambda x (eq (type x) "list"))
    
    _ (print "is-int(42):" (is-int 42))
    _ (print "is-int('hello'):" (is-int "hello"))
    _ (print "is-string('world'):" (is-string "world"))
    _ (print "is-list([1 2 3]):" (is-list [1 2 3]))
    
    nil
)`
	executeProgram(program)
}

func advancedDemo() {
	program := `
(let
    # Mutual recursion
    _ (print "=== Mutual Recursion ===")
    even (lambda n (match {n <= 0}
        true true
        (odd {n - 1})
    ))
    odd (lambda n (match {n <= 0}
        true false
        (even {n - 1})
    ))
    _ (print "even(4):" (even 4))
    _ (print "odd(4):" (odd 4))
    _ (print "even(5):" (even 5))
    _ (print "odd(5):" (odd 5))
    
    # Complex data structures
    _ (print "\n=== Complex Data Structures ===")
    person [name "Alice" age 30 city "NYC"]
    _ (print "person:" person)
    
    # Unwrap operator
    _ (print "\n=== Unwrap Operator ===")
    numbers [1 2 3]
    _ (print "numbers:" numbers)
    _ (print "sum with unwrap:" (add *numbers))
    _ (print "product with unwrap:" (mul *numbers))
    
    # Nested function calls
    _ (print "\n=== Nested Function Calls ===")
    compose (lambda f g (lambda x (f (g x))))
    square (lambda x {x * x})
    increment (lambda x {x + 1})
    square-then-increment (compose increment square)
    _ (print "square-then-increment(3):" (square-then-increment 3))
    
    # Complex arithmetic with precedence
    _ (print "\n=== Complex Arithmetic ===")
    _ (print "2^3 + 4*5 - 6/2 =" {2*2*2 + 4*5 - 6/2})
    _ (print "(2+3)*(4-1) =" {(2+3)*(4-1)})
    
    # String operations (basic)
    _ (print "\n=== String Operations ===")
    _ (print "Hello" "World")
    _ (print "Numbers:" 1 2 3)
    _ (print "Mixed:" "x =" 42 "y =" 3.14)
    
    nil
)`
	executeProgram(program)
}
