(let
    # Functional Programming Techniques
    _ (print "=== Functional Programming ===")

    # Map, Filter, Reduce patterns
    _ (let
        numbers [1 2 3 4 5 6 7 8 9 10]
        _ (print "Original numbers:" numbers)
        nil
    )

    # Map transformations
    _ (let
        numbers [1 2 3 4 5 6 7 8 9 10]
        squares (map numbers (lambda x {x x x}))
        cubes (map numbers (lambda x {x x x x}))
        _ (print "Squares:" squares)
        _ (print "Cubes:" cubes)
        nil
    )

    # Filter operations
    _ (let
        numbers [1 2 3 4 5 6 7 8 9 10]
        evens (filter numbers (lambda x {x % 2 == 0}))
        odds (filter numbers (lambda x {x % 2 == 1}))
        primes (filter numbers (lambda x (match x
            1 false
            2 true
            3 true
            5 true
            7 true
            false
        )))
        _ (print "Even numbers:" evens)
        _ (print "Odd numbers:" odds)
        _ (print "Primes in range:" primes)
        nil
    )

    # Reduce operations
    _ (let
        numbers [1 2 3 4 5 6 7 8 9 10]
        sum_all (foldl numbers 0 add)
        product_all (foldl numbers 1 mul)
        _ (print "Sum of all numbers:" sum_all)
        _ (print "Product of all numbers:" product_all)
        nil
    )

    # String processing
    words ["hello" "world" "functional" "programming"]
    (print "Words:" words)

    # Map over strings (simulated)
    word_lengths (map words (lambda w (len w)))
    (print "Word lengths:" word_lengths)

    # Filter long words
    long_words (filter words (lambda w {len w > 5}))
    (print "Long words (>5 chars):" long_words)

    # Higher-order functions
    # Function that takes a predicate and returns a filter function
    make_filter (lambda pred (lambda lst (filter lst pred)))
    even_filter (make_filter (lambda x {x % 2 == 0}))
    (print "Even filter applied:" (even_filter numbers))

    # Function that takes a transformation and returns a map function
    make_mapper (lambda transform (lambda lst (map lst transform)))
    square_mapper (make_mapper (lambda x {x * x}))
    (print "Square mapper applied:" (square_mapper [1 2 3 4]))

    # Partial application
    multiply_by (lambda n (lambda x {x * n}))
    double (multiply_by 2)
    triple (multiply_by 3)
    (print "Double [1 2 3]:" (map [1 2 3] double))
    (print "Triple [1 2 3]:" (map [1 2 3] triple))

    # Function composition
    compose (lambda f g (lambda x (f (g x))))
    add_one (lambda x {x + 1})
    square (lambda x {x * x})
    add_one_then_square (compose square add_one)
    (print "Add one then square 3:" (add_one_then_square 3))

    # Pipeline operations
    pipeline (lambda lst (let
        step1 (map lst (lambda x {x + 1}))
        step2 (filter step1 (lambda x {x % 2 == 0}))
        step3 (map step2 (lambda x {x * 2}))
        step3
    ))
    (print "Pipeline [1 2 3 4 5]:" (pipeline [1 2 3 4 5]))

    # Memoization concept (simplified)
    # Note: This is a conceptual example - real memoization would require state
    fib_memo (lambda n (match {n <= 1}
        true n
        (let
            # In a real implementation, we'd check a memo table here
            p (fib_memo {n - 1})
            q (fib_memo {n - 2})
            {p + q}
        )
    ))

    (print "Fibonacci with memo concept (n=10):" (fib_memo 10))

    # Lazy evaluation simulation
    # Generate infinite sequence (limited by implementation)
    generate_nats (lambda start (cons start (generate_nats {start + 1})))
    first_10_nats (take 10 (generate_nats 0))
    (print "First 10 natural numbers:" first_10_nats)

    # Monadic operations (conceptually)
    # Maybe/Option type simulation
    safe_div (lambda a b (match {b == 0}
        true nil
        {a / b}
    ))

    safe_sqrt (lambda x (match {x < 0}
        true nil
        (sqrt_approx x)
    ))

    # Chain operations
    result1 (safe_div 10 2)
    result2 (match result1
        nil nil
        (safe_sqrt result1)
    )
    (print "Safe sqrt of (10/2):" result2)

    # List comprehensions (simulated)
    # Generate all pairs (x, y) where x + y = 10
    pairs_sum_10 (let
        xs (range 1 10)
        (foldl xs [] (lambda acc x (let
            y {10 - x}
            (cons (list x y) acc)
        )))
    )
    (print "Pairs that sum to 10:" pairs_sum_10)

    # Functional data structures
    # Tree operations (binary tree simulation)
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

    # Build a tree
    tree (tree_insert (tree_insert (tree_insert [] 5) 3) 7)
    _ (let
        tree (tree_insert (tree_insert tree 1) 9)
        _ (print "Tree as list (in-order):" (tree_to_list tree))
        nil
    )

    nil
)