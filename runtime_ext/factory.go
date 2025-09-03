package runtime_ext

import (
	runtime "el/runtime"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Runtime = runtime.Runtime
type Frame = runtime.Frame

func NewBasicRuntime() (Runtime, Frame) {
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
	f :=
		(&frameHelper{frame: runtime.BuiltinFrame}).
			LoadExtension(listExtension, lenExtension, sliceExtension, rangeExtension).
			Load("true", True).Load("false", False).
			LoadExtension(eqExtension, neExtension, ltExtension, leExtension, gtExtension, geExtension).
			LoadExtension(addExtension, subExtension, mulExtension, divExtension, modExtension).frame

	return r, f
}

type frameHelper struct {
	frame Frame
}

func (sh *frameHelper) Load(name Name, value Object) *frameHelper {
	sh.frame = sh.frame.Set(name, value)
	return sh
}

func (sh *frameHelper) LoadExtension(exts ...Extension) *frameHelper {
	for _, ext := range exts {
		sh.Load(ext.Name, ext.Module())
	}
	return sh
}
