# Math Problems Demo
# Demonstrates various mathematical problems and solutions in El

(let
    # Number theory problems
    _ (print "=== Number Theory Problems ===")
    
    # Perfect numbers
    is-perfect (lambda n (let
        sum-divisors (lambda n (let
            sum-helper (lambda n i sum (match {i >= n}
                true sum
                (match (eq (mod n i) 0)
                    true (sum-helper n {i + 1} {sum + i})
                    (sum-helper n {i + 1} sum)
                )
            ))
            (sum-helper n 1 0)
        ))
        (eq (sum-divisors n) n)
    ))
    _ (print "is-perfect(6):" (is-perfect 6))
    _ (print "is-perfect(28):" (is-perfect 28))
    _ (print "is-perfect(12):" (is-perfect 12))
    
    # Amicable numbers
    sum-proper-divisors (lambda n (let
        sum-helper (lambda n i sum (match {i >= n}
            true sum
            (match (eq (mod n i) 0)
                true (sum-helper n {i + 1} {sum + i})
                (sum-helper n {i + 1} sum)
            )
        ))
        (sum-helper n 1 0)
    ))
    
    are-amicable (lambda a b (and
        (eq (sum-proper-divisors a) b)
        (eq (sum-proper-divisors b) a)
        (ne a b)
    ))
    _ (print "are-amicable(220, 284):" (are-amicable 220 284))
    _ (print "are-amicable(1184, 1210):" (are-amicable 1184 1210))
    
    # Combinatorics
    _ (print "\n=== Combinatorics ===")
    
    # Factorial
    factorial (lambda n (match {n <= 1}
        true 1
        {n * factorial {n - 1}}
    ))
    _ (print "factorial(5):" (factorial 5))
    _ (print "factorial(10):" (factorial 10))
    
    # Permutations (concept)
    permutations (lambda n r (div (factorial n) (factorial {n - r})))
    _ (print "permutations(5, 3):" (permutations 5 3))
    _ (print "permutations(10, 2):" (permutations 10 2))
    
    # Combinations (concept)
    combinations (lambda n r (div (permutations n r) (factorial r)))
    _ (print "combinations(5, 3):" (combinations 5 3))
    _ (print "combinations(10, 2):" (combinations 10 2))
    
    # Fibonacci variations
    _ (print "\n=== Fibonacci Variations ===")
    
    # Tribonacci
    tribonacci (lambda n (match n
        0 0
        1 1
        2 1
        {tribonacci {n - 1} + tribonacci {n - 2} + tribonacci {n - 3}}
    ))
    _ (print "tribonacci(10):" (tribonacci 10))
    _ (print "tribonacci(15):" (tribonacci 15))
    
    # Lucas numbers
    lucas (lambda n (match n
        0 2
        1 1
        {lucas {n - 1} + lucas {n - 2}}
    ))
    _ (print "lucas(10):" (lucas 10))
    _ (print "lucas(15):" (lucas 15))
    
    # Pell numbers
    pell (lambda n (match n
        0 0
        1 1
        {2 * pell {n - 1} + pell {n - 2}}
    ))
    _ (print "pell(10):" (pell 10))
    _ (print "pell(15):" (pell 15))
    
    # Geometric problems
    _ (print "\n=== Geometric Problems ===")
    
    # Pythagorean triples
    is-pythagorean (lambda a b c (eq {a*a + b*b} {c*c}))
    _ (print "is-pythagorean(3, 4, 5):" (is-pythagorean 3 4 5))
    _ (print "is-pythagorean(5, 12, 13):" (is-pythagorean 5 12 13))
    _ (print "is-pythagorean(1, 2, 3):" (is-pythagorean 1 2 3))
    
    # Triangle area (Heron's formula)
    triangle-area (lambda a b c (let
        s {a + b + c / 2}
        {s * {s - a} * {s - b} * {s - c}}
    ))
    _ (print "triangle-area(3, 4, 5):" (triangle-area 3 4 5))
    
    # Number sequences
    _ (print "\n=== Number Sequences ===")
    
    # Triangular numbers
    triangular (lambda n {n * {n + 1} / 2})
    _ (print "triangular(10):" (triangular 10))
    _ (print "triangular(20):" (triangular 20))
    
    # Square numbers
    square (lambda n {n * n})
    _ (print "square(5):" (square 5))
    _ (print "square(10):" (square 10))
    
    # Cubic numbers
    cube (lambda n {n * n * n})
    _ (print "cube(3):" (cube 3))
    _ (print "cube(5):" (cube 5))
    
    # Pentagonal numbers
    pentagonal (lambda n {n * {3*n - 1} / 2})
    _ (print "pentagonal(5):" (pentagonal 5))
    _ (print "pentagonal(10):" (pentagonal 10))
    
    # Mathematical constants and approximations
    _ (print "\n=== Mathematical Constants ===")
    
    # Pi approximation (Leibniz formula)
    pi-approx (lambda terms (let
        pi-helper (lambda n sum (match {n >= terms}
            true {sum * 4}
            (let
                term {1 / {2*n + 1}}
                sign (match {n % 2}
                    0 1
                    -1
                )
                (pi-helper {n + 1} {sum + sign * term})
            )
        ))
        (pi-helper 0 0)
    ))
    _ (print "pi approximation (10 terms):" (pi-approx 10))
    _ (print "pi approximation (100 terms):" (pi-approx 100))
    
    # E approximation
    e-approx (lambda terms (let
        e-helper (lambda n sum (match {n >= terms}
            true sum
            (e-helper {n + 1} {sum + 1 / factorial n})
        ))
        (e-helper 0 0)
    ))
    _ (print "e approximation (10 terms):" (e-approx 10))
    
    # Golden ratio
    golden-ratio (lambda n (let
        phi-helper (lambda n (match {n <= 0}
            true 1
            {1 + 1 / phi-helper {n - 1}}
        ))
        (phi-helper n)
    ))
    _ (print "golden ratio approximation (10 iterations):" (golden-ratio 10))
    
    # Number base conversions
    _ (print "\n=== Number Base Conversions ===")
    
    # Decimal to binary (concept)
    decimal-to-binary (lambda n (match n
        0 "0"
        1 "1"
        (let
            _ (print "Binary conversion concept - would need string operations")
            "concept"
        )
    ))
    _ (print "decimal-to-binary(10):" (decimal-to-binary 10))
    
    # Digit sum
    digit-sum (lambda n (match {n < 10}
        true n
        {digit-sum {n / 10} + digit-sum {n % 10}}
    ))
    _ (print "digit-sum(12345):" (digit-sum 12345))
    _ (print "digit-sum(98765):" (digit-sum 98765))
    
    # Digital root
    digital-root (lambda n (match {n < 10}
        true n
        (digital-root (digit-sum n))
    ))
    _ (print "digital-root(12345):" (digital-root 12345))
    _ (print "digital-root(98765):" (digital-root 98765))
    
    # Prime factorization
    _ (print "\n=== Prime Factorization ===")
    prime-factors (lambda n (let
        factor-helper (lambda n i factors (match {i * i > n}
            true (match {n > 1}
                true [*factors n]
                factors
            )
            (match (eq (mod n i) 0)
                true (factor-helper {n / i} i [*factors i])
                (factor-helper n {i + 1} factors)
            )
        ))
        (factor-helper n 2 [])
    ))
    _ (print "prime-factors(60):" (prime-factors 60))
    _ (print "prime-factors(100):" (prime-factors 100))
    
    nil
)
