package ts

import "github.com/fbundle/lab_public/lab/go_util/pkg/adt"

func MustSingle(level int, name string, parent Sort) Sort {
	var sort Sort
	if ok := Single(level, name, parent).Unwrap(&sort); !ok {
		panic("type_error")
	}
	return sort
}

func Single(level int, name string, parent Sort) adt.Option[Sort] {
	if parent != nil && level+1 != parent.Level() {
		// if parent is specified, then its level must be valid
		return adt.None[Sort]()
	}
	return adt.Some[Sort](single{
		level:  level,
		name:   name,
		parent: parent,
	})
}

// single - representing all single sorts
// level 1: Int, Bool
// level 2: Type
type single struct {
	level  int
	name   string
	parent Sort
}

func (s single) Level() int {
	return s.level
}

func (s single) String() string {
	return s.name
}

func (s single) Parent() Sort {
	if s.parent != nil {
		return s.parent
	}
	// default parent - must have this to avoid infinity
	return single{
		level: s.level + 1,
		name:  DefaultSortName,
	}
}

func (s single) Length() int {
	return 1
}

func (s single) LessEqual(dst Sort) bool {
	if s.Length() != dst.Length() || s.Level() != dst.Level() {
		return false
	}
	return le(s.String(), dst.String())
}

func (s single) prepend(param Sort) Sort {
	return chain{
		params: adt.MustNonEmpty([]Sort{param}),
		body:   s,
	}
}
