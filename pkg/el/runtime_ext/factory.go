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
		ParseLiteralOpt: func(lit string) adt.Option[runtime_core.Object] {
			return adt.Wrap(func() (Object, error) {
				return parseLiteral(lit)
			})()
		},
		UnwrapArgsOpt: func(args []runtime_core.Object) adt.Option[[]runtime_core.Object] {
			return adt.Wrap(func() ([]Object, error) {
				return unwrapArgs(args)
			})()
		},
	}
	s := runtime_core.Stack{}.Push(runtime_core.Frame{})
	s = loadModule(s, letModule, lambdaModule, matchModule)
	s = loadExtension(s, listExtension, lenExtension, sliceExtension)
	return r, s
}

func loadExtension(s Stack, exts ...Extension) Stack {
	return runtime_core.PeekAndUpdate(s, func(frame runtime_core.Frame) runtime_core.Frame {
		for _, ext := range exts {
			frame = frame.Set(ext.Name, ext.Module())
		}
		return frame
	})
}

func loadModule(s Stack, ms ...Module) Stack {
	return runtime_core.PeekAndUpdate(s, func(frame runtime_core.Frame) runtime_core.Frame {
		for _, m := range ms {
			frame = frame.Set(m.name, m)
		}
		return frame
	})
}
