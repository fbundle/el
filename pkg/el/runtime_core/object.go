package runtime_core

import (
	"context"
	"el/pkg/el/expr"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Value interface {
	String() string
	MustValue()
}

type Command interface {
	Value
	Apply(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Option[Value]
}

type Lambda struct {
	Params  []Name    `json:"params,omitempty"`
	Body    expr.Expr `json:"body,omitempty"`
	Closure Frame     `json:"closure,omitempty"`
}

func (l Lambda) String() string {
	s := fmt.Sprintf("(<closure_%p>; lambda ", l.Closure)
	for _, param := range l.Params {
		s += string(param) + " "
	}
	s += l.Body.String()
	s += ")"
	return s
}

func (l Lambda) MustValue() {}

func (l Lambda) Apply(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Option[Value] {
	// 0. sanity check
	if len(argList) < len(l.Params) {
		errorValue(ErrorNotEnoughArguments)
	}
	// 1. evaluate arguments
	args := make([]Value, len(argList))
	for i, argExpr := range argList {
		if err := r.StepOpt(ctx, s, argExpr).Unwrap(&args[i]); err != nil {
			return errorValue(err)
		}
	}
	if err := r.PostProcessArgsOpt(args).Unwrap(&args); err != nil {
		return errorValue(err)
	}
	// 2. make call stack
	local := l.Closure
	for i := 0; i < len(l.Params); i++ {
		param, arg := l.Params[i], args[i]
		local = local.Set(param, arg)
	}
	callStack := s.Push(local)
	// 3. make call with new stack
	var o Value
	if err := r.StepOpt(ctx, callStack, l.Body).Unwrap(&o); err != nil {
		return errorValue(err)
	}
	return value(o)
}

type Module struct {
	Name Name
	Exec func(r Runtime, ctx context.Context, s Stack, args []expr.Expr) adt.Option[Value]
	Man  string
}

func (m Module) String() string {
	return fmt.Sprintf("[%s]", m.Man)
}
func (m Module) MustValue() {}

func (m Module) Apply(r Runtime, ctx context.Context, s Stack, args []expr.Expr) adt.Option[Value] {
	return m.Exec(r, ctx, s, args)
}
