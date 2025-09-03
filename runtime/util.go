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

type cmd struct {
	cmdExpr ast.Expr
	argList []ast.Expr
}

func getCmd(e ast.Lambda) adt.Option[cmd] {
	if len(e) == 0 {
		return adt.None[cmd]()
	}
	return adt.Some(cmd{
		cmdExpr: e[0],
		argList: e[1:],
	})
}
