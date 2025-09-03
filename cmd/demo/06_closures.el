# Closures Demo
# Demonstrates closure behavior and lexical scoping

(let
    # Simple counter closure
    _ (print "=== Simple Counter Closure ===")
    make-counter (lambda start (lambda {start + 1}))
    counter1 (make-counter 10)
    counter2 (make-counter 100)
    _ (print "counter1 (start=10):" (counter1))
    _ (print "counter1 again:" (counter1))
    _ (print "counter2 (start=100):" (counter2))
    _ (print "counter2 again:" (counter2))
    
    # Adder closure
    _ (print "\n=== Adder Closure ===")
    make-adder (lambda n (lambda x {x + n}))
    add-5 (make-adder 5)
    add-10 (make-adder 10)
    add-100 (make-adder 100)
    _ (print "add-5(3) =" (add-5 3))
    _ (print "add-5(7) =" (add-5 7))
    _ (print "add-10(3) =" (add-10 3))
    _ (print "add-100(50) =" (add-100 50))
    
    # Multiplier closure
    _ (print "\n=== Multiplier Closure ===")
    make-multiplier (lambda n (lambda x {x * n}))
    double (make-multiplier 2)
    triple (make-multiplier 3)
    times-10 (make-multiplier 10)
    _ (print "double(5) =" (double 5))
    _ (print "triple(4) =" (triple 4))
    _ (print "times-10(7) =" (times-10 7))
    
    # Accumulator closure (stateful)
    _ (print "\n=== Accumulator Closure ===")
    make-accumulator (lambda initial (lambda x {initial + x}))
    acc1 (make-accumulator 0)
    acc2 (make-accumulator 100)
    _ (print "acc1(5):" (acc1 5))
    _ (print "acc1(3):" (acc1 3))
    _ (print "acc1(10):" (acc1 10))
    _ (print "acc2(20):" (acc2 20))
    _ (print "acc2(30):" (acc2 30))
    
    # Function factory
    _ (print "\n=== Function Factory ===")
    make-comparator (lambda op (lambda x y (op x y)))
    greater-than (make-comparator gt)
    less-than (make-comparator lt)
    _ (print "greater-than(5, 3):" (greater-than 5 3))
    _ (print "greater-than(2, 7):" (greater-than 2 7))
    _ (print "less-than(3, 8):" (less-than 3 8))
    
    # Nested closures
    _ (print "\n=== Nested Closures ===")
    make-math-ops (lambda base (let
        add-base (lambda x {x + base})
        mul-base (lambda x {x * base})
        [add-base mul-base]
    ))
    ops (make-math-ops 5)
    add-5 (slice ops [0])
    mul-5 (slice ops [1])
    _ (print "add-5(3):" (add-5 3))
    _ (print "mul-5(4):" (mul-5 4))
    
    # Closure with multiple captured variables
    _ (print "\n=== Multiple Captured Variables ===")
    make-linear (lambda a b (lambda x {a * x + b}))
    line1 (make-linear 2 3)  # y = 2x + 3
    line2 (make-linear -1 10)  # y = -x + 10
    _ (print "line1(0):" (line1 0))  # should be 3
    _ (print "line1(2):" (line1 2))  # should be 7
    _ (print "line2(5):" (line2 5))  # should be 5
    _ (print "line2(10):" (line2 10))  # should be 0
    
    nil
)
