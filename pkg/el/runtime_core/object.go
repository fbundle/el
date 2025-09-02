package runtime_core

import (
	"context"
	"el/pkg/el/expr"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Object interface {
	String() string
	MustTypeObject()
}

type Command interface {
	Object
	Exec(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Option[Object]
}

type Lambda struct {
	Params  []Name    `json:"params,omitempty"`
	Impl    expr.Expr `json:"impl,omitempty"`
	Closure Frame     `json:"closure,omitempty"`
}

func (l Lambda) String() string {
	s := fmt.Sprintf("(<closure_%p>; lambda ", l.Closure)
	for _, param := range l.Params {
		s += string(param) + " "
	}
	s += l.Impl.String()
	s += ")"
	return s
}

func (l Lambda) MustTypeObject() {}

func (l Lambda) Exec(r Runtime, ctx context.Context, s Stack, argList []expr.Expr) adt.Option[Object] {
	// 0. sanity check
	if len(argList) < len(l.Params) {
		errorObject(ErrorNotEnoughArguments)
	}
	// 1. evaluate arguments
	args := make([]Object, len(argList))
	for i, argExpr := range argList {
		if err := r.StepOpt(ctx, s, argExpr).Unwrap(&args[i]); err != nil {
			return errorObject(err)
		}
	}
	if err := r.UnwrapArgsOpt(args).Unwrap(&args); err != nil {
		return errorObject(err)
	}
	// 2. make call stack
	local := l.Closure
	for i := 0; i < len(l.Params); i++ {
		param, arg := l.Params[i], args[i]
		local = local.Set(param, arg)
	}
	callStack := s.Push(local)
	// 3. make call with new stack
	var o Object
	if err := r.StepOpt(ctx, callStack, l.Impl).Unwrap(&o); err != nil {
		return errorObject(err)
	}
	return object(o)
}
