package ts

import "github.com/fbundle/lab_public/lab/go_util/pkg/adt"

func MustSingleName(level int, name string, parent Sort) Sort {
	var sort Sort
	if ok := SingleName(level, name, parent).Unwrap(&sort); !ok {
		panic("type_error")
	}
	return sort
}

func SingleName(level int, name string, parent Sort) adt.Option[Sort] {
	if parent != nil && level+1 != parent.Level() {
		// if parent is specified, then its level must be valid
		return adt.None[Sort]()
	}
	return adt.Some[Sort](singleName{
		level:  level,
		name:   name,
		parent: parent,
	})
}

// singleName - representing all single sorts
// this is a helper since with only singleData, one needs infinitely many objects to construct
// level 1: Int, Bool
// level 2: Type
type singleName struct {
	level  int
	name   string
	parent Sort
}

func (s singleName) Level() int {
	return s.level
}

func (s singleName) String() string {
	return s.name
}

func (s singleName) Parent() Sort {
	if s.parent != nil {
		return s.parent
	}
	return singleName{
		level: s.level + 1,
		name:  DefaultSortName,
	}
}

func (s singleName) Length() int {
	return 1
}

func (s singleName) LessEqual(dst Sort) bool {
	if s.Length() != dst.Length() || s.Level() != dst.Level() {
		return false
	}
	return le(s.String(), dst.String())
}

func (s singleName) prepend(param Sort) Sort {
	return chain{
		params: adt.MustNonEmpty([]Sort{param}),
		body:   s,
	}
}
