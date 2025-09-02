package runtime_core

import (
	"context"
	"el/pkg/el/expr"
	"errors"
	"fmt"
	"time"

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

type Runtime struct {
	MaxStackDepth   int
	ParseLiteralOpt func(lit string) adt.Option[Object]
	UnwrapArgsOpt   func(args []Object) adt.Option[[]Object]
}

func (r Runtime) StepOpt(ctx context.Context, s Stack, e expr.Expr) adt.Option[Object] {
	deadline, ok := ctx.Deadline()
	if ok && time.Now().After(deadline) {
		return errorObject(ErrorTimeout(ctx.Err()))
	}
	select {
	case <-ctx.Done():
		return errorObject(ErrorInterrupt(ctx.Err()))
	default:
	}

	/*
		the whole language is every simple
			1. parse literal or search on stack
			2. function application: push a new frame, exec the function, pop
			3. builtin module
				a. lambda: capture the current frame and save the implementation
				b. let: push a new frame, exec the function, pop
				c. match: eval and match

		only let and function application push a new frame since
			- let requires local scope to bind new variables
			- function application requires local scope to
				bind parameters and previously captured variables in lambda
	*/

	if s.Depth() > r.MaxStackDepth {
		return errorObject(ErrorStackOverflow)
	}
	switch e := e.(type) {
	case expr.Name:
		// parse literal
		var o Object
		if err := r.ParseLiteralOpt(string(e)).Unwrap(&o); err == nil {
			return object(o)
		}
		// search name on stack
		if o, ok := searchOnStack(s, Name(e)); ok {
			return object(o)
		}
		return errorObject(ErrorNameNotFound(Name(e)))
	case expr.Lambda:
		var cmd Object
		if err := r.StepOpt(ctx, s, e.Cmd).Unwrap(&cmd); err != nil {
			return errorObject(err)
		}
		if cmd, ok := cmd.(Command); ok {
			return cmd.Exec(r, ctx, s, e.Args)
		}
		return errorObject(ErrorCannotExecuteExpression(e))
	default:
		return errorObject(ErrorUnknownExpression(e))
	}
}

func object(o Object) adt.Option[Object] {
	return adt.Some[Object](o)
}

func errorObject(err error) adt.Option[Object] {
	return adt.Error[Object](err)
}
