package runtime

import (
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

type Nil struct{}

func (Nil) String() string {
	return "nil"
}

var NilType Object = MakeType("nil_type")
