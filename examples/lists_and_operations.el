(let
    # Lists and List Operations
    _ (print "=== Lists and Operations ===")

    # Basic list creation
    _ (print "Empty list:" [])
    _ (print "Simple list:" [1 2 3 4 5])
    _ (print "Mixed list:" ["hello" 42 true false])

    # List operations
    _ (let
        numbers [1 2 3 4 5]
        _ (print "Original list:" numbers)
        _ (print "Head (first element):" (head numbers))
        _ (print "Rest (all but first):" (rest numbers))
        _ (print "Last element:" (last numbers))
        _ (print "Init (all but last):" (init numbers))
        _ (print "Length:" (len numbers))
        nil
    )

    # List construction
    _ (print "Cons 0 to list:" (cons 0 [1 2 3 4 5]))
    _ (print "Append lists:" (append [1 2 3] [4 5 6]))

    # List processing
    _ (let
        numbers [1 2 3 4 5]
        _ (print "Map (double each):" (map numbers (lambda n {n x 2})))
        _ (print "Filter (even numbers):" (filter numbers (lambda n {n % 2 == 0})))
        _ (print "Sum:" (sum numbers))
        _ (print "Product:" (product numbers))
        _ (print "Maximum:" (max_list numbers))
        _ (print "Minimum:" (min_list numbers))
        nil
    )

    # Range generation
    _ (print "Range 1 to 10:" (range 1 11))
    _ (print "Range with step:" (range_step 0 20 2))

    # List utilities
    _ (let
        numbers [1 2 3 4 5]
        _ (print "Reverse:" (reverse numbers))
        _ (print "Take first 3:" (take 3 numbers))
        _ (print "Drop first 2:" (drop 2 numbers))
        nil
    )

    nil
)