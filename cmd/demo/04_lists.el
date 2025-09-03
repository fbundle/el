# Lists Demo
# Demonstrates list operations, creation, and manipulation

(let
    # Basic list operations
    _ (print "=== Basic List Operations ===")
    numbers [1 2 3 4 5]
    _ (print "numbers:" numbers)
    _ (print "length:" (len numbers))
    
    # List creation methods
    _ (print "\n=== List Creation ===")
    list1 [1 2 3 4 5]
    list2 (list 6 7 8 9 10)
    mixed [1 "hello" true 42]
    _ (print "list1:" list1)
    _ (print "list2:" list2)
    _ (print "mixed list:" mixed)
    
    # Range function
    _ (print "\n=== Range Function ===")
    range1 (range 0 5)
    range2 (range 5 10)
    range3 (range 1 11)
    _ (print "range(0, 5):" range1)
    _ (print "range(5, 10):" range2)
    _ (print "range(1, 11):" range3)
    
    # List slicing
    _ (print "\n=== List Slicing ===")
    data [10 20 30 40 50 60 70 80 90]
    _ (print "original:" data)
    _ (print "first 3:" (slice data [0 1 2]))
    _ (print "last 3:" (slice data [6 7 8]))
    _ (print "every other:" (slice data [0 2 4 6 8]))
    _ (print "middle 3:" (slice data [3 4 5]))
    
    # Nested lists
    _ (print "\n=== Nested Lists ===")
    matrix [[1 2 3] [4 5 6] [7 8 9]]
    _ (print "matrix:" matrix)
    _ (print "first row:" (slice matrix [0]))
    _ (print "second row:" (slice matrix [1]))
    
    # List of functions
    _ (print "\n=== List of Functions ===")
    operations [add sub mul div]
    _ (print "operations:" operations)
    
    # Empty list
    _ (print "\n=== Empty List ===")
    empty []
    _ (print "empty list:" empty)
    _ (print "empty list length:" (len empty))
    
    # List with nil
    _ (print "\n=== List with Nil ===")
    with-nil [1 nil 3 nil 5]
    _ (print "list with nil:" with-nil)
    
    nil
)
