# Pattern Matching and Conditionals
(print "=== Pattern Matching ===")

# Basic pattern matching
(print "Match 1 == 1:" (match 1 1 "yes" "no"))
(print "Match 1 == 2:" (match 1 2 "yes" "no"))

# Multiple conditions
grade_check (lambda score (match score
    100 "Perfect!"
    90 "Excellent"
    80 "Good"
    70 "Average"
    60 "Below Average"
    "Fail"
))
(print "Grade for 95:" (grade_check 95))
(print "Grade for 75:" (grade_check 75))
(print "Grade for 45:" (grade_check 45))

# Boolean operations
(print "True and True:" (and true true))
(print "True and False:" (and true false))
(print "False or True:" (or false true))
(print "Not True:" (not true))

# Conditional expressions
(print "If 5 > 3:" (if {5 > 3} "greater" "not greater"))
(print "If 2 > 5:" (if {2 > 5} "greater" "not greater"))

# Complex pattern matching with calculations
abs (lambda x (match {x < 0}
    true {-x}
    x
))
(print "Absolute value of -5:" (abs -5))
(print "Absolute value of 5:" (abs 5))

# String pattern matching
greet (lambda name (match name
    "Alice" "Hello Alice!"
    "Bob" "Hi Bob!"
    "Hello stranger!"
))
(print "Greet Alice:" (greet "Alice"))
(print "Greet Bob:" (greet "Bob"))
(print "Greet Charlie:" (greet "Charlie"))

# List pattern matching
list_length (lambda lst (match (len lst)
    0 0
    (let
        rest_lst (rest lst)
        {1 + (list_length rest_lst)}
    )
))
(print "Length of [1 2 3 4]:" (list_length [1 2 3 4]))
(print "Length of []:" (list_length []))

# Type checking with pattern matching
type_check (lambda x (match (type x)
    "number" "It's a number"
    "string" "It's a string"
    "list" "It's a list"
    "function" "It's a function"
    "unknown type"
))
(print "Type of 42:" (type_check 42))
(print "Type of 'hello':" (type_check "hello"))
(print "Type of [1 2 3]:" (type_check [1 2 3]))
