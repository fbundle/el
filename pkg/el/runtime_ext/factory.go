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
	s := loadExtension(
		runtime_core.NewBuiltinStack(),
		listExtension, lenExtension, sliceExtension,
	)
	return r, s
}

func loadExtension(s Stack, exts ...Extension) Stack {
	return runtime_core.UpdateHead(s, func(frame runtime_core.Frame) runtime_core.Frame {
		for _, ext := range exts {
			frame = frame.Set(ext.Name, ext.Module())
		}
		return frame
	})
}
