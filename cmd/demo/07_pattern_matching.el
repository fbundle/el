# Pattern Matching Demo
# Demonstrates pattern matching and conditional logic

(let
    # Simple pattern matching
    _ (print "=== Simple Pattern Matching ===")
    classify (lambda n (match n
        0 "zero"
        1 "one"
        2 "two"
        3 "three"
        "many"
    ))
    _ (print "classify(0):" (classify 0))
    _ (print "classify(1):" (classify 1))
    _ (print "classify(2):" (classify 2))
    _ (print "classify(3):" (classify 3))
    _ (print "classify(5):" (classify 5))
    
    # Boolean pattern matching
    _ (print "\n=== Boolean Pattern Matching ===")
    sign (lambda n (match {n < 0}
        true "negative"
        (match {n > 0}
            true "positive"
            "zero"
        )
    ))
    _ (print "sign(-5):" (sign -5))
    _ (print "sign(0):" (sign 0))
    _ (print "sign(5):" (sign 5))
    
    # Grade classification
    _ (print "\n=== Grade Classification ===")
    grade (lambda score (match {score >= 90}
        true "A"
        (match {score >= 80}
            true "B"
            (match {score >= 70}
                true "C"
                (match {score >= 60}
                    true "D"
                    "F"
                )
            )
        )
    ))
    _ (print "grade(95):" (grade 95))
    _ (print "grade(85):" (grade 85))
    _ (print "grade(75):" (grade 75))
    _ (print "grade(65):" (grade 65))
    _ (print "grade(55):" (grade 55))
    
    # List length pattern matching
    _ (print "\n=== List Length Pattern Matching ===")
    describe-list (lambda lst (match (len lst)
        0 "empty"
        1 "single"
        2 "pair"
        3 "triple"
        "many"
    ))
    _ (print "describe-list([]):" (describe-list []))
    _ (print "describe-list([1]):" (describe-list [1]))
    _ (print "describe-list([1 2]):" (describe-list [1 2]))
    _ (print "describe-list([1 2 3]):" (describe-list [1 2 3]))
    _ (print "describe-list([1 2 3 4]):" (describe-list [1 2 3 4]))
    
    # Parity checking
    _ (print "\n=== Parity Checking ===")
    parity (lambda n (match {n % 2}
        0 "even"
        "odd"
    ))
    _ (print "parity(4):" (parity 4))
    _ (print "parity(7):" (parity 7))
    _ (print "parity(0):" (parity 0))
    _ (print "parity(1):" (parity 1))
    
    # Age category
    _ (print "\n=== Age Category ===")
    age-category (lambda age (match {age < 13}
        true "child"
        (match {age < 20}
            true "teen"
            (match {age < 65}
                true "adult"
                "senior"
            )
        )
    ))
    _ (print "age-category(8):" (age-category 8))
    _ (print "age-category(16):" (age-category 16))
    _ (print "age-category(30):" (age-category 30))
    _ (print "age-category(70):" (age-category 70))
    
    # Temperature classification
    _ (print "\n=== Temperature Classification ===")
    temp-class (lambda temp (match {temp < 0}
        true "freezing"
        (match {temp < 20}
            true "cold"
            (match {temp < 30}
                true "warm"
                "hot"
            )
        )
    ))
    _ (print "temp-class(-5):" (temp-class -5))
    _ (print "temp-class(10):" (temp-class 10))
    _ (print "temp-class(25):" (temp-class 25))
    _ (print "temp-class(35):" (temp-class 35))
    
    # Multiple condition matching
    _ (print "\n=== Multiple Condition Matching ===")
    triangle-type (lambda a b c (match {a + b > c}
        true (match {a == b}
            true (match {b == c}
                true "equilateral"
                "isosceles"
            )
            (match {b == c}
                true "isosceles"
                (match {a == c}
                    true "isosceles"
                    "scalene"
                )
            )
        )
        "invalid"
    ))
    _ (print "triangle-type(3, 3, 3):" (triangle-type 3 3 3))
    _ (print "triangle-type(3, 3, 4):" (triangle-type 3 3 4))
    _ (print "triangle-type(3, 4, 5):" (triangle-type 3 4 5))
    
    nil
)
