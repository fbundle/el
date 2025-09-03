# Lists and List Operations
(print "=== Lists and Operations ===")

# Basic list creation
(print "Empty list:" [])
(print "Simple list:" [1 2 3 4 5])
(print "Mixed list:" ["hello" 42 true false])

# List operations
numbers [1 2 3 4 5]
(print "Original list:" numbers)
(print "Head (first element):" (head numbers))
(print "Rest (all but first):" (rest numbers))
(print "Last element:" (last numbers))
(print "Init (all but last):" (init numbers))
(print "Length:" (len numbers))

# List construction
(print "Cons 0 to list:" (cons 0 numbers))
(print "Append lists:" (append [1 2 3] [4 5 6]))

# List processing
(print "Map (double each):" (map numbers (lambda x {x * 2})))
(print "Filter (even numbers):" (filter numbers (lambda x {x % 2 == 0})))
(print "Sum:" (sum numbers))
(print "Product:" (product numbers))
(print "Maximum:" (max_list numbers))
(print "Minimum:" (min_list numbers))

# Range generation
(print "Range 1 to 10:" (range 1 11))
(print "Range with step:" (range_step 0 20 2))

# List utilities
(print "Reverse:" (reverse numbers))
(print "Take first 3:" (take 3 numbers))
(print "Drop first 2:" (drop 2 numbers))
