package runtime

import (
	"context"
	"el/pkg/el/expr"
	"fmt"
	"strings"
)

// Object : union - TODO : introduce new data types
type Object interface {
	String() string
	MustValue() // for type-safety every Object must implement this
}

type Int int

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

func (i Int) MustValue() {}

var True = Int(1)
var False = Int(0)

type Unwrap struct{}

func (u Unwrap) String() string {
	return "*"
}

func (u Unwrap) MustValue() {}

type Wildcard struct{}

func (w Wildcard) String() string {
	return "_"
}

func (w Wildcard) MustValue() {}

type Lambda struct {
	ParamNameList  []Name    `json:"paramnamelist,omitempty"`
	Implementation expr.Expr `json:"implementation,omitempty"`
	Closure        Frame     `json:"closure,omitempty"`
}

func (l Lambda) String() string {
	s := fmt.Sprintf("(<closure_%p>; lambda ", l.Closure)
	for _, param := range l.ParamNameList {
		s += string(param) + " "
	}
	s += l.Implementation.String()
	s += ")"
	return s
}

func (l Lambda) MustValue() {}

type Module struct {
	Name Name `json:"name,omitempty"`
	Exec func(ctx context.Context, r *Runtime, e expr.Lambda) (Object, error)
	Man  string `json:"man,omitempty"`
}

func (m Module) String() string {
	return fmt.Sprintf("[%s]", m.Man)
}

func (m Module) MustValue() {}

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

func (l List) MustValue() {}
