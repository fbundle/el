# Basic Arithmetic Operations
_ (print "=== Basic Arithmetic ===")

# Simple arithmetic
_ (print "2 + 3 =" {2 + 3})
_ (print "10 - 4 =" {10 - 4})
_ (print "6 x 7 =" {6 x 7})
_ (print "15 / 3 =" {15 / 3})
_ (print "17 % 5 =" {17 % 5})

# Complex expressions
_ (print "Complex: (2 + 3) x 4 =" {2 + 3 x 4})
_ (print "Complex: 2 + 3 x 4 =" {2 + 3 x 4})

# Using variables
_ (let
    x 10
    y 20
    z {x + y}
    _ (print "x =" x)
    _ (print "y =" y)
    _ (print "x + y =" z)
    nil
)

nil
