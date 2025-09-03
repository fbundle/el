# Advanced Algorithms and Data Structures
_ (print "=== Advanced Algorithms ===")

# Sorting algorithms
# Bubble sort
bubble_sort (lambda lst (let
    n (len lst)
    (match {n <= 1}
        true lst
        (let
            # One pass of bubble sort
            sorted_once (bubble_pass lst 0)
            (bubble_sort sorted_once)
        )
    )
))

bubble_pass (lambda lst i (let
    n (len lst)
    (match {i >= {n - 1}}
        true lst
        (let
            current (get lst i)
            next (get lst {i + 1})
            (match {current > next}
                true (let
                    # Swap elements
                    new_lst (swap lst i {i + 1})
                    (bubble_pass new_lst {i + 1})
                )
                (bubble_pass lst {i + 1})
            )
        )
    )
))

swap (lambda lst i j (let
    n (len lst)
    (match {i >= n}
        true lst
        (match {j >= n}
            true lst
            (let
                # Create new list with swapped elements
                before_i (take i lst)
                after_i (drop {i + 1} lst)
                before_j (take j after_i)
                after_j (drop {j - i} after_i)
                elem_i (get lst i)
                elem_j (get lst j)
                (append (append (append before_i [elem_j]) before_j) (append [elem_i] after_j))
            )
        )
    )
))

# Test sorting
_ (let
    unsorted [5 2 8 1 9 3]
    _ (print "Original list:" unsorted)
    _ (print "Bubble sorted:" (bubble_sort unsorted))
    nil
)

# Binary search
binary_search (lambda lst target (let
    n (len lst)
    (match {n == 0}
        true -1
        (let
            mid {n / 2}
            mid_val (get lst mid)
            (match {mid_val == target}
                true mid
                (match {mid_val > target}
                    true (let
                        left_half (take mid lst)
                        result (binary_search left_half target)
                        (match {result == -1}
                            true -1
                            result
                        )
                    )
                    (let
                        right_half (drop {mid + 1} lst)
                        result (binary_search right_half target)
                        (match {result == -1}
                            true -1
                            {result + mid + 1}
                        )
                    )
                )
            )
        )
    )
))

_ (let
    sorted_list [1 3 5 7 9 11 13 15]
    _ (print "Searching in:" sorted_list)
    _ (print "Search for 7:" (binary_search sorted_list 7))
    _ (print "Search for 4:" (binary_search sorted_list 4))
    nil
)

# Prime number generation
is_prime (lambda n (match {n < 2}
    true false
    (let
        sqrt_n (sqrt_approx n)
        (check_divisors n 2 sqrt_n)
    )
))

sqrt_approx (lambda n (let
    guess {n / 2}
    (refine_sqrt n guess 0)
))

refine_sqrt (lambda n guess i (match {i >= 10}
    true guess
    (let
        new_guess {{guess + {n / guess}} / 2}
        (refine_sqrt n new_guess {i + 1})
    )
))

check_divisors (lambda n i limit (match {i > limit}
    true true
    (match {n % i == 0}
        true false
        (check_divisors n {i + 1} limit)
    )
))

# Generate primes up to n
_ (let
    primes_up_to (lambda n (let
        candidates (range 2 {n + 1})
        (filter candidates is_prime)
    ))
    _ (print "Primes up to 30:" (primes_up_to 30))
    nil
)

# Greatest Common Divisor
_ (let
    gcd (lambda a b (match {b == 0}
        true a
        (gcd b {a % b})
    ))
    _ (print "GCD of 48 and 18:" (gcd 48 18))
    _ (print "GCD of 17 and 13:" (gcd 17 13))
    nil
)

# Matrix operations (simple 2x2)
matrix_multiply (lambda m1 m2 (let
    a (get (get m1 0) 0)
    b (get (get m1 0) 1)
    c (get (get m1 1) 0)
    d (get (get m1 1) 1)
    e (get (get m2 0) 0)
    f (get (get m2 0) 1)
    g (get (get m2 1) 0)
    h (get (get m2 1) 1)
    (list
        (list {a * e + b * g} {a * f + b * h})
        (list {c * e + d * g} {c * f + d * h})
    )
))

_ (let
    matrix1 (list (list 1 2) (list 3 4))
    matrix2 (list (list 5 6) (list 7 8))
    _ (print "Matrix 1:" matrix1)
    _ (print "Matrix 2:" matrix2)
    _ (print "Matrix multiplication:" (matrix_multiply matrix1 matrix2))
    nil
)

nil
