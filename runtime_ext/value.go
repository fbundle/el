package runtime_ext

import (
	"el/runtime"
	"fmt"
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/seq"
)

type Type = runtime.Type
type Value = runtime.Value

type Unwrap struct{}

func (u Unwrap) String() string {
	return "*"
}

func (u Unwrap) Type() Type {
	return runtime.DataType("unwrap")
}

type Int struct {
	int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.int)
}

func (i Int) Type() Type {
	return runtime.DataType("int")
}

type List struct {
	seq.Seq[Value]
}

func (l List) String() string {
	ls := make([]string, 0, l.Len())
	for _, o := range l.Iter {
		ls = append(ls, o.String())
	}
	s := strings.Join(ls, ",")
	s = fmt.Sprintf("[%s]", s)
	return s
}

func (l List) Type() Type {
	return runtime.DataType("list")
}

type String struct {
	string
}

func (s String) String() string {
	return fmt.Sprintf("\"%s\"", s.string)
}

func (s String) Type() Type {
	return runtime.DataType("string")
}
