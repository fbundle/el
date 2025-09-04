package ts

import "github.com/fbundle/lab_public/lab/go_util/pkg/adt"

func MustObject(level int, name string, parent Sort) Sort {
	var sort Sort
	if ok := Object(level, name, parent).Unwrap(&sort); !ok {
		panic("type_error")
	}
	return sort
}

func Object(level int, name string, parent Sort) adt.Option[Sort] {
	if parent != nil && level+1 != parent.Level() {
		// if parent is specified, then its level must be valid
		return adt.None[Sort]()
	}
	return adt.Some[Sort](object{
		level:  level,
		name:   name,
		parent: parent,
	})
}

// object - representing all primitive sorts
// level 1: Int, Bool
// level 2: Type
type object struct {
	level  int
	name   string
	parent Sort
}

func (s object) Level() int {
	return s.level
}

func (s object) String() string {
	return s.name
}

func (s object) Parent() Sort {
	if s.parent != nil {
		return s.parent
	}
	// default parent - must have this to avoid infinity
	return object{
		level: s.level + 1,
		name:  DefaultSortName,
	}
}

func (s object) Length() int {
	return 1
}

func (s object) LessEqual(dst Sort) bool {
	if s.Length() != dst.Length() || s.Level() != dst.Level() {
		return false
	}
	return le(s.String(), dst.String())
}

func (s object) prepend(param Sort) Sort {
	return morphism{
		params: adt.MustNonEmpty([]Sort{param}),
		body:   s,
	}
}
