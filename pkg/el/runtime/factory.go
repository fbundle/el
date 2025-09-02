package runtime

// NewCoreRuntime - Runtime and core control flow extensions
func NewCoreRuntime() *Runtime {
	return (&Runtime{
		Stack: newFrameStack(),
	}).LoadModule(letModule, lambdaModule, matchModule)
}

// NewBasicRuntime - NewCoreRuntime and minimal set of arithmetic extensions for Turing completeness
func NewBasicRuntime() *Runtime {
	return NewCoreRuntime().
		// list extension
		LoadExtension(listExtension, lenExtension, rangeExtension, sliceExtension).
		// arithmetic extension
		LoadConstant("true", True).LoadConstant("false", False).
		LoadExtension(eqExtension, neExtension, ltExtension, leExtension, gtExtension, geExtension).
		LoadExtension(addExtension, subExtension, mulExtension, divExtension, modExtension)
	// extra
}
