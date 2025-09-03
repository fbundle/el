package runtime

import (
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
