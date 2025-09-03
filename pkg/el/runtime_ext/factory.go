package runtime_ext

import (
	"el/pkg/el/runtime"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Runtime = runtime.Runtime
type Stack = runtime.Stack

func NewBasicRuntime() (Runtime, Stack) {
	r := Runtime{
		MaxStackDepth: 1000,
		ParseLiteral: func(lit string) adt.Result[Value] {
			val, err := parseLiteral(lit)
			return adt.Result[Value]{
				Val: val,
				Err: err,
			}
		},
		UnwrapArgs: func(argsOpt adt.Result[[]Value]) adt.Result[[]Value] {
			var args []Value
			if err := argsOpt.Unwrap(&args); err != nil {
				return adt.Result[[]Value]{
					Err: err,
				}
			}
			unwrappedArgs, err := unwrapArgs(args)
			return adt.Result[[]Value]{
				Val: unwrappedArgs,
				Err: err,
			}
		},
	}
	sh := stackHelper{stack: runtime.NewBuiltinStack()}
	sh = sh.LoadExtension(listExtension, lenExtension, sliceExtension, rangeExtension)
	sh = sh.Load("true", True).Load("false", False)
	sh = sh.LoadExtension(eqExtension, neExtension, ltExtension, leExtension, gtExtension, geExtension)
	sh = sh.LoadExtension(addExtension, subExtension, mulExtension, divExtension, modExtension)

	return r, sh.stack
}

type stackHelper struct {
	stack Stack
}

func (sh stackHelper) Load(name Name, value Value) stackHelper {
	stack := runtime.UpdateHead(sh.stack, func(frame runtime.Frame) runtime.Frame {
		return frame.Set(name, value)
	})
	return stackHelper{stack: stack}
}

func (sh stackHelper) LoadExtension(exts ...Extension) stackHelper {
	for _, ext := range exts {
		sh = sh.Load(ext.Name, ext.Module())
	}
	return sh
}
