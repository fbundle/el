package runtime

import (
	"context"
	"el/ast"
	"errors"
	"fmt"
	"time"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var ErrorNameNotFound = func(name Name) error {
	return fmt.Errorf("object not found %s", name)
}
var ErrorInterrupt = errors.New("interrupted")
var ErrorTimeout = errors.New("timeout")
var ErrorStackOverflow = errors.New("stack overflow")

var ErrorUnknownExpression = func(e ast.Expr) error {
	return fmt.Errorf("unknown expression type %s", e.String())
}
var ErrorCannotExecuteExpression = func(e ast.Expr) error {
	return fmt.Errorf("expression cannot be executed: %s", e.String())
}
var ErrorNotEnoughArguments = errors.New("not enough arguments")

type Runtime struct {
	MaxStackDepth int
	ParseLiteral  func(lit string) adt.Result[Object]
	UnwrapArgs    func(argsOpt adt.Result[[]Object]) adt.Result[[]Object]
}

func (r Runtime) Step(ctx context.Context, s Stack, e ast.Expr) adt.Result[Object] {
	deadline, ok := ctx.Deadline()
	if ok && time.Now().After(deadline) {
		return errValue(ErrorTimeout)
	}
	select {
	case <-ctx.Done():
		return errValue(ErrorInterrupt)
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
		return errValue(ErrorStackOverflow)
	}
	switch e := e.(type) {
	case ast.Name:
		name := Name(e)
		var o Object
		if ok := r.resolveName(s, name).Unwrap(&o); !ok {
			return errValue(ErrorNameNotFound(name))
		}
		return value(o)
	case ast.Lambda:
		var cmd cmd
		if ok := getCmd(e).Unwrap(&cmd); !ok {
			return errValue(nil) // empty expression
		}
		var cmdObject Object
		if err := r.Step(ctx, s, cmd.cmdExpr).Unwrap(&cmdObject); err != nil {
			return errValue(err)
		}
		var funcObject Function
		if ok := adt.Cast[Function](cmdObject).Unwrap(&funcObject); !ok {
			return errValue(ErrorCannotExecuteExpression(e))
		}
		return funcObject.exec(r, ctx, s, cmd.argList)

	default:
		return errValue(ErrorUnknownExpression(e))
	}
}

// stepAndUnwrapArgs executes the argument expressions in parallel and unwraps the results
func (r Runtime) stepAndUnwrapArgs(ctx context.Context, s Stack, argList []ast.Expr) adt.Result[[]Object] {
	args := make([]Object, len(argList))
	for i, e := range argList {
		if err := r.Step(ctx, s, e).Unwrap(&args[i]); err != nil {
			return adt.Err[[]Object](err)
		}
	}
	return r.UnwrapArgs(adt.Ok(args))
}

func (r Runtime) resolveName(s Stack, name Name) adt.Option[Object] {
	var o Object
	// search name on the stack
	if ok := findStack(s, name).Unwrap(&o); ok {
		return adt.Some(o)
	}
	// parse literal
	if err := r.ParseLiteral(string(name)).Unwrap(&o); err == nil {
		return adt.Some(o)
	}
	return adt.None[Object]()
}
