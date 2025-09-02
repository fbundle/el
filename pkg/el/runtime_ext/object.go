package runtime_ext

import (
	"context"
	"el/pkg/el/expr"
	"el/pkg/el/runtime_core"
	"fmt"
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/seq"
)

type Module struct {
	name runtime_core.Name
	exec func(r runtime_core.Runtime, ctx context.Context, s runtime_core.Stack, args []expr.Expr) adt.Option[runtime_core.Object]
	man  string
}

func (m Module) String() string {
	return fmt.Sprintf("[%s]", m.man)
}
func (m Module) MustTypeObject() {}

func (m Module) Exec(r runtime_core.Runtime, ctx context.Context, s runtime_core.Stack, args []expr.Expr) adt.Option[runtime_core.Object] {
	return m.exec(r, ctx, s, args)
}

type Unwrap struct{}

func (u Unwrap) String() string {
	return "*"
}

func (u Unwrap) MustTypeObject() {}

type Int struct {
	int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.int)
}

func (i Int) MustTypeObject() {}

type List struct {
	seq.Seq[Object]
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

func (l List) MustTypeObject() {}
