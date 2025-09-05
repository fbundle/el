package runtime_ext

import (
	"el/runtime"
	"fmt"
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/seq"
)

type Object = runtime.Object
type Data = runtime.Data
type Unwrap struct{}

func (u Unwrap) String() string {
	return "*"
}

type Int struct {
	Val int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.Val)
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

type String struct {
	Val string
}

func (s String) String() string {
	return s.Val
}
