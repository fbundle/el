(let
    # Pattern Matching and Conditionals
    _ (print "=== Pattern Matching ===")

    # Basic pattern matching
    _ (print "Match 1 == 1:" (match 1 1 "yes" "no"))
    _ (print "Match 1 == 2:" (match 1 2 "yes" "no"))

    # Multiple conditions
    _ (let
        grade_check (lambda score (match score
            100 "Perfect!"
            90 "Excellent"
            80 "Good"
            70 "Average"
            60 "Below Average"
            "Fail"
        ))
        _ (print "Grade for 95:" (grade_check 95))
        _ (print "Grade for 75:" (grade_check 75))
        _ (print "Grade for 45:" (grade_check 45))
        nil
    )

    # Boolean operations
    _ (print "True and True:" (and true true))
    _ (print "True and False:" (and true false))
    _ (print "False or True:" (or false true))
    _ (print "Not True:" (not true))

    # Conditional expressions
    _ (print "If 5 > 3:" (if {5 > 3} "greater" "not greater"))
    _ (print "If 2 > 5:" (if {2 > 5} "greater" "not greater"))

    # Complex pattern matching with calculations
    _ (let
        abs (lambda x (match {x < 0}
            true {-x}
            x
        ))
        _ (print "Absolute value of -5:" (abs -5))
        _ (print "Absolute value of 5:" (abs 5))
        nil
    )

    # String pattern matching
    _ (let
        greet (lambda name (match name
            "Alice" "Hello Alice!"
            "Bob" "Hi Bob!"
            "Hello stranger!"
        ))
        _ (print "Greet Alice:" (greet "Alice"))
        _ (print "Greet Bob:" (greet "Bob"))
        _ (print "Greet Charlie:" (greet "Charlie"))
        nil
    )

    # List pattern matching
    _ (let
        list_length (lambda lst (match (len lst)
            0 0
            (let
                rest_lst (rest lst)
                {1 + (list_length rest_lst)}
            )
        ))
        _ (print "Length of [1 2 3 4]:" (list_length [1 2 3 4]))
        _ (print "Length of []:" (list_length []))
        nil
    )

    # Type checking with pattern matching
    _ (let
        type_check (lambda x (match (type x)
            "number" "It's a number"
            "string" "It's a string"
            "list" "It's a list"
            "function" "It's a function"
            "unknown type"
        ))
        _ (print "Type of 42:" (type_check 42))
        _ (print "Type of 'hello':" (type_check "hello"))
        _ (print "Type of [1 2 3]:" (type_check [1 2 3]))
        nil
    )

    nil
)