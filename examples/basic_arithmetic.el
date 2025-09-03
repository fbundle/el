# Basic Arithmetic Operations
(print "=== Basic Arithmetic ===")

# Simple arithmetic
(print "2 + 3 =" {2 + 3})
(print "10 - 4 =" {10 - 4})
(print "6 * 7 =" {6 * 7})
(print "15 / 3 =" {15 / 3})
(print "17 % 5 =" {17 % 5})

# Complex expressions
(print "Complex: (2 + 3) * 4 =" {(2 + 3) * 4})
(print "Complex: 2 + 3 * 4 =" {2 + 3 * 4})

# Using variables
(let
    x 10
    y 20
    z {x + y}
    (print "x =" x)
    (print "y =" y)
    (print "x + y =" z)
)
