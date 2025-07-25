package runtime

type Name string

// Frame - a collection of bindings NameExpr -> Object
type Frame = map[Name]Object

// FrameStack - stack of frames
type FrameStack interface {
	Push(Frame)
	Pop() Frame
	Depth() uint
	Iter(func(Frame) bool)
}

func newFrameStack() FrameStack {
	s := &frameStack{
		stack: nil,
	}
	s.Push(Frame{})
	return s
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
