package runtime

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/ordered_map"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/stack"
)

type Name string

// Frame - a collection of bindings NameExpr -> Object
type Frame = ordered_map.OrderedMap[Name, Object]

type FrameStack = stack.Stack[Frame]

func newFrameStack() FrameStack {
	return stack.Empty[Frame]().Push(Frame{})
}

func updateHead(s FrameStack, f func(Frame) Frame) FrameStack {
	s, h := s.Pop()
	h = f(h)
	return s.Push(h)
}
