package runtime_ext

import (
	"el/pkg/el/runtime_core"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Runtime = runtime_core.Runtime
type Stack = runtime_core.Stack

var InitRuntime = Runtime{
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

var InitStack Stack

func init() {
	InitStack = runtime_core.InitStack
	InitStack = runtime_core.PeekAndUpdate(InitStack, func(frame runtime_core.Frame) runtime_core.Frame {
		var m runtime_core.Module
		m = runtime_core.MakeModuleFromExtension(listExtension)
		frame = frame.Set(m.Name, m)
		m = runtime_core.MakeModuleFromExtension(lenExtension)
		frame = frame.Set(m.Name, m)
		m = runtime_core.MakeModuleFromExtension(sliceExtension)
		frame = frame.Set(m.Name, m)
		return frame
	})
}
