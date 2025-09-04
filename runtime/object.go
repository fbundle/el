package runtime

import (
	"context"
	"el/ast"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Exec = func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object]

type Object interface {
	String() string
}

type Function struct {
	exec Exec
	repr string
}

func (f Function) String() string {
	return f.repr
}

type Nil struct{}

func (Nil) String() string {
	return "nil"
}
