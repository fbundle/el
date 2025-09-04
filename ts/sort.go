package ts

import (
	"el/ast"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

const (
	TypeName = "type"
	Arrow    = "->"
)

func Unmarshal(typeMake string, e ast.Expr) adt.Result[Sort] {
	panic("not implemented")
}

type Sort interface {
	Level() int
	String() string
	Type() Sort
	Cast(sort Sort) adt.Option[Sort]
	Chain() adt.NonEmptySlice[Sort]

	le(dst Sort) bool
	prepend(param Sort) Sort
}
