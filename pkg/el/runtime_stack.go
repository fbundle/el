package el

import (
	"el/pkg/stack"
)

type Frame = map[string]Object

func newFrameStack() Stack {
	return &frameStack{
		stack: stack.Empty[Frame]().Push(Frame{}),
	}
}

// Stack - read-only stack
type Stack interface {
	Push(Frame) Stack
	Pop() (Stack, Frame)
	Depth() uint
}

type frameStack struct {
	stack stack.Node[Frame]
}

func (s *frameStack) Push(frame Frame) Stack {
	return &frameStack{
		stack: s.stack.Push(frame),
	}
}

func (s *frameStack) Pop() (Stack, Frame) {
	stack1, frame := s.stack.Pop()
	return &frameStack{
		stack: stack1,
	}, frame
}

func (s *frameStack) Depth() uint {
	return s.stack.Depth()
}
