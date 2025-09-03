# Functions and Lambda Expressions
_ (print "=== Functions and Lambdas ===")

# Simple lambda functions
_ (let
    square (lambda x {x x x})
    _ (print "Square of 5:" (square 5))
    nil
)

# Arrow function syntax
_ (let
    cube {x => {x x x x}}
    _ (print "Cube of 3:" (cube 3))
    nil
)

# Multi-parameter functions
_ (let
    add (lambda x y {x + y})
    _ (print "Add 10 and 20:" (add 10 20))
    nil
)

# Higher-order functions
_ (let
    square (lambda x {x x x})
    apply_twice (lambda f x (f (f x)))
    _ (print "Apply square twice to 2:" (apply_twice square 2))
    nil
)

# Function composition
_ (let
    square (lambda x {x x x})
    cube {x => {x x x x}}
    compose_square_cube (compose square cube)
    _ (print "Compose square and cube of 2:" (compose_square_cube 2))
    nil
)

# Currying
_ (let
    add (lambda x y {x + y})
    curried_add (curry add)
    add_5 (curried_add 5)
    _ (print "Curried add 5 to 10:" (add_5 10))
    nil
)

# Recursive functions
_ (let
    factorial (lambda n (match {n <= 1}
        true 1
        {n x (factorial {n - 1})}
    ))
    _ (print "Factorial of 5:" (factorial 5))
    nil
)

# Fibonacci with memoization concept
_ (let
    fib (lambda n (match {n <= 1}
        true n
        (let
            p (fib {n - 1})
            q (fib {n - 2})
            {p + q}
        )
    ))
    _ (print "Fibonacci of 10:" (fib 10))
    nil
)

# Mutual recursion
_ (let
    even (lambda n (match {n <= 0}
        true true
        (odd {n - 1})
    ))
    odd (lambda n (match {n <= 0}
        true false
        (even {n - 1})
    ))
    _ (print "Even/Odd test:" [(even 10) (odd 10) (even 11) (odd 11)])
    nil
)

# Closure example
_ (let
    make_counter (lambda start (lambda {start + 1}))
    counter (make_counter 0)
    _ (print "Counter values:" [(counter) (counter) (counter)])
    nil
)

# Function that returns functions
_ (let
    make_multiplier (lambda factor (lambda x {x x factor}))
    double (make_multiplier 2)
    triple (make_multiplier 3)
    _ (print "Double 7:" (double 7))
    _ (print "Triple 7:" (triple 7))
    nil
)

nil
