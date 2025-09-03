package runtime

import (
	"context"
	"el/ast"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
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

var ErrorUnknownExpression = func(e ast.Expr) error {
	return fmt.Errorf("unknown expression type %s", e.String())
}
var ErrorCannotExecuteExpression = func(e ast.Expr) error {
	return fmt.Errorf("expression cannot be executed: %s", e.String())
}
var ErrorNotEnoughArguments = errors.New("not enough arguments")

type Runtime struct {
	MaxStackDepth int
	ParseLiteral  func(lit string) adt.Result[Value]
	UnwrapArgs    func(argsOpt adt.Result[[]Value]) adt.Result[[]Value]
}

func (r Runtime) Step(ctx context.Context, s Stack, e ast.Expr) adt.Result[Value] {
	deadline, ok := ctx.Deadline()
	if ok && time.Now().After(deadline) {
		return errValue(ErrorTimeout(ctx.Err()))
	}
	select {
	case <-ctx.Done():
		return errValue(ErrorInterrupt(ctx.Err()))
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
	case ast.Leaf:
		// parse literal
		var o Value
		if err := r.ParseLiteral(string(e)).Unwrap(&o); err == nil {
			return value(o)
		}
		// search name on the stack
		if ok := searchOnStack(s, Name(e)).Unwrap(&o); ok {
			return value(o)
		}
		return errValue(ErrorNameNotFound(Name(e)))
	case ast.Node:
		var cmd Value
		if err := r.Step(ctx, s, e.Cmd).Unwrap(&cmd); err != nil {
			return errValue(err)
		}
		var function Function
		if ok := adt.Cast[Function](cmd).Unwrap(&function); ok {
			return function.apply(r, ctx, s, e.Args)
		}
		return errValue(ErrorCannotExecuteExpression(e))
	default:
		return errValue(ErrorUnknownExpression(e))
	}
}

// StepAndUnwrapArgs executes the argument expressions in parallel and unwraps the results
func (r Runtime) StepAndUnwrapArgs(ctx context.Context, s Stack, argList []ast.Expr) adt.Result[[]Value] {
	args := make([]Value, len(argList))
	errHolder := &atomic.Value{}
	subCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := &sync.WaitGroup{}
	for i := 0; i < len(argList); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			e := argList[i]
			if subErr := r.Step(subCtx, s, e).Unwrap(&args[i]); subErr != nil {
				errHolder.CompareAndSwap(nil, subErr) // set the error only once
				cancel()                              // stop all other goroutines
			}
		}(i)
	}
	wg.Wait()

	var argsOpt = adt.Result[[]Value]{
		Val: args,
	}
	if err := errHolder.Load(); err != nil {
		argsOpt.Err = err.(error)
	}
	return r.UnwrapArgs(argsOpt)
}
