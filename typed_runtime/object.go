package runtime

import (
	"el/sorts"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Data interface {
	String() string
}

type Object interface {
	Data() Data
	String() string
	Sort() Sort
	Type() Object
	Cast(dtype Object) adt.Option[Object]
}

var BuiltinType = makeType("builtin_type")

var NilType = makeType(sorts.Unit)
var AnyType = makeType(sorts.Any)
