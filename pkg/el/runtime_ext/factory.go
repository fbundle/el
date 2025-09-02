package runtime_ext

import (
	"el/pkg/el/runtime_core"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Runtime = runtime_core.Runtime

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
