package runtime_core

var CoreStack Stack

func init() {
	CoreStack = emptyStack.Push(emptyFrame)
	CoreStack = PeekAndUpdate(CoreStack, func(f Frame) Frame {
		for _, m := range []Module{letModule, lambdaModule, matchModule} {
			f = f.Set(m.Name, m)
		}
		return f
	})
}
