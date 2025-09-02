package runtime_core

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/ordered_map"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/stack"
)

type Name string

// Frame - a collection of bindings Name -> Object
type Frame = ordered_map.OrderedMap[Name, Object]

// Stack - stack frame
type Stack = stack.Stack[Frame]

var emptyFrame = ordered_map.EmptyOrderedMap[Name, Object]()

var emptyStack = stack.Empty[Frame]()

// PeekAndUpdate - update top of the stack
func PeekAndUpdate(s Stack, f func(Frame) Frame) Stack {
	h := s.Peek()
	h = f(h)
	return s.Pop().Push(h)
}

var InitStack Stack

func init() {
	InitStack = emptyStack.Push(emptyFrame)
	InitStack = PeekAndUpdate(InitStack, func(f Frame) Frame {
		for _, m := range []Module{letModule, lambdaModule, matchModule} {
			f = f.Set(m.Name, m)
		}
		return f
	})
}
