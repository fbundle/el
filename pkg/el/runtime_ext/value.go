package runtime_ext

import (
	"el/pkg/el/runtime"
	"errors"
	"fmt"
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/seq"
)

type Value = runtime.Value

type Unwrap struct{}

func (u Unwrap) String() string {
	return "*"
}

type Int struct {
	int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.int)
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

type String struct {
	string
}

func (s String) String() string {
	return fmt.Sprintf("\"%s\"", s.string)
}

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
