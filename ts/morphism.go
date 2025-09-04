package ts

import (
	"strings"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

func MustMorphism(sorts ...Sort) Sort {
	var sort Sort
	if ok := Morphism(sorts...).Unwrap(&sort); !ok {
		panic("type_error")
	}
	return sort
}

func Morphism(sorts ...Sort) adt.Option[Sort] {
	if len(sorts) == 0 {
		return adt.None[Sort]()
	}
	sort := sorts[len(sorts)-1]
	for i := len(sorts) - 2; i >= 0; i-- {
		sort = sort.prepend(sorts[i])
	}
	return adt.Some[Sort](sort)
}

// morphism - represent arrow type A -> B -> C
type morphism struct {
	params adt.NonEmptySlice[Sort]
	body   Sort
}

func (s morphism) Level() int {
	level := s.body.Level()
	for _, param := range s.params.Repr() {
		level = max(level, param.Level())
	}
	return level
}

func (s morphism) String() string {
	strList := make([]string, 0, len(s.params.Repr())+1)
	for _, param := range s.params.Repr() {
		strList = append(strList, param.String())
	}
	strList = append(strList, s.body.String())
	return "{" + strings.Join(strList, " -> ") + "}"
}

func (s morphism) Parent() Sort {
	return object{
		level: s.Level() + 1,
		name:  DefaultSortName,
	}
}

func (s morphism) Length() int {
	return len(s.params.Repr()) + 1
}

func (s morphism) LessEqual(dst Sort) bool {
	if s.Length() != dst.Length() || s.Level() != dst.Level() {
		return false
	}
	var d morphism
	if ok := adt.Cast[morphism](dst).Unwrap(&d); !ok {
		return false
	}
	length := len(s.params.Repr())
	for i := 0; i < length; i++ {
		sParam := s.params.Repr()[i]
		dParam := d.params.Repr()[i]
		if !dParam.LessEqual(sParam) {
			// reverse cast - similar to contravariant functor
			// {int} can be cast into {any}
			// {any -> int} can be cast into {int -> int}
			return false
		}
	}
	return s.body.LessEqual(d.body)
}

func (s morphism) prepend(param Sort) Sort {
	return morphism{
		params: adt.MustNonEmpty[Sort](append([]Sort{param}, s.params.Repr()...)),
		body:   s.body,
	}
}
