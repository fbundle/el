package runtime

import (
	"context"
	"el/ast"
	"errors"
	"fmt"
	"time"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/ordered_map"
)

type Name string

type Frame = ordered_map.OrderedMap[Name, Object]

type Runtime struct {
	ParseLiteral func(lit string) adt.Result[Object]
	UnwrapArgs   func(argsOpt adt.Result[[]Object]) adt.Result[[]Object]
}

var ErrorNameNotFound = func(name Name) error {
	return fmt.Errorf("object not found %s", name)
}
var ErrorInterrupt = errors.New("interrupted")
var ErrorTimeout = errors.New("timeout")

var ErrorUnknownExpression = func(e ast.Expr) error {
	return fmt.Errorf("unknown expression type %s", e.String())
}
var ErrorCannotExecuteExpression = func(e ast.Expr) error {
	return fmt.Errorf("expression cannot be executed: %s", e.String())
}

func (r Runtime) Step(ctx context.Context, frame Frame, e ast.Expr) adt.Result[Object] {
	deadline, ok := ctx.Deadline()
	if ok && time.Now().After(deadline) {
		return resultErr(ErrorTimeout)
	}
	select {
	case <-ctx.Done():
		return resultErr(ErrorInterrupt)
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

	switch e := e.(type) {
	case ast.Name:
		name := Name(e)
		var o Object
		if ok := r.resolveName(frame, name).Unwrap(&o); !ok {
			return resultErr(ErrorNameNotFound(name))
		}
		return resultObj(o)
	case ast.Lambda:
		var cmd cmd
		if ok := getCmd(e).Unwrap(&cmd); !ok {
			return resultData(nil, NilType) // empty expression
		}
		var cmdObject Object
		if err := r.Step(ctx, frame, cmd.cmdExpr).Unwrap(&cmdObject); err != nil {
			return resultErr(err)
		}
		funcData, ok := cmdObject.Data().(FuncData)
		if !ok {
			return resultErr(ErrorCannotExecuteExpression(e))
		}
		return funcData.exec(r, ctx, frame, cmd.argExprList)

	default:
		return resultErr(ErrorUnknownExpression(e))
	}
}

// stepAndUnwrapArgs executes the argument expressions and unwraps the results
func (r Runtime) stepAndUnwrapArgs(ctx context.Context, frame Frame, argExprList []ast.Expr) adt.Result[[]Object] {
	args := make([]Object, len(argExprList))
	for i, e := range argExprList {
		if err := r.Step(ctx, frame, e).Unwrap(&args[i]); err != nil {
			return adt.Err[[]Object](err)
		}
	}
	return r.UnwrapArgs(adt.Ok(args))
}

func (r Runtime) resolveName(frame Frame, name Name) adt.Option[Object] {
	// search name on the stack
	o, ok := frame.Get(name)
	if ok {
		return adt.Some(o)
	}
	// parse literal
	if err := r.ParseLiteral(string(name)).Unwrap(&o); err == nil {
		return adt.Some(o)
	}
	return adt.None[Object]()
}

func resultObj(o Object) adt.Result[Object] {
	return adt.Ok(o)
}
func resultData(data Data, parent Object) adt.Result[Object] {
	return adt.Ok(makeData(data, parent))
}

func resultErr(err error) adt.Result[Object] {
	return adt.Err[Object](err)
}
func resultErrStrf(format string, args ...any) adt.Result[Object] {
	return adt.Err[Object](fmt.Errorf(format, args...))
}

type cmd struct {
	cmdExpr     ast.Expr
	argExprList []ast.Expr
}

func getCmd(e ast.Lambda) adt.Option[cmd] {
	if len(e) == 0 {
		return adt.None[cmd]()
	}
	return adt.Some(cmd{
		cmdExpr:     e[0],
		argExprList: e[1:],
	})
}
