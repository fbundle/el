package ts

import (
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

var _ Sort = chain{}

// chain - represent arrow type A -> B -> C
type chain struct {
	params adt.NonEmptySlice[Sort]
	body   Sort
}

func (c chain) Data() adt.Option[Data] {
	return adt.None[Data]()
}

func (c chain) Level() int {
	level := c.body.Level()
	for _, param := range c.params.Repr() {
		level = max(level, param.Level())
	}
	return level
}

func (c chain) String() string {
	strList := make([]string, 0, len(c.params.Repr())+1)
	for _, param := range c.params.Repr() {
		strList = append(strList, param.String())
	}
	strList = append(strList, c.body.String())
	return "{" + strings.Join(strList, " -> ") + "}"
}

func (c chain) Type() Sort {
	return singleName{
		level: c.Level() + 1,
		name:  TypeName,
	}
}

func (c chain) Cast(parent Sort) adt.Option[Sort] {
	// cannot cast sort of chain
	return adt.None[Sort]()
}

func (c chain) Len() int {
	return len(c.params.Repr()) + 1
}

func (c chain) le(dst Sort) bool {
	//TODO implement me
	panic("implement me")
}

func (c chain) prepend(param Sort) Sort {
	return chain{
		params: adt.MustNonEmpty[Sort](append([]Sort{param}, c.params.Repr()...)),
		body:   c.body,
	}
}
