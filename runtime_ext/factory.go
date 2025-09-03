package runtime_ext

import (
	runtime "el/runtime"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Runtime = runtime.Runtime
type Stack = runtime.Stack

func NewBasicRuntime() (Runtime, Stack) {
	r := Runtime{
		MaxStackDepth: 1000,
		ParseLiteral: func(lit string) adt.Result[Object] {
			val, err := parseLiteral(lit)
			return adt.Result[Object]{
				Val: val,
				Err: err,
			}
		},
		UnwrapArgs: func(argsOpt adt.Result[[]Object]) adt.Result[[]Object] {
			var args []Object
			if err := argsOpt.Unwrap(&args); err != nil {
				return adt.Result[[]Object]{
					Err: err,
				}
			}
			unwrappedArgs, err := unwrapArgs(args)
			return adt.Result[[]Object]{
				Val: unwrappedArgs,
				Err: err,
			}
		},
	}
	s :=
		(&stackHelper{stack: runtime.BuiltinStack}).
			LoadExtension(listExtension, lenExtension, sliceExtension, rangeExtension).
			Load("true", True).Load("false", False).
			LoadExtension(eqExtension, neExtension, ltExtension, leExtension, gtExtension, geExtension).
			LoadExtension(addExtension, subExtension, mulExtension, divExtension, modExtension).stack

	return r, s
}

type stackHelper struct {
	stack Stack
}

func (sh *stackHelper) Load(name Name, value Object) *stackHelper {
	head := sh.stack.Peek()
	sh.stack = sh.stack.Pop()
	head = head.Set(name, value)
	sh.stack = sh.stack.Push(head)
	return sh
}

func (sh *stackHelper) LoadExtension(exts ...Extension) *stackHelper {
	for _, ext := range exts {
		sh.Load(ext.Name, ext.Module())
	}
	return sh
}
