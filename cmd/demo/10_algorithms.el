# Algorithms Demo
# Demonstrates common algorithms implemented in El

(let
    # Sorting algorithms
    _ (print "=== Sorting Algorithms ===")
    
    # Bubble sort
    bubble-sort (lambda lst (let
        swap (lambda lst i j (let
            temp (slice lst [i])
            new-lst (slice lst (range 0 i))
            new-lst [*new-lst *slice lst [j] *slice lst (range {i + 1} j) *temp *slice lst (range {j + 1} (len lst))]
            new-lst
        ))
        # Simplified bubble sort for demonstration
        _ (print "Bubble sort not fully implemented - showing concept")
        lst
    ))
    
    # Quick sort
    quicksort (lambda lst (match (len lst)
        0 []
        1 lst
        (let
            pivot (slice lst [0])
            rest (slice lst (range 1 (len lst)))
            smaller []
            larger []
            # Simplified - would need filter function
            _ (print "Quick sort concept - pivot:" pivot)
            [*pivot *rest]
        )
    ))
    
    # Search algorithms
    _ (print "\n=== Search Algorithms ===")
    
    # Linear search
    linear-search (lambda lst target (let
        search-helper (lambda lst target index (match (len lst)
            0 -1
            (match (eq *slice lst [0] target)
                true index
                (search-helper (slice lst (range 1 (len lst))) target {index + 1})
            )
        ))
        (search-helper lst target 0)
    ))
    
    numbers [1 3 5 7 9 11 13 15]
    _ (print "numbers:" numbers)
    _ (print "linear-search(numbers, 7):" (linear-search numbers 7))
    _ (print "linear-search(numbers, 4):" (linear-search numbers 4))
    
    # Binary search (simplified)
    binary-search (lambda lst target (let
        search-helper (lambda lst target left right (match {left > right}
            true -1
            (let
                mid {left + (right - left) / 2}
                mid-val (slice lst [mid])
                (match (eq *mid-val target)
                    true mid
                    (match (lt *mid-val target)
                        true (search-helper lst target {mid + 1} right)
                        (search-helper lst target left {mid - 1})
                    )
                )
            )
        ))
        (search-helper lst target 0 {len lst - 1})
    ))
    
    _ (print "binary-search(numbers, 7):" (binary-search numbers 7))
    _ (print "binary-search(numbers, 4):" (binary-search numbers 4))
    
    # Mathematical algorithms
    _ (print "\n=== Mathematical Algorithms ===")
    
    # Greatest Common Divisor (Euclidean algorithm)
    gcd (lambda a b (match b
        0 a
        (gcd b (mod a b))
    ))
    _ (print "gcd(48, 18):" (gcd 48 18))
    _ (print "gcd(100, 25):" (gcd 100 25))
    _ (print "gcd(17, 13):" (gcd 17 13))
    
    # Least Common Multiple
    lcm (lambda a b {a * b / gcd a b})
    _ (print "lcm(12, 18):" (lcm 12 18))
    _ (print "lcm(4, 6):" (lcm 4 6))
    
    # Prime checking
    is-prime (lambda n (let
        check-divisor (lambda n i (match {i * i > n}
            true true
            (match (eq (mod n i) 0)
                true false
                (check-divisor n {i + 1})
            )
        ))
        (match {n < 2}
            true false
            (check-divisor n 2)
        )
    ))
    _ (print "is-prime(17):" (is-prime 17))
    _ (print "is-prime(25):" (is-prime 25))
    _ (print "is-prime(29):" (is-prime 29))
    
    # Fibonacci with memoization concept
    _ (print "\n=== Fibonacci with Memoization Concept ===")
    fib-memo (lambda n (match n
        0 0
        1 1
        {fib-memo {n - 1} + fib-memo {n - 2}}
    ))
    _ (print "fib-memo(10):" (fib-memo 10))
    _ (print "fib-memo(15):" (fib-memo 15))
    
    # String algorithms
    _ (print "\n=== String Algorithms ===")
    
    # String length (simulation)
    string-length (lambda s (let
        _ (print "String length concept - would need string operations")
        0
    ))
    
    # Palindrome checking (concept)
    is-palindrome (lambda s (let
        _ (print "Palindrome checking concept")
        true
    ))
    
    # List algorithms
    _ (print "\n=== List Algorithms ===")
    
    # Reverse list
    reverse-list (lambda lst (match (len lst)
        0 []
        [*reverse-list (slice lst (range 1 (len lst))) *slice lst [0]]
    ))
    _ (print "reverse-list([1 2 3 4 5]):" (reverse-list [1 2 3 4 5]))
    
    # Concatenate lists
    concat-lists (lambda lst1 lst2 [*lst1 *lst2])
    _ (print "concat-lists([1 2 3], [4 5 6]):" (concat-lists [1 2 3] [4 5 6]))
    
    # Flatten nested lists (concept)
    flatten (lambda lst (match (len lst)
        0 []
        (let
            first (slice lst [0])
            rest (slice lst (range 1 (len lst)))
            (match (eq (type first) "list")
                true [*first *flatten rest]
                [first *flatten rest]
            )
        )
    ))
    _ (print "flatten([[1 2] [3 4] [5 6]]):" (flatten [[1 2] [3 4] [5 6]]))
    
    # Graph algorithms (concept)
    _ (print "\n=== Graph Algorithms (Concept) ===")
    _ (print "Graph algorithms would require more complex data structures")
    _ (print "Such as adjacency lists or matrices")
    
    nil
)