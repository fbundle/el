package ts

import "github.com/fbundle/lab_public/lab/go_util/pkg/adt"

func MustSingleName(level int, name string) Sort {
	return singleName{
		level: level,
		name:  name,
	}
}

// singleName - representing all single sorts
// this is a helper since with only singleData, one needs infinitely many objects to construct
// level 1: Int, Bool
// level 2: Type
type singleName struct {
	level int
	name  string
}

func (s singleName) Level() int {
	return s.level
}

func (s singleName) String() string {
	return s.name
}

func (s singleName) Parent() Sort {
	return singleName{
		level: s.level + 1,
		name:  DefaultSortName,
	}
}

func (s singleName) Length() int {
	return 1
}

func (s singleName) LE(dst Sort) bool {
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
