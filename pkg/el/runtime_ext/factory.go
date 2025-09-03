package runtime_ext

import (
	"el/pkg/el/runtime_core"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Runtime = runtime_core.Runtime
type Stack = runtime_core.Stack

func NewBasicRuntime() (Runtime, Stack) {
	r := Runtime{
		MaxStackDepth: 1000,
		ParseLiteralOpt: func(lit string) adt.Result[Value] {
			val, err := parseLiteral(lit)
			return adt.Result[Value]{
				Val: val,
				Err: err,
			}
		},
		PostProcessArgsOpt: func(args []Value) adt.Result[[]Value] {
			unwrappedArgs, err := unwrapArgs(args)
			return adt.Result[[]Value]{
				Val: unwrappedArgs,
				Err: err,
			}
		},
	}
	sh := stackHelper{stack: runtime_core.NewBuiltinStack()}
	sh = sh.LoadExtension(listExtension, lenExtension, sliceExtension)
	sh = sh.Load("true", True).Load("false", False)
	sh = sh.LoadExtension(eqExtension, neExtension, ltExtension, leExtension, gtExtension, geExtension)
	sh = sh.LoadExtension(addExtension, subExtension, mulExtension, divExtension, modExtension)

	return r, sh.stack
}

type stackHelper struct {
	stack Stack
}

func (sh stackHelper) Load(name Name, value Value) stackHelper {
	stack := runtime_core.UpdateHead(sh.stack, func(frame runtime_core.Frame) runtime_core.Frame {
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
