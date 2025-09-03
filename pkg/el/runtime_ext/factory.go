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
		ParseLiteralOpt: func(lit string) adt.Result[runtime_core.Value] {
			return adt.Wrap(func() (Value, error) {
				return parseLiteral(lit)
			})()
		},
		PostProcessArgsOpt: func(args []runtime_core.Value) adt.Result[[]runtime_core.Value] {
			return adt.Wrap(func() ([]Value, error) {
				return unwrapArgs(args)
			})()
		},
	}
	s := loadExtension(runtime_core.NewBuiltinStack(), listExtension, lenExtension, sliceExtension)
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
