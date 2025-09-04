package ts

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

func SortData(level int, data Data, parent Sort) adt.Option[Sort] {
	if parent.Level() != level+1 {
		return adt.None[Sort]()
	}
	return adt.Some[Sort](singleData{
		level:  level,
		data:   data,
		parent: parent,
	})
}

type singleData struct {
	level  int
	data   Data
	parent Sort
}

func (s singleData) Data() adt.Option[Data] {
	return adt.Some(s.data)
}

func (s singleData) Level() int {
	return s.level
}

func (s singleData) String() string {
	return s.data.String()
}

func (s singleData) Type() Sort {
	return s.parent
}

func (s singleData) Cast(parent Sort) adt.Option[Sort] {
	if ok := s.Type().le(parent); !ok {
		return adt.None[Sort]()
	}
	return adt.Some[Sort](singleData{
		level:  s.level,
		data:   s.data,
		parent: parent,
	})
}

func (s singleData) Len() int {
	return 1
}

func (s singleData) le(dst Sort) bool {
	if s.Len() != dst.Len() || s.Level() != dst.Level() {
		return false
	}
	return le(s.String(), dst.String())
}

func (s singleData) prepend(param Sort) Sort {
	return chain{
		params: adt.MustNonEmpty([]Sort{param}),
		body:   s,
	}
}
