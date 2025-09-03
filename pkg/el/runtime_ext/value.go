package runtime_ext

import (
	"el/pkg/el/runtime_core"
	"errors"
	"fmt"
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/seq"
)

type Value = runtime_core.Value

type Unwrap struct{}

func (u Unwrap) String() string {
	return "*"
}

func (u Unwrap) MustValue() {}

type Int struct {
	int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.int)
}

func (i Int) MustValue() {}

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

func (l List) MustValue() {}

// helpers
func value(o Value) adt.Result[Value] {
	return adt.Ok[Value](o)
}

func errValue(err error) adt.Result[Value] {
	return adt.Err[Value](err)
}

func errValueString(msg string) adt.Result[Value] {
	return errValue(errors.New(msg))
}
