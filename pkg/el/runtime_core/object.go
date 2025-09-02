package runtime_core

import (
	"context"
	"el/pkg/el/expr"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

// Object : union
type Object interface {
	String() string
	MustTypeObject() // for type-safety every Object must implement this
}

type Wildcard struct{}

func (w Wildcard) String() string {
	return "_"
}

func (w Wildcard) MustTypeObject() {}

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

type Module struct {
	Name Name `json:"name,omitempty"`
	Exec func(r Runtime, ctx context.Context, s Stack, e expr.Lambda) adt.Option[Object]
	Man  string `json:"man,omitempty"`
}

func (m Module) String() string {
	return fmt.Sprintf("[%s]", m.Man)
}

func (m Module) MustTypeObject() {}
