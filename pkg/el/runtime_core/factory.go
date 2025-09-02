package runtime_core

func NewCoreStack() Stack {
	stack := emptyStack.Push(emptyFrame)
	return PeekAndUpdate(stack, func(f Frame) Frame {
		for _, m := range []Module{letModule, lambdaModule, matchModule} {
			f = f.Set(m.Name, m)
		}
		return f
	})
}
