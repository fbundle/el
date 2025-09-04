package ts

import "github.com/fbundle/lab_public/lab/go_util/pkg/adt"

func MustSingleName(level int, name string) Sort {
	if level < 1 {
		panic("type_error")
	}
	return singleName{
		level: level,
		name:  name,
	}
}

func SingleName(level int, name string) adt.Option[Sort] {
	if level < 1 {
		return adt.None[Sort]()
	}
	return adt.Some[Sort](singleName{
		level: level,
		name:  name,
	})
}

// singleName - representing all single sorts with level >= 1
// this is a helper since with only singleData, one needs infinitely many objects to construct
// level 1: Int, Bool
// level 2: Type
type singleName struct {
	level int
	name  string
}

func (s singleName) Data() adt.Option[Data] {
	return adt.None[Data]()
}

func (s singleName) Level() int {
	return s.level
}

func (s singleName) String() string {
	return s.name
}

func (s singleName) Type() Sort {
	return singleName{
		level: s.level + 1,
		name:  TypeName,
	}
}

func (s singleName) Cast(parent Sort) adt.Option[Sort] {
	// cannot cast sort of name
	return adt.None[Sort]()
}

func (s singleName) Len() int {
	return 1
}

func (s singleName) le(dst Sort) bool {
	if s.Len() != s.Len() || s.Level() != dst.Level() {
		return false
	}
	return le(s.String(), dst.String())
}

func (s singleName) prepend(param Sort) Sort {
	return chain{
		parm: adt.MustNonEmpty([]Sort{param}),
		ret:  s,
	}
}
