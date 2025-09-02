package runtime_core

import (
	"el/pkg/el/expr"
	"errors"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var ErrorNameNotFound = func(name Name) error {
	return fmt.Errorf("object not found %s", name)
}
var ErrorInterrupt = func(err error) error {
	return fmt.Errorf("interrupted: %s", err)
}
var ErrorTimeout = func(err error) error {
	return fmt.Errorf("timeout: %s", err)
}
var ErrorStackOverflow = errors.New("stack overflow")

var ErrorUnknownExpression = func(e expr.Expr) error {
	return fmt.Errorf("unknown expression type %s", e.String())
}
var ErrorCannotExecuteExpression = func(e expr.Expr) error {
	return fmt.Errorf("expression cannot be executed: %s", e.String())
}
var ErrorNotEnoughArguments = errors.New("not enough arguments")
var ErrorInternal = errors.New("internal")

func searchOnStack(s Stack, name Name) (Object, bool) {
	for _, frame := range s.Iter {
		if o, ok := frame.Get(name); ok {
			return o, true
		}
	}
	return nil, false
}

func object(o Object) adt.Option[Object] {
	return adt.Some[Object](o)
}

func errorObject(err error) adt.Option[Object] {
	return adt.Error[Object](err)
}

func errorObjectString(msg string) adt.Option[Object] {
	return errorObject(errors.New(msg))
}
