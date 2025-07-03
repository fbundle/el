package el

type Frame = map[string]Object

func newFrameStack() Stack {
	return &frameStack{
		stack: []Frame{
			{},
		},
	}
}

// Stack - read-only stack
type Stack interface {
	Push(Frame) Stack
	Pop() (Stack, Frame)
	Depth() uint
}

type frameStack struct {
	stack []Frame
}

func (s *frameStack) Push(frame Frame) Stack {
	return &frameStack{
		stack: append(s.stack, frame),
	}
}

func (s *frameStack) Pop() (Stack, Frame) {
	return &frameStack{
		stack: s.stack[:len(s.stack)-1],
	}, s.stack[len(s.stack)-1]
}

func (s *frameStack) Depth() uint {
	return uint(len(s.stack))
}
