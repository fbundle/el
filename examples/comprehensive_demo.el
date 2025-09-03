# Comprehensive Demo - Showcasing All EL Features
_ (print "=== EL Programming Language Comprehensive Demo ===")

# 1. Basic Operations
(print "\n1. Basic Operations:")
(print "Arithmetic: 2 + 3 * 4 =" {2 + 3 * 4})
(print "Comparison: 5 > 3 =" {5 > 3})
(print "Boolean: true and false =" (and true false))

# 2. Variables and Let Bindings
(print "\n2. Variables and Let Bindings:")
(let
    x 10
    y 20
    z {x + y}
    (print "x =" x ", y =" y ", z =" z)
)

# 3. Lists and List Operations
(print "\n3. Lists and Operations:")
numbers [1 2 3 4 5 6 7 8 9 10]
(print "Original:" numbers)
(print "Head:" (head numbers))
(print "Tail:" (rest numbers))
(print "Last:" (last numbers))
(print "Length:" (len numbers))
(print "Sum:" (sum numbers))
(print "Product:" (product numbers))
(print "Maximum:" (max_list numbers))
(print "Minimum:" (min_list numbers))

# 4. List Processing
(print "\n4. List Processing:")
(print "Squares:" (map numbers (lambda x {x * x})))
(print "Evens:" (filter numbers (lambda x {x % 2 == 0})))
(print "Odds:" (filter numbers (lambda x {x % 2 == 1})))
(print "Take 5:" (take 5 numbers))
(print "Drop 3:" (drop 3 numbers))
(print "Reverse:" (reverse numbers))

# 5. Functions and Lambdas
(print "\n5. Functions and Lambdas:")
square (lambda x {x * x})
cube {x => {x * x * x}}
(print "Square of 5:" (square 5))
(print "Cube of 3:" (cube 3))

# Higher-order functions
apply_twice (lambda f x (f (f x)))
(print "Apply square twice to 2:" (apply_twice square 2))

# Function composition
compose_square_cube (compose square cube)
(print "Compose square and cube of 2:" (compose_square_cube 2))

# 6. Recursive Functions
(print "\n6. Recursive Functions:")
factorial (lambda n (match {n <= 1}
    true 1
    {n * (factorial {n - 1})}
))
(print "Factorial of 5:" (factorial 5))

fibonacci (lambda n (match {n <= 1}
    true n
    (let
        p (fibonacci {n - 1})
        q (fibonacci {n - 2})
        {p + q}
    )
))
(print "Fibonacci of 10:" (fibonacci 10))

# 7. Pattern Matching
(print "\n7. Pattern Matching:")
grade (lambda score (match score
    100 "Perfect!"
    90 "Excellent"
    80 "Good"
    70 "Average"
    60 "Below Average"
    "Fail"
))
(print "Grade for 95:" (grade 95))
(print "Grade for 75:" (grade 75))

# 8. Type System
(print "\n8. Type System:")
(print "Type of 42:" (type 42))
(print "Type of 'hello':" (type "hello"))
(print "Type of [1 2 3]:" (type [1 2 3]))
(print "Type of true:" (type true))
(print "Type of lambda:" (type (lambda x x)))

# 9. Advanced List Operations
(print "\n9. Advanced List Operations:")
# Zip two lists
list1 [1 2 3 4]
list2 ["a" "b" "c" "d"]
zipped (zip list1 list2)
(print "Zip [1 2 3 4] and ['a' 'b' 'c' 'd']:" zipped)

# Range generation
(print "Range 1 to 10:" (range 1 11))
(print "Range with step 2:" (range_step 0 20 2))

# 10. Mathematical Functions
(print "\n10. Mathematical Functions:")
# Greatest Common Divisor
gcd (lambda a b (match {b == 0}
    true a
    (gcd b {a % b})
))
(print "GCD of 48 and 18:" (gcd 48 18))

# Prime checking (simplified)
is_prime (lambda n (match {n < 2}
    true false
    (match n
        2 true
        3 true
        5 true
        7 true
        11 true
        13 true
        17 true
        19 true
        23 true
        29 true
        false
    )
))
(print "Is 17 prime?" (is_prime 17))
(print "Is 15 prime?" (is_prime 15))

# 11. String Operations (simulated)
(print "\n11. String Operations:")
words ["hello" "world" "functional" "programming"]
(print "Words:" words)
(print "Word lengths:" (map words (lambda w (len w))))
(print "Long words:" (filter words (lambda w {len w > 5})))

# 12. Complex Data Structures
(print "\n12. Complex Data Structures:")
# Simple tree operations
tree_insert (lambda tree val (match (len tree)
    0 (list val [] [])
    (let
        root (get tree 0)
        left (get tree 1)
        right (get tree 2)
        (match {val < root}
            true (list root (tree_insert left val) right)
            (list root left (tree_insert right val))
        )
    )
))

tree_to_list (lambda tree (match (len tree)
    0 []
    (let
        root (get tree 0)
        left (get tree 1)
        right (get tree 2)
        left_list (tree_to_list left)
        right_list (tree_to_list right)
        (append (append left_list [root]) right_list)
    )
))

# Build and traverse tree
tree (tree_insert (tree_insert (tree_insert [] 5) 3) 7)
tree (tree_insert (tree_insert tree 1) 9)
(print "Tree as list (in-order):" (tree_to_list tree))

# 13. Performance and Optimization
(print "\n13. Performance and Optimization:")
# Tail-recursive countdown
countdown (lambda n (match {n <= 0}
    true 0
    (countdown {n - 1})
))
(print "Countdown from 1000:" (countdown 1000))

# 14. Error Handling and Edge Cases
(print "\n14. Error Handling:")
# Safe division
safe_div (lambda a b (match {b == 0}
    true "Division by zero"
    {a / b}
))
(print "10 / 2:" (safe_div 10 2))
(print "10 / 0:" (safe_div 10 0))

# Safe list access
safe_get (lambda lst i (match {i >= 0}
    true (match {i < (len lst)}
        true (get lst i)
        "Index out of bounds"
    )
    "Negative index"
))
(print "Safe get [1 2 3] at index 1:" (safe_get [1 2 3] 1))
(print "Safe get [1 2 3] at index 5:" (safe_get [1 2 3] 5))

# 15. Functional Programming Patterns
(print "\n15. Functional Programming Patterns:")
# Pipeline
pipeline (lambda lst (let
    step1 (map lst (lambda x {x + 1}))
    step2 (filter step1 (lambda x {x % 2 == 0}))
    step3 (map step2 (lambda x {x * 2}))
    step3
))
(print "Pipeline [1 2 3 4 5]:" (pipeline [1 2 3 4 5]))

# Currying
curried_add (curry add)
add_5 (curried_add 5)
(print "Curried add 5 to 10:" (add_5 10))

_ (print "\n=== Demo Complete ===")

nil
