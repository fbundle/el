# Type System and Introspection
_ (print "=== Type System ===")

# Basic type checking
_ (print "Type of 42:" (type 42))
_ (print "Type of 'hello':" (type "hello"))
_ (print "Type of [1 2 3]:" (type [1 2 3]))
_ (print "Type of true:" (type true))
_ (print "Type of nil:" (type nil))

# Function type
_ (let
    my_func (lambda x {x + 1})
    _ (print "Type of function:" (type my_func))
    nil
)

# Type of type
_ (print "Type of type function:" (type type))

# Type introspection
type_info (lambda x (let
    t (type x)
    (list x t (type t))
))
(print "Type info for 42:" (type_info 42))
(print "Type info for 'hello':" (type_info "hello"))

# Type checking functions
is_number (lambda x (match (type x)
    "number" true
    false
))
is_string (lambda x (match (type x)
    "string" true
    false
))
is_list (lambda x (match (type x)
    "list" true
    false
))
is_function (lambda x (match (type x)
    "function" true
    false
))

(print "Is 42 a number?" (is_number 42))
(print "Is 'hello' a string?" (is_string "hello"))
(print "Is [1 2 3] a list?" (is_list [1 2 3]))
(print "Is lambda a function?" (is_function (lambda x x)))

# Type-safe operations
safe_add (lambda a b (let
    type_a (type a)
    type_b (type b)
    (match type_a
        "number" (match type_b
            "number" {a + b}
            "Cannot add number and non-number"
        )
        "Cannot add non-number"
    )
))

(print "Safe add 5 + 3:" (safe_add 5 3))
(print "Safe add 5 + 'hello':" (safe_add 5 "hello"))

# Type conversion (conceptual)
to_string (lambda x (match (type x)
    "number" (let
        # Simple number to string conversion
        (match x
            0 "0"
            1 "1"
            2 "2"
            3 "3"
            4 "4"
            5 "5"
            6 "6"
            7 "7"
            8 "8"
            9 "9"
            "number"
        )
    )
    "string"
))

(print "Convert 5 to string:" (to_string 5))

# Polymorphic functions
identity (lambda x x)
(print "Identity of 42:" (identity 42))
(print "Identity of 'hello':" (identity "hello"))
(print "Identity of [1 2 3]:" (identity [1 2 3]))

# Type-safe list operations
safe_head (lambda lst (match (type lst)
    "list" (match (len lst)
        0 nil
        (head lst)
    )
    nil
))

(print "Safe head of [1 2 3]:" (safe_head [1 2 3]))
(print "Safe head of []:" (safe_head []))
(print "Safe head of 'not a list':" (safe_head "not a list"))

# Type checking with pattern matching
process_value (lambda x (match (type x)
    "number" (let
        doubled {x * 2}
        (list "number" doubled)
    )
    "string" (let
        length (len x)
        (list "string" length)
    )
    "list" (let
        length (len x)
        (list "list" length)
    )
    (list "unknown" x)
))

(print "Process 42:" (process_value 42))
(print "Process 'hello':" (process_value "hello"))
(print "Process [1 2 3]:" (process_value [1 2 3]))
(print "Process true:" (process_value true))

# Type hierarchy simulation
is_primitive (lambda x (match (type x)
    "number" true
    "string" true
    "boolean" true
    false
))

is_composite (lambda x (match (type x)
    "list" true
    "function" true
    false
))

(print "Is 42 primitive?" (is_primitive 42))
(print "Is [1 2 3] composite?" (is_composite [1 2 3]))
(print "Is lambda primitive?" (is_primitive (lambda x x)))

# Type-safe arithmetic
arithmetic_result (lambda op a b (let
    type_a (type a)
    type_b (type b)
    (match type_a
        "number" (match type_b
            "number" (match op
                "+" {a + b}
                "-" {a - b}
                "*" {a * b}
                "/" (match {b == 0}
                    true "Division by zero"
                    {a / b}
                )
                "Invalid operator"
            )
            "Cannot perform arithmetic with non-number"
        )
        "Cannot perform arithmetic with non-number"
    )
))

(print "5 + 3:" (arithmetic_result "+" 5 3))
(print "10 - 4:" (arithmetic_result "-" 10 4))
(print "6 * 7:" (arithmetic_result "*" 6 7))
(print "15 / 3:" (arithmetic_result "/" 15 3))
_ (print "10 / 0:" (arithmetic_result "/" 10 0))
_ (print "5 + 'hello':" (arithmetic_result "+" 5 "hello"))

nil
