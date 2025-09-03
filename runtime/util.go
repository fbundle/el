package runtime

import (
	"el/ast"
	"errors"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

// helpers
func value(o Object) adt.Result[Object] {
	return adt.Ok[Object](o)
}

func errValue(err error) adt.Result[Object] {
	return adt.Err[Object](err)
}
func errValueString(msg string) adt.Result[Object] {
	return errValue(errors.New(msg))
}

func getCmd(e ast.Node) adt.Option2[ast.Expr, []ast.Expr] {
	if len(e) == 0 {
		return adt.None2[ast.Expr, []ast.Expr]()
	}
	cmd := e[0]
	args := e[1:]
	return adt.Some2[ast.Expr, []ast.Expr](cmd, args)
}
