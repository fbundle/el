package stack

func PopAtMost[T any](s Node[T], cap int) (Node[T], []T) {
	a := make([]T, 0, cap)
	var v T
	for len(a) < cap && s.Depth() > 0 {
		s, v = s.Pop()
		a = append(a, v)
	}
	return s, a
}
