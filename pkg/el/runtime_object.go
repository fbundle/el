package el

import (
	"context"
	"fmt"
)

// Object : union - TODO : introduce new data types
type Object interface {
	String() string
	MustTypeObject() // for type-safety every Object must implement this
}

type Int int

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

func (i Int) MustTypeObject() {}

var True = Int(1)
var False = Int(0)

type Unwrap struct{}

func (u Unwrap) String() string {
	return "*"
}

func (u Unwrap) MustTypeObject() {}

type Wildcard struct{}

func (w Wildcard) String() string {
	return "_"
}

func (w Wildcard) MustTypeObject() {}

type Lambda struct {
	Params []string `json:"params,omitempty"`
	Impl   Expr     `json:"impl,omitempty"`
	Frame  Frame    `json:"frame,omitempty"`
}

func (l Lambda) String() string {
	s := fmt.Sprintf("(<frame_%p>; lambda ", l.Frame)
	for _, param := range l.Params {
		s += param + " "
	}
	s += l.Impl.String()
	s += ")"
	return s
}

func (l Lambda) MustTypeObject() {}

type Module struct {
	Name string `json:"name,omitempty"`
	Exec func(ctx context.Context, r *Runtime, expr LambdaExpr) (Object, error)
	Man  string `json:"man,omitempty"`
}

func (m Module) String() string {
	return fmt.Sprintf("[%s]", m.Man)
}

func (m Module) MustTypeObject() {}

type List []Object

func (l List) String() string {
	s := ""
	s += "["
	for _, obj := range l {
		s += fmt.Sprintf("%v,", obj)
	}
	s += "]"
	return s
}

func (l List) MustTypeObject() {}
