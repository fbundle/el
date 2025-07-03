package el

type Frame = map[NameExpr]Object

func newFrameStack() Stack {
	return &frameStack{
		stack: []Frame{
			{},
		},
	}
}

// Stack - call stack
type Stack interface {
	Push(Frame)
	Pop() Frame
	Depth() uint
	Iter(func(Frame) bool)
}

type frameStack struct {
	stack []Frame
}

func (s *frameStack) Push(frame Frame) {
	s.stack = append(s.stack, frame)
}

func (s *frameStack) Pop() Frame {
	head := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return head
}

func (s *frameStack) Depth() uint {
	return uint(len(s.stack))
}

func (s *frameStack) Iter(f func(Frame) bool) {
	for i := len(s.stack) - 1; i >= 0; i-- {
		if !f(s.stack[i]) {
			break
		}
	}
}
