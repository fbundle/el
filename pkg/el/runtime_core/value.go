package runtime_core

import (
	"context"
	"el/pkg/el/expr"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Value interface {
	String() string
	MustValue() // for type-safety every Value must implement this
}

type Function interface {
	Value
	Apply(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Result[Value]
}

// Lambda - a function written in S-expression
type Lambda struct {
	ParamList []Name    `json:"param_list,omitempty"`
	Body      expr.Expr `json:"body,omitempty"`
	Closure   Frame     `json:"closure,omitempty"`
}

func (l Lambda) String() string {
	s := fmt.Sprintf("(<closure_%p>; lambda ", l.Closure)
	for _, param := range l.ParamList {
		s += string(param) + " "
	}
	s += l.Body.String()
	s += ")"
	return s
}

func (l Lambda) MustValue() {}

func (l Lambda) Apply(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Result[Value] {
	// 0. sanity check
	if len(argList) < len(l.ParamList) {
		errValue(ErrorNotEnoughArguments)
	}
	// 1. evaluate arguments
	args := make([]Value, len(argList))
	for i, argExpr := range argList {
		if err := r.Step(ctx, s, argExpr).Unwrap(&args[i]); err != nil {
			return errValue(err)
		}
	}
	if err := r.PostProcessArgsOpt(args).Unwrap(&args); err != nil {
		return errValue(err)
	}
	// 2. make call stack
	local := l.Closure
	for i := 0; i < len(l.ParamList); i++ {
		param, arg := l.ParamList[i], args[i]
		local = local.Set(param, arg)
	}
	callStack := s.Push(local)
	// 3. make call with new stack
	var o Value
	if err := r.Step(ctx, callStack, l.Body).Unwrap(&o); err != nil {
		return errValue(err)
	}
	return value(o)
}

type Exec = func(r Runtime, ctx context.Context, s Stack, args []expr.Expr) adt.Result[Value]

// Module - a function built-in to the language
type Module struct {
	Name Name
	Exec Exec
	Man  string
}

func (m Module) String() string {
	return fmt.Sprintf("[%s]", m.Man)
}
func (m Module) MustValue() {}

func (m Module) Apply(r Runtime, ctx context.Context, s Stack, args []expr.Expr) adt.Result[Value] {
	return m.Exec(r, ctx, s, args)
}

// helpers
func value(o Value) adt.Result[Value] {
	return adt.Ok[Value](o)
}

func errValue(err error) adt.Result[Value] {
	return adt.Err[Value](err)
}
