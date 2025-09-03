# Recursion Demo
# Demonstrates recursive functions and algorithms

(let
    # Factorial function
    _ (print "=== Factorial ===")
    factorial (lambda n (match {n <= 1}
        true 1
        {n * factorial {n - 1}}
    ))
    _ (print "factorial(0) =" (factorial 0))
    _ (print "factorial(1) =" (factorial 1))
    _ (print "factorial(5) =" (factorial 5))
    _ (print "factorial(6) =" (factorial 6))
    
    # Fibonacci sequence
    _ (print "\n=== Fibonacci ===")
    fib (lambda n (match {n <= 1}
        true n
        {fib {n - 1} + fib {n - 2}}
    ))
    _ (print "fib(0) =" (fib 0))
    _ (print "fib(1) =" (fib 1))
    _ (print "fib(2) =" (fib 2))
    _ (print "fib(5) =" (fib 5))
    _ (print "fib(8) =" (fib 8))
    
    # Sum of list elements
    _ (print "\n=== Sum of List ===")
    sum-list (lambda lst (match (len lst)
        0 0
        {*slice lst [0] + sum-list (slice lst (range 1 (len lst)))}
    ))
    numbers [1 2 3 4 5]
    _ (print "sum of" numbers "=" (sum-list numbers))
    
    # Product of list elements
    _ (print "\n=== Product of List ===")
    product-list (lambda lst (match (len lst)
        0 1
        {*slice lst [0] * product-list (slice lst (range 1 (len lst)))}
    ))
    _ (print "product of" numbers "=" (product-list numbers))
    
    # Count down
    _ (print "\n=== Count Down ===")
    countdown (lambda n (match {n <= 0}
        true []
        [n *countdown {n - 1}]
    ))
    _ (print "countdown(5):" (countdown 5))
    _ (print "countdown(3):" (countdown 3))
    
    # Count up
    _ (print "\n=== Count Up ===")
    countup (lambda n (match {n <= 0}
        true []
        [*countup {n - 1} n]
    ))
    _ (print "countup(5):" (countup 5))
    _ (print "countup(3):" (countup 3))
    
    # Greatest common divisor
    _ (print "\n=== Greatest Common Divisor ===")
    gcd (lambda a b (match b
        0 a
        (gcd b (mod a b))
    ))
    _ (print "gcd(48, 18) =" (gcd 48 18))
    _ (print "gcd(100, 25) =" (gcd 100 25))
    
    # Power function
    _ (print "\n=== Power Function ===")
    power (lambda base exp (match exp
        0 1
        {base * power base {exp - 1}}
    ))
    _ (print "2^3 =" (power 2 3))
    _ (print "3^4 =" (power 3 4))
    _ (print "5^2 =" (power 5 2))
    
    nil
)
