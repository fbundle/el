# Type System Demo
# Demonstrates type introspection and type checking

(let
    # Type introspection
    _ (print "=== Type Introspection ===")
    x 42
    y "hello"
    z [1 2 3]
    w true
    v nil
    f (lambda x {x + 1})
    
    _ (print "type of 42:" (type x))
    _ (print "type of 'hello':" (type y))
    _ (print "type of [1 2 3]:" (type z))
    _ (print "type of true:" (type w))
    _ (print "type of nil:" (type v))
    _ (print "type of lambda:" (type f))
    
    # Type of types
    _ (print "\n=== Type of Types ===")
    _ (print "type of (type 42):" (type (type x)))
    _ (print "type of (type (type 42)):" (type (type (type x))))
    _ (print "type of (type (type (type 42))):" (type (type (type (type x)))))
    
    # Function types
    _ (print "\n=== Function Types ===")
    add-func add
    sub-func sub
    lambda-func (lambda x {x + 1})
    _ (print "type of add function:" (type add-func))
    _ (print "type of sub function:" (type sub-func))
    _ (print "type of lambda function:" (type lambda-func))
    
    # Type checking functions
    _ (print "\n=== Type Checking Functions ===")
    is-int (lambda x (eq (type x) "int"))
    is-string (lambda x (eq (type x) "string"))
    is-list (lambda x (eq (type x) "list"))
    is-function (lambda x (eq (type x) "function"))
    is-nil (lambda x (eq (type x) "nil"))
    
    _ (print "is-int(42):" (is-int 42))
    _ (print "is-int('hello'):" (is-int "hello"))
    _ (print "is-string('world'):" (is-string "world"))
    _ (print "is-string(123):" (is-string 123))
    _ (print "is-list([1 2 3]):" (is-list [1 2 3]))
    _ (print "is-list(42):" (is-list 42))
    _ (print "is-function(add):" (is-function add))
    _ (print "is-function(42):" (is-function 42))
    _ (print "is-nil(nil):" (is-nil nil))
    _ (print "is-nil(42):" (is-nil 42))
    
    # Type-safe operations
    _ (print "\n=== Type-Safe Operations ===")
    safe-add (lambda a b (match (is-int a)
        true (match (is-int b)
            true (add a b)
            "second argument must be int"
        )
        "first argument must be int"
    ))
    _ (print "safe-add(3, 4):" (safe-add 3 4))
    _ (print "safe-add(3, 'hello'):" (safe-add 3 "hello"))
    _ (print "safe-add('world', 4):" (safe-add "world" 4))
    
    # Type conversion simulation
    _ (print "\n=== Type Conversion Simulation ===")
    to-string (lambda x (match (type x)
        "int" (let
            # Simple int to string conversion simulation
            _ (print "Converting int to string")
            "converted-string"
        )
        "string" x
        "unknown-type"
    ))
    _ (print "to-string(42):" (to-string 42))
    _ (print "to-string('hello'):" (to-string "hello"))
    _ (print "to-string([1 2 3]):" (to-string [1 2 3]))
    
    # Polymorphic function
    _ (print "\n=== Polymorphic Function ===")
    identity (lambda x x)
    _ (print "identity(42):" (identity 42))
    _ (print "identity('hello'):" (identity "hello"))
    _ (print "identity([1 2 3]):" (identity [1 2 3]))
    _ (print "identity(true):" (identity true))
    
    # Type-based dispatch
    _ (print "\n=== Type-Based Dispatch ===")
    process (lambda x (match (type x)
        "int" {x * 2}
        "string" (let
            _ (print "Processing string")
            "processed"
        )
        "list" (len x)
        "unknown"
    ))
    _ (print "process(21):" (process 21))
    _ (print "process('test'):" (process "test"))
    _ (print "process([1 2 3 4]):" (process [1 2 3 4]))
    _ (print "process(true):" (process true))
    
    # Type hierarchy simulation
    _ (print "\n=== Type Hierarchy Simulation ===")
    is-numeric (lambda x (match (type x)
        "int" true
        "string" false
        "list" false
        false
    ))
    _ (print "is-numeric(42):" (is-numeric 42))
    _ (print "is-numeric('hello'):" (is-numeric "hello"))
    _ (print "is-numeric([1 2 3]):" (is-numeric [1 2 3]))
    
    nil
)
