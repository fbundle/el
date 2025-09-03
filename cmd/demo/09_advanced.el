# Advanced Features Demo
# Demonstrates advanced language features and complex examples

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
    _ (print "even(0):" (even 0))
    _ (print "odd(0):" (odd 0))
    
    # Complex data structures
    _ (print "\n=== Complex Data Structures ===")
    person [name "Alice" age 30 city "NYC" salary 75000]
    _ (print "person:" person)
    
    # Unwrap operator
    _ (print "\n=== Unwrap Operator ===")
    numbers [1 2 3 4 5]
    _ (print "numbers:" numbers)
    _ (print "sum with unwrap:" (add *numbers))
    _ (print "product with unwrap:" (mul *numbers))
    
    # Nested function calls
    _ (print "\n=== Nested Function Calls ===")
    compose (lambda f g (lambda x (f (g x))))
    square (lambda x {x * x})
    increment (lambda x {x + 1})
    square-then-increment (compose increment square)
    increment-then-square (compose square increment)
    _ (print "square-then-increment(3):" (square-then-increment 3))
    _ (print "increment-then-square(3):" (increment-then-square 3))
    
    # Complex arithmetic with precedence
    _ (print "\n=== Complex Arithmetic ===")
    _ (print "2^3 + 4*5 - 6/2 =" {2*2*2 + 4*5 - 6/2})
    _ (print "(2+3)*(4-1) =" {(2+3)*(4-1)})
    _ (print "10 - 2*3 + 1 =" {10 - 2*3 + 1})
    _ (print "2*(3+4) - 5 =" {2*(3+4) - 5})
    
    # String operations (basic)
    _ (print "\n=== String Operations ===")
    _ (print "Hello" "World")
    _ (print "Numbers:" 1 2 3)
    _ (print "Mixed:" "x =" 42 "y =" 3.14)
    _ (print "Result:" "sum =" {1 + 2 + 3})
    
    # Function currying simulation
    _ (print "\n=== Function Currying ===")
    curry-add (lambda x (lambda y {x + y}))
    add-3 (curry-add 3)
    add-7 (curry-add 7)
    _ (print "add-3(5):" (add-3 5))
    _ (print "add-7(10):" (add-7 10))
    
    # Memoization simulation
    _ (print "\n=== Memoization Simulation ===")
    memo-fib (lambda n (match n
        0 0
        1 1
        {memo-fib {n - 1} + memo-fib {n - 2}}
    ))
    _ (print "memo-fib(8):" (memo-fib 8))
    _ (print "memo-fib(10):" (memo-fib 10))
    
    # List processing pipeline
    _ (print "\n=== List Processing Pipeline ===")
    numbers [1 2 3 4 5 6 7 8 9 10]
    _ (print "original:" numbers)
    
    # Filter even numbers (simulation)
    evens (slice numbers [1 3 5 7 9])  # indices of even numbers
    _ (print "evens:" evens)
    
    # Square each number
    squares (let
        square (lambda x {x * x})
        [*square *slice evens [0] *square *slice evens [1] *square *slice evens [2] *square *slice evens [3] *square *slice evens [4]]
    )
    _ (print "squares of evens:" squares)
    
    # Sum the squares
    sum-squares (add *squares)
    _ (print "sum of squares:" sum-squares)
    
    # Complex nested structures
    _ (print "\n=== Complex Nested Structures ===")
    data [[1 2 3] [4 5 6] [7 8 9]]
    _ (print "matrix:" data)
    
    # Access nested elements
    first-row (slice data [0])
    second-row (slice data [1])
    _ (print "first row:" first-row)
    _ (print "second row:" second-row)
    
    # Function that returns multiple values (as list)
    _ (print "\n=== Multiple Return Values ===")
    div-mod (lambda a b [(div a b) (mod a b)])
    result (div-mod 17 5)
    quotient (slice result [0])
    remainder (slice result [1])
    _ (print "17 / 5 =" quotient "remainder" remainder)
    
    # Higher-order function with multiple functions
    _ (print "\n=== Higher-Order Function with Multiple Functions ===")
    apply-all (lambda funcs x (let
        apply-one (lambda f (f x))
        [*apply-one *slice funcs [0] *apply-one *slice funcs [1] *apply-one *slice funcs [2]]
    ))
    funcs [square increment double]
    _ (print "apply-all to 3:" (apply-all funcs 3))
    
    # Complex conditional logic
    _ (print "\n=== Complex Conditional Logic ===")
    classify-number (lambda n (match {n < 0}
        true "negative"
        (match {n == 0}
            true "zero"
            (match {n < 10}
                true "single-digit"
                (match {n < 100}
                    true "double-digit"
                    "large"
                )
            )
        )
    ))
    _ (print "classify-number(-5):" (classify-number -5))
    _ (print "classify-number(0):" (classify-number 0))
    _ (print "classify-number(7):" (classify-number 7))
    _ (print "classify-number(42):" (classify-number 42))
    _ (print "classify-number(123):" (classify-number 123))
    
    nil
)
