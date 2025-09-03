package runtime_ext

import (
	"el/runtime"
	"fmt"
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/seq"
)

type Type = runtime.Type
type Object = runtime.Object

type Unwrap struct{}

func (u Unwrap) String() string {
	return "*"
}

func (u Unwrap) Type() Object {
	return runtime.DataType("unwrap")
}

type Int struct {
	int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.int)
}

func (i Int) Type() Object {
	return runtime.DataType("int")
}

type List struct {
	seq.Seq[Object]
}

func (l List) String() string {
	ls := make([]string, 0, l.Len())
	for _, o := range l.Iter {
		ls = append(ls, fmt.Sprint(o))
	}
	s := strings.Join(ls, " ")
	s = fmt.Sprintf("[%s]", s)
	return s
}

func (l List) Type() Object {
	return runtime.DataType("list")
}

type String struct {
	string
}

func (s String) String() string {
	return s.string
}

func (s String) Type() Object {
	return runtime.DataType("string")
}
