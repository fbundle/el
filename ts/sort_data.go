package ts

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Data interface {
	String() string
}

func SortData(data Data, sort string) Sort {
	return sortData{
		data:   data,
		parent: MustSortGE1(1, sort),
	}
}

type sortData struct {
	data   Data
	parent Sort
}

func (s sortData) Level() int {
	return 0
}

func (s sortData) String() string {
	return s.data.String()
}

func (s sortData) Type() Sort {
	return s.parent
}

func (s sortData) Cast(sort Sort) adt.Option[Sort] {
	//TODO implement me
	panic("implement me")
}

func (s sortData) Chain() adt.NonEmptySlice[Sort] {
	return adt.MustNonEmpty([]Sort{s})
}

func (s sortData) le(dst Sort) bool {
	//TODO implement me
	panic("implement me")
}

func (s sortData) prepend(param Sort) Sort {
	//TODO implement me
	panic("implement me")
}
