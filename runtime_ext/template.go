package runtime_ext

import "fmt"

// WithTemplate - add some common template to the code
func WithTemplate(s string) string {
	return fmt.Sprintf(`
(let

# identity - identity function
unit (lambda x x) 

# get - get element from list
get (lambda l i (unit * (slice l (range i (add i 1)))))			# get l[i]
head (lambda l (get l 0))							# get l[0]
rest (lambda l (slice l (range 1 (len l))))			# get l[1:]

# operators
+ add - sub x mul / div %% mod			# short hand for common operator
== eq != ne <= le < lt > gt >= ge

# map
map (lambda l f (match (len l)
	0 []					# if len l == 0 then return empty list
	(let
		first_elem (head l)
		first_elem2 (f first_elem)
		rest_elems (rest l)
		rest_elems2 (map rest_elems f)	# recursive call
		(list first_elem2 *rest_elems2)
	)
))

%s

)`, s)
}
