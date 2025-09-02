package runtime_core

func NewRuntime() Stack {
	stack := emptyStack.Push(emptyFrame)
	return LoadModule(stack, letModule, lambdaModule, matchModule)
}
