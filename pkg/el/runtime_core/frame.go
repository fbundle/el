package runtime_core

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/ordered_map"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/stack"
)

type Name string

// Frame - a collection of bindings Name -> Value
type Frame = ordered_map.OrderedMap[Name, Value]

// Stack - stack frame
type Stack = stack.Stack[Frame]

// PeekAndUpdate - update top of the stack
func PeekAndUpdate(s Stack, f func(Frame) Frame) Stack {
	h := s.Peek()
	h = f(h)
	return s.Pop().Push(h)
}

func searchOnStack(s Stack, name Name) (Value, bool) {
	for _, frame := range s.Iter {
		if o, ok := frame.Get(name); ok {
			return o, true
		}
	}
	return nil, false
}
