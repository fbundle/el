package el

import (
	"context"
	"fmt"
	"strings"
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
	Params  []NameExpr `json:"params,omitempty"`
	Impl    Expr       `json:"impl,omitempty"`
	Closure Frame      `json:"closure,omitempty"`
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
	Name NameExpr `json:"name,omitempty"`
	Exec func(ctx context.Context, r *Runtime, expr LambdaExpr) (Object, error)
	Man  string `json:"man,omitempty"`
}

func (m Module) String() string {
	return fmt.Sprintf("[%s]", m.Man)
}

func (m Module) MustTypeObject() {}

type List []Object

func (l List) String() string {
	ls := make([]string, 0, len(l))
	for _, o := range l {
		ls = append(ls, o.String())
	}
	s := strings.Join(ls, ",")
	s = fmt.Sprintf("[%s]", s)
	return s

}

func (l List) MustTypeObject() {}
