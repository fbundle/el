package runtime

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/ordered_map"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/stack"
)

type Name string

// Frame - a collection of bindings Name -> Object
type Frame = ordered_map.OrderedMap[Name, Object]

// Stack - stack frame
type Stack = stack.Stack[Frame]

// updateHead - update top of the stack
func updateHead(s Stack, f func(Frame) Frame) Stack {
	h := s.Peek()
	h = f(h)
	return s.Pop().Push(h)
}

func collapseStack(s Stack) Frame {
	head := Frame{}
	for _, frame := range s.Iter {
		for k, v := range frame.Iter {
			if _, ok := head.Get(k); !ok {
				head = head.Set(k, v)
			}
		}
	}
	return head
}

func findStack(s Stack, name Name) adt.Option[Object] {
	for _, frame := range s.Iter {
		if o, ok := frame.Get(name); ok {
			return adt.Some[Object](o)
		}
	}
	return adt.None[Object]()
}
