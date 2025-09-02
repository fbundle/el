package runtime

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/ordered_map"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/stack"
)

type Name string

// Frame - a collection of bindings Name -> Object
type Frame = ordered_map.OrderedMap[Name, Object]

type FrameStack = stack.Stack[Frame]

func newFrameStack() FrameStack {
	return stack.Empty[Frame]().Push(ordered_map.EmptyOrderedMap[Name, Object]())
}

func updateHead(s FrameStack, f func(Frame) Frame) FrameStack {
	h := s.Peek()
	h = f(h)
	return s.Pop().Push(h)
}
