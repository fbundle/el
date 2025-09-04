package ts

import (
	"el/ast"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

func Unmarshal(typeMake string, e ast.Expr) adt.Result[Sort] {
	panic("not implemented")
}

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
