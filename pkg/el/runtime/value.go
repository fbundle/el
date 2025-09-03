package runtime

import (
	"context"
	"el/pkg/el/expr"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Value interface {
	Type() Type
	String() string
}
type Function interface {
	Value
	Apply(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Result[Value]
}

func DataType(name string) Type {
	return Type{
		level: 0,
		name:  name,
	}
}

type Type struct {
	level int
	name  string
}

func (t Type) String() string {
	if t.level == 0 {
		return t.name
	}
	return fmt.Sprintf("type_%d", t.level)
}

func (t Type) Type() Type {
	return Type{
		level: t.level + 1,
	}
}

// Lambda - a function written in S-expression
type Lambda struct {
	ParamList []Name    `json:"param_list,omitempty"`
	Body      expr.Expr `json:"body,omitempty"`
	Closure   Frame     `json:"closure,omitempty"`
}

func (l Lambda) Type() Type {
	return DataType("lambda")
}

func (l Lambda) String() string {
	s := fmt.Sprintf("(<closure_%s>; lambda ", l.Closure.Repr())
	for _, param := range l.ParamList {
		s += string(param) + " "
	}
	s += l.Body.String()
	s += ")"
	return s
}

func (l Lambda) Apply(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Result[Value] {
	// 0. sanity check
	if len(argList) < len(l.ParamList) {
		errValue(ErrorNotEnoughArguments)
	}
	// 1. evaluate arguments
	var args []Value
	if err := r.StepAndUnwrapArgs(ctx, s, argList).Unwrap(&args); err != nil {
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
	Man  string
	Exec Exec
}

func (m Module) Type() Type {
	return DataType("module")
}

func (m Module) String() string {
	return fmt.Sprintf("[%s]", m.Man)
}

func (m Module) Apply(r Runtime, ctx context.Context, s Stack, args []expr.Expr) adt.Result[Value] {
	return m.Exec(r, ctx, s, args)
}
