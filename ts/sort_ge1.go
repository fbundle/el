package ts

import (
	"strconv"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

func MustSortGE1(level int, name string) Sort {
	var sort Sort
	if ok := SortGE(level, name).Unwrap(&sort); !ok {
		panic("type_error")
	}
	return sort
}

func SortGE(level int, name string) adt.Option[Sort] {
	if level < 1 {
		return adt.None[Sort]()
	}
	return adt.Some[Sort](sortGE1{
		level: level,
		name:  name,
	})
}

// sortGE1 - representing a single sort with level >= 1
// level 1: Int, Bool, etc.
// level 2: TypeName, etc.
type sortGE1 struct {
	level int
	name  string
}

func (s sortGE1) Level() int {
	return s.level
}

func (s sortGE1) String() string {
	return s.name + "_" + strconv.Itoa(s.level)
}

func (s sortGE1) Type() Sort {
	return sortGE1{
		level: s.level + 1,
		name:  TypeName, // everything from level 2 is just TypeName
	}
}

func (s sortGE1) Cast(sort Sort) adt.Option[Sort] {
	//TODO implement me
	panic("implement me")
}

func (s sortGE1) Chain() adt.NonEmptySlice[Sort] {
	return adt.MustNonEmpty([]Sort{s})
}

func (s sortGE1) le(dst Sort) bool {
	var d sortGE1
	if ok := adt.Cast[sortGE1](dst).Unwrap(&d); !ok {
		return false
	}
	if s.level != d.level {
		return false
	}
	return leName(s.name, d.name)
}

func (s sortGE1) prepend(param Sort) Sort {
	return chain{
		par: adt.MustNonEmpty([]Sort{param}),
		ret: s,
	}
}

var _ Sort = sortGE1{}
