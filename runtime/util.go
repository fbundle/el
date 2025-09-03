package runtime

import (
	"errors"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

// helpers
func value(o Value) adt.Result[Value] {
	return adt.Ok[Value](o)
}

func errValue(err error) adt.Result[Value] {
	return adt.Err[Value](err)
}
func errValueString(msg string) adt.Result[Value] {
	return errValue(errors.New(msg))
}
