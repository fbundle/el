package stack

// Node - implementation of read-only stack
// useful for program stack since we can make a clone of stack without copying
type Node[T any] interface {
	Push(T) Node[T]
	Pop() (Node[T], T)
	Depth() uint
}

func Empty[T any]() Node[T] {
	return &node[T]{
		depth: uint(0),
	}
}

type node[T any] struct {
	depth uint
	value T
	next  *node[T]
}

func (n *node[T]) Push(value T) Node[T] {
	return &node[T]{
		depth: n.depth + 1,
		value: value,
		next:  n,
	}
}

func (n *node[T]) Pop() (Node[T], T) {
	if n.depth == 0 {
		panic(n)
	}
	return n.next, n.value
}

func (n *node[T]) Depth() uint {
	return n.depth
}
