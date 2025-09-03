package runtime

import (
	"context"
	"el/pkg/el/expr"
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

var ErrorUnknownExpression = func(e expr.Expr) error {
	return fmt.Errorf("unknown expression type %s", e.String())
}
var ErrorCannotExecuteExpression = func(e expr.Expr) error {
	return fmt.Errorf("expression cannot be executed: %s", e.String())
}
var ErrorNotEnoughArguments = errors.New("not enough arguments")

type Runtime struct {
	MaxStackDepth      int
	ParseLiteralOpt    func(lit string) adt.Result[Value]
	PostProcessArgsOpt func(args []Value) adt.Result[[]Value]
}

func (r Runtime) Step(ctx context.Context, s Stack, e expr.Expr) adt.Result[Value] {
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
	case expr.Name:
		// parse literal
		var o Value
		if err := r.ParseLiteralOpt(string(e)).Unwrap(&o); err == nil {
			return value(o)
		}
		// search name on stack
		if ok := searchOnStack(s, Name(e)).Unwrap(&o); ok {
			return value(o)
		}
		return errValue(ErrorNameNotFound(Name(e)))
	case expr.Lambda:
		var cmd Value
		if err := r.Step(ctx, s, e.Cmd).Unwrap(&cmd); err != nil {
			return errValue(err)
		}
		var function Function
		if ok := adt.Cast[Function](cmd).Unwrap(&function); ok {
			return function.Apply(r, ctx, s, e.Args)
		}
		return errValue(ErrorCannotExecuteExpression(e))
	default:
		return errValue(ErrorUnknownExpression(e))
	}
}

// StepMany executes the expressions in parallel and returns the results in the same order as the input expressions
func (r Runtime) StepMany(ctx context.Context, s Stack, es ...expr.Expr) adt.Result[[]Value] {
	outputs := make([]Value, len(es))
	errHolder := &atomic.Value{}
	subCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := &sync.WaitGroup{}
	for i := 0; i < len(es); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			e := es[i]
			if subErr := r.Step(subCtx, s, e).Unwrap(&outputs[i]); subErr != nil {
				errHolder.CompareAndSwap(nil, subErr) // set the error only once
				cancel()                              // stop all other goroutines
			}
		}(i)
	}
	wg.Wait()
	if err := errHolder.Load(); err != nil {
		return adt.Err[[]Value](err.(error))
	}
	return adt.Ok[[]Value](outputs)
}
