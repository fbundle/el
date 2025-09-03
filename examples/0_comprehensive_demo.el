(let
    # Comprehensive Demo - Showcasing All EL Features
    _ (print "=== EL Programming Language Comprehensive Demo ===")

    # 1. Basic Operations
    _ (print "\n1. Basic Operations:")
    _ (print "Arithmetic: 2 + 3 * 4 =" {2 + 3 * 4})
    _ (print "Comparison: 5 > 3 =" {5 > 3})

    # 2. Variables and Let Bindings
    _ (print "\n2. Variables and Let Bindings:")
    _ (let
        x 10
        y 20
        z {x + y}
        (print "x =" x ", y =" y ", z =" z)
    )

    # 3. Lists and List Operations
    _ (print "\n3. Lists and Operations:")
    numbers [1 2 3 4 5 6 7 8 9 10]
    _ (print "Original:" numbers)
    _ (print "Head:" (head numbers))
    _ (print "Tail:" (rest numbers))
    _ (print "Last:" (last numbers))
    _ (print "Length:" (len numbers))
    _ (print "Sum:" (sum numbers))
    _ (print "Product:" (product numbers))
    _ (print "Maximum:" (max_list numbers))
    _ (print "Minimum:" (min_list numbers))

    # 4. List Processing
    _ (print "\n4. List Processing:")
    _ (print "Squares:" (map numbers (lambda x {x * x})))
    _ (print "Evens:" (filter numbers (lambda x {x % 2 == 0})))
    _ (print "Odds:" (filter numbers (lambda x {x % 2 == 1})))
    _ (print "Take 5:" (take 5 numbers))
    _ (print "Drop 3:" (drop 3 numbers))
    _ (print "Reverse:" (reverse numbers))

    # 5. Functions and Lambdas
    _ (print "\n5. Functions and Lambdas:")
    square (lambda x {x * x})
    cube {x => {x * x * x}}
    _ (print "Square of 5:" (square 5))
    _ (print "Cube of 3:" (cube 3))

    # Higher-order functions
    apply_twice (lambda f x (f (f x)))
    _ (print "Apply square twice to 2:" (apply_twice square 2))

    # Function composition
    compose {f g => {x => (f (g x))}}
    compose_square_cube (compose square cube)
    _ (print "Compose square and cube of 2:" (compose_square_cube 2))

    # 6. Recursive Functions
    _ (print "\n6. Recursive Functions:")
    factorial (lambda n (match {n <= 1}
        true 1
        {n * (factorial {n - 1})}
    ))
    _ (print "Factorial of 5:" (factorial 5))

    fibonacci (lambda n (match {n <= 1}
        true n
        (let
            p (fibonacci {n - 1})
            q (fibonacci {n - 2})
            {p + q}
        )
    ))
    _ (print "Fibonacci of 10:" (fibonacci 10))

    # 7. Pattern Matching
    _ (print "\n7. Pattern Matching:")
    grade (lambda score (match score
        100 "Perfect!"
        90 "Excellent"
        80 "Good"
        70 "Average"
        60 "Below Average"
        "Fail"
    ))
    _ (print "Grade for 95:" (grade 95))
    _ (print "Grade for 75:" (grade 75))

    # 8. Type System
    _ (print "\n8. Type System:")
    _ (print "Type of nil:" (type nil))
    _ (print "Type of 42:" (type 42))
    _ (print "Type of 'hello':" (type "hello"))
    _ (print "Type of [1 2 3]:" (type [1 2 3]))
    _ (print "Type of true:" (type true))
    _ (print "Type of lambda:" (type (lambda x x)))
    _ (print "Type of type:" (type (type nil)))
    _ (print "Type of type of type:" (type (type (type nil))))

    # 9. Advanced List Operations
    # _ (print "\n9. Advanced List Operations:")
    # Zip two lists
    # list1 [1 2 3 4]
    # list2 ["a" "b" "c" "d"]
    # zipped (zip list1 list2)
    # _ (print "Zip [1 2 3 4] and ['a' 'b' 'c' 'd']:" zipped)

    # Range generation
    _ (print "Range 1 to 10:" (range 1 11))
    _ (print "Range with step 2:" (range_step 0 20 2))

    # 10. Mathematical Functions
    _ (print "\n10. Mathematical Functions:")
    # Greatest Common Divisor
    gcd (lambda a b (match {b == 0}
        true a
        (gcd b {a % b})
    ))
    _ (print "GCD of 48 and 18:" (gcd 48 18))

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
    _ (print "Is 17 prime?" (is_prime 17))
    _ (print "Is 15 prime?" (is_prime 15))

    # 11. String Operations (simulated)
    # _ (print "\n11. String Operations:")
    # words ["hello" "world" "functional" "programming"]
    # _ (print "Words:" words)
    # _ (print "Word lengths:" (map words (lambda w (len w))))
    # _ (print "Long words:" (filter words (lambda w {len w > 5})))

    # 12. Complex Data Structures
    _ (print "\n12. Complex Data Structures:")
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
    _ (print "Tree as list (in-order):" (tree_to_list tree))

    # 13. Performance and Optimization
    _ (print "\n13. Performance and Optimization:")
    # Tail-recursive countdown
    countdown (lambda n (match {n <= 0}
        true 0
        (countdown {n - 1})
    ))
    _ (print "Countdown from 1000:" (countdown 1000))

    # 14. Error Handling and Edge Cases
    _ (print "\n14. Error Handling:")
    # Safe division
    safe_div (lambda a b (match {b == 0}
        true "Division by zero"
        {a / b}
    ))
    _ (print "10 / 2:" (safe_div 10 2))
    _ (print "10 / 0:" (safe_div 10 0))

    # Safe list access
    safe_get (lambda lst i (match {i >= 0}
        true (match {i < (len lst)}
            true (get lst i)
            "Index out of bounds"
        )
        "Negative index"
    ))
    _ (print "Safe get [1 2 3] at index 1:" (safe_get [1 2 3] 1))
    _ (print "Safe get [1 2 3] at index 5:" (safe_get [1 2 3] 5))

    # 15. Functional Programming Patterns
    _ (print "\n15. Functional Programming Patterns:")
    # Pipeline
    pipeline (lambda lst (let
        step1 (map lst (lambda x {x + 1}))
        step2 (filter step1 (lambda x {x % 2 == 0}))
        step3 (map step2 (lambda x {x * 2}))
        step3
    ))
    _ (print "Pipeline [1 2 3 4 5]:" (pipeline [1 2 3 4 5]))

    # Currying
    # curried_add (curry add)
    # add_5 (curried_add 5)
    # (print "Curried add 5 to 10:" (add_5 10))

    _ (print "\n=== Demo Complete ===")

    nil
)