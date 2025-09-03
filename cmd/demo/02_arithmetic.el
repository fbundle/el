# Arithmetic Operations Demo
# Demonstrates basic arithmetic and comparison operations

(let
    # Basic arithmetic
    _ (print "=== Basic Arithmetic ===")
    _ (print "1 + 2 =" {1 + 2})
    _ (print "10 - 3 =" {10 - 3})
    _ (print "4 * 5 =" (mul 4 5))
    _ (print "20 / 4 =" (div 20 4))
    _ (print "17 % 5 =" (mod 17 5))
    
    # Complex expressions with operator precedence
    _ (print "\n=== Complex Expressions ===")
    _ (print "2 + 3 * 4 =" (add 2 (mul 3 4)))
    _ (print "(2 + 3) * 4 =" (mul (add 2 3) 4))
    _ (print "10 - 2 * 3 + 1 =" (add (sub 10 (mul 2 3)) 1))
    _ (print "2^3 + 4*5 - 6/2 =" (sub (add (mul 2 (mul 2 2)) (mul 4 5)) (div 6 2)))
    
    # Comparisons
    _ (print "\n=== Comparisons ===")
    _ (print "5 == 5:" (eq 5 5))
    _ (print "5 != 3:" (ne 5 3))
    _ (print "3 < 5:" (lt 3 5))
    _ (print "5 <= 5:" (le 5 5))
    _ (print "7 > 3:" (gt 7 3))
    _ (print "4 >= 4:" (ge 4 4))
    
    # Variables and calculations
    _ (print "\n=== Variables and Calculations ===")
    x 10
    y 20
    z (add x (mul y 2))
    _ (print "x =" x)
    _ (print "y =" y)
    _ (print "z = x + y * 2 =" z)
    
    nil
)
