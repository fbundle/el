package ts

import "github.com/fbundle/lab_public/lab/go_util/pkg/adt"

func MustSortGE1(level int, name string) Sort {
	if level < 1 {
		panic("type_error")
	}
	return sortGE1{
		level: level,
		name:  name,
	}
}

func SortGE1(level int, name string) adt.Option[Sort] {
	if level < 1 {
		return adt.None[Sort]()
	}
	return adt.Some[Sort](sortGE1{
		level: level,
		name:  name,
	})
}

// sortGE1 - representing all single sort with level >= 1
// level 1: Int, Bool
// level 2: Type
type sortGE1 struct {
	level int
	name  string
}

func (s sortGE1) Level() int {
	return s.level
}

func (s sortGE1) String() string {
	return s.name
}

func (s sortGE1) Type() Sort {
	return sortGE1{
		level: s.level + 1,
		name:  TypeName,
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
	//TODO implement me
	panic("implement me")
}

func (s sortGE1) prepend(param Sort) Sort {
	//TODO implement me
	panic("implement me")
}
