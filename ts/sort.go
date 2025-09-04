package ts

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

const (
	DefaultSortName = "type"
)

type Data interface {
	String() string
}

type Sort interface {
	Data() adt.Option[Data]
	Level() int
	String() string
	Type() Sort
	Cast(parent Sort) adt.Option[Sort]
	Len() int

	le(dst Sort) bool
	prepend(param Sort) Sort
}
