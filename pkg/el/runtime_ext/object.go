package runtime_ext

import (
	"el/pkg/el/runtime_core"
	"fmt"
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/seq"
)

type Wildcard = runtime_core.Wildcard

type Unwrap struct{}

func (u Unwrap) String() string {
	return "*"
}

func (u Unwrap) MustTypeObject() {}

type Int int

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

func (i Int) MustTypeObject() {}

var True = Int(1)
var False = Int(0)

type List seq.Seq[Object]

func (l List) String() string {
	lst := (seq.Seq[Object])(l)
	ls := make([]string, 0, lst.Len())
	for _, o := range lst.Iter {
		ls = append(ls, o.String())
	}
	s := strings.Join(ls, ",")
	s = fmt.Sprintf("[%s]", s)
	return s

}

func (l List) MustTypeObject() {}
