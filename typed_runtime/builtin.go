package runtime

import (
	"context"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Exec = func(r Runtime, ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[Object]

type FuncData struct {
	exec Exec
	repr string
}

func (f FuncData) String() string {
	return f.repr
}
