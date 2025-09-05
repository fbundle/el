package runtime_ext

import (
	"el/runtime"
	"fmt"
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/seq"
)

type Object = runtime.Object
type Data = runtime.Data

// TypedData -
// the reason why we don't include TypedData in core runtime since functions are Data but don't have a type name,
// hence we handle it by _object{data, sort, parent} where parent (data type) is created as a sort
type TypedData interface {
	Data
	TypeName() string
}

type Unwrap struct{}

func (u Unwrap) String() string {
	return "*"
}
func (u Unwrap) TypeName() string {
	return "unwrap"
}

type Int struct {
	Val int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.Val)
}
func (i Int) TypeName() string {
	return "int"
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
func (l List) TypeName() string {
	return "list"
}

type String struct {
	Val string
}

func (s String) String() string {
	return s.Val
}

func (s String) TypeName() string {
	return "string"
}
