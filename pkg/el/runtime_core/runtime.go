package runtime

import (
	"context"
	"el/pkg/el/expr"
	"time"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

func StepOpt(ctx context.Context, s Stack, e expr.Expr) adt.Option[Object] {
	deadline, ok := ctx.Deadline()
	if ok && time.Now().After(deadline) {
		return errorObject(TimeoutError(ctx.Err()))
	}
	select {
	case <-ctx.Done():
		return errorObject(InterruptError(ctx.Err()))
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

	if s.Depth() > MAX_STACK_DEPTH {
		return errorObject(StackOverflowError)
	}
	switch e := e.(type) {
	case expr.Name:
		// search name on stack
		if o, ok := searchOnStack(s, Name(e)); ok {
			return object(o)
		}
		// parse literal
		o, err := parseLiteral(string(e))
		if err != nil {
			return errorObject(err)
		}
		return object(o)
	case expr.Lambda:
		var cmd Object
		if err := StepOpt(ctx, s, e.Cmd).Unwrap(&cmd); err != nil {
			return errorObject(err)
		}
		switch cmd := cmd.(type) {
		case Module:
			return cmd.Exec(ctx, s, e)
		case Lambda:
			// 0. sanity check
			if len(e.Args) < len(cmd.ParamNameList) {
				errorObject(NotEnoughArguments)
			}
			// 1. evaluate arguments
			var args []Object
			if err := stepManyOpt(ctx, s, e.Args...).Unwrap(&args); err != nil {
				return errorObject(err)
			}
			if err := unwrapArgsOpt(args).Unwrap(&args); err != nil {
				return errorObject(err)
			}
			// 2. make call stack
			local := emptyFrame
			for i := 0; i < len(cmd.ParamNameList); i++ {
				param, arg := cmd.ParamNameList[i], args[i]
				local = local.Set(param, arg)
			}
			callStack := cmd.Closure.Push(local)
			// 3. make call with new stack
			var o Object
			if err := StepOpt(ctx, callStack, cmd.Implementation).Unwrap(&o); err != nil {
				return errorObject(err)
			}
			return object(o)
		default:
			return errorObject(CannotExecuteExpression(e))
		}

	default:
		return errorObject(UnknownExpression(e))
	}
}

func stepManyOpt(ctx context.Context, s Stack, exprList ...expr.Expr) adt.Option[[]Object] {
	var outputs = make([]Object, len(exprList))
	for i, e := range exprList {
		if err := StepOpt(ctx, s, e).Unwrap(&outputs[i]); err != nil {
			return adt.Error[[]Object](err)
		}
	}
	return adt.Some(outputs)
}
