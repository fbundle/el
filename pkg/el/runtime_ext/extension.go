package runtime_ext

import (
	"context"
	"el/pkg/el/runtime_core"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Extension = runtime_core.Extension
type Object = runtime_core.Object

var listExtension = Extension{
	Name: "list",
	Exec: func(ctx context.Context, values ...Object) adt.Option[Object] {
		l := List{}
		for _, v := range values {
			l = append(l, v)
		}
		return adt.Some[Object](l)
	},
	Man: "module: (list 1 2 (lambda x (add x 1))) - make a list",
}
