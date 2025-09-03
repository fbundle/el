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

func (u Unwrap) MustString() string {
	return "*"
}

type Int struct {
	int
}

func (i Int) MustString() string {
	return fmt.Sprintf("%d", i.int)
}

type List struct {
	seq.Seq[Value]
}

func (l List) MustString() string {
	ls := make([]string, 0, l.Len())
	for _, o := range l.Iter {
		ls = append(ls, o.MustString())
	}
	s := strings.Join(ls, ",")
	s = fmt.Sprintf("[%s]", s)
	return s

}

type Bool struct {
	bool
}

func (b Bool) MustString() string {
	if b.bool {
		return "true"
	} else {
		return "false"
	}
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
