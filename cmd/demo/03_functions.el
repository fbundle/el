# Functions Demo
# Demonstrates function definitions, calls, and higher-order functions

(let
    # Simple function definitions
    _ (print "=== Simple Functions ===")
    square (lambda x {x * x})
    _ (print "square(5) =" (square 5))
    _ (print "square(7) =" (square 7))
    
    # Multi-parameter function
    add (lambda x y {x + y})
    multiply (lambda x y {x * y})
    _ (print "add(3, 4) =" (add 3 4))
    _ (print "multiply(3, 4) =" (multiply 3 4))
    
    # Arrow function syntax
    _ (print "\n=== Arrow Functions ===")
    subtract {x y => {x - y}}
    divide {x y => {x / y}}
    _ (print "subtract(10, 3) =" (subtract 10 3))
    _ (print "divide(20, 4) =" (divide 20 4))
    
    # Function composition
    _ (print "\n=== Function Composition ===")
    double (lambda x {x * 2})
    increment (lambda x {x + 1})
    _ (print "double(5) =" (double 5))
    _ (print "increment(5) =" (increment 5))
    
    # Higher-order functions
    _ (print "\n=== Higher-Order Functions ===")
    apply-twice (lambda f x (f (f x)))
    _ (print "apply-twice increment 5 =" (apply-twice increment 5))
    _ (print "apply-twice double 3 =" (apply-twice double 3))
    
    # Function that returns a function
    _ (print "\n=== Functions Returning Functions ===")
    make-adder (lambda n (lambda x {x + n}))
    add-5 (make-adder 5)
    add-10 (make-adder 10)
    _ (print "add-5(3) =" (add-5 3))
    _ (print "add-10(7) =" (add-10 7))
    
    # Complex function composition
    _ (print "\n=== Complex Function Composition ===")
    compose (lambda f g (lambda x (f (g x))))
    square-then-increment (compose increment square)
    _ (print "square-then-increment(3):" (square-then-increment 3))
    
    nil
)
