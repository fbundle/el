package runtime_core

import (
	"context"
	"el/pkg/el/expr"
	"time"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

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
		switch cmd := cmd.(type) {
		case Module:
			return cmd.Exec(r, ctx, s, e)
		case Lambda:
			// 0. sanity check
			if len(e.Args) < len(cmd.Params) {
				errorObject(ErrorNotEnoughArguments)
			}
			// 1. evaluate arguments
			args := make([]Object, len(e.Args))
			for i, argExpr := range e.Args {
				if err := r.StepOpt(ctx, s, argExpr).Unwrap(&args[i]); err != nil {
					return errorObject(err)
				}
			}
			if err := r.UnwrapArgsOpt(args).Unwrap(&args); err != nil {
				return errorObject(err)
			}
			// 2. make call stack
			local := cmd.Closure
			for i := 0; i < len(cmd.Params); i++ {
				param, arg := cmd.Params[i], args[i]
				local = local.Set(param, arg)
			}
			callStack := s.Push(local)
			// 3. make call with new stack
			var o Object
			if err := r.StepOpt(ctx, callStack, cmd.Impl).Unwrap(&o); err != nil {
				return errorObject(err)
			}
			return object(o)
		default:
			return errorObject(ErrorCannotExecuteExpression(e))
		}

	default:
		return errorObject(ErrorUnknownExpression(e))
	}
}
