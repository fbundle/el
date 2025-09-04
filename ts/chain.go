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

func (s chain) Data() adt.Option[Data] {
	return adt.None[Data]()
}

func (s chain) Level() int {
	level := s.body.Level()
	for _, param := range s.params.Repr() {
		level = max(level, param.Level())
	}
	return level
}

func (s chain) String() string {
	strList := make([]string, 0, len(s.params.Repr())+1)
	for _, param := range s.params.Repr() {
		strList = append(strList, param.String())
	}
	strList = append(strList, s.body.String())
	return "{" + strings.Join(strList, " -> ") + "}"
}

func (s chain) Type() Sort {
	return singleName{
		level: s.Level() + 1,
		name:  TypeName,
	}
}

func (s chain) Cast(parent Sort) adt.Option[Sort] {
	// cannot cast sort of chain
	return adt.None[Sort]()
}

func (s chain) Len() int {
	return len(s.params.Repr()) + 1
}

func (s chain) le(dst Sort) bool {
	if s.Len() != dst.Len() || s.Level() != dst.Level() {
		return false
	}
	var d chain
	if ok := adt.Cast[chain](dst).Unwrap(&d); !ok {
		return false
	}
	length := len(s.params.Repr())
	for i := 0; i < length; i++ {
		sParam := s.params.Repr()[i]
		dParam := d.params.Repr()[i]
		if !dParam.le(sParam) {
			// reverse cast - similar to contravariant functor
			// {int} can be cast into {any}
			// {any -> int} can be cast into {int -> int}
			return false
		}
	}
	return s.body.le(d.body)
}

func (s chain) prepend(param Sort) Sort {
	return chain{
		params: adt.MustNonEmpty[Sort](append([]Sort{param}, s.params.Repr()...)),
		body:   s.body,
	}
}
