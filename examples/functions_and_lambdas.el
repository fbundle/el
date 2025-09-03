# Functions and Lambda Expressions
(print "=== Functions and Lambdas ===")

# Simple lambda functions
square (lambda x {x * x})
(print "Square of 5:" (square 5))

# Arrow function syntax
cube {x => {x * x * x}}
(print "Cube of 3:" (cube 3))

# Multi-parameter functions
add (lambda x y {x + y})
(print "Add 10 and 20:" (add 10 20))

# Higher-order functions
apply_twice (lambda f x (f (f x)))
(print "Apply square twice to 2:" (apply_twice square 2))

# Function composition
compose_square_cube (compose square cube)
(print "Compose square and cube of 2:" (compose_square_cube 2))

# Currying
curried_add (curry add)
add_5 (curried_add 5)
(print "Curried add 5 to 10:" (add_5 10))

# Recursive functions
factorial (lambda n (match {n <= 1}
    true 1
    {n * (factorial {n - 1})}
))
(print "Factorial of 5:" (factorial 5))

# Fibonacci with memoization concept
fib (lambda n (match {n <= 1}
    true n
    (let
        p (fib {n - 1})
        q (fib {n - 2})
        {p + q}
    )
))
(print "Fibonacci of 10:" (fib 10))

# Mutual recursion
even (lambda n (match {n <= 0}
    true true
    (odd {n - 1})
))
odd (lambda n (match {n <= 0}
    true false
    (even {n - 1})
))
(print "Even/Odd test:" [(even 10) (odd 10) (even 11) (odd 11)])

# Closure example
make_counter (lambda start (lambda {start + 1}))
counter (make_counter 0)
(print "Counter values:" [(counter) (counter) (counter)])

# Function that returns functions
make_multiplier (lambda factor (lambda x {x * factor}))
double (make_multiplier 2)
triple (make_multiplier 3)
(print "Double 7:" (double 7))
(print "Triple 7:" (triple 7))
