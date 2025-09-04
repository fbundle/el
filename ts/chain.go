package ts

import "github.com/fbundle/lab_public/lab/go_util/pkg/adt"

var _ Sort = chain{}

type chain struct {
	par adt.NonEmptySlice[Sort]
	ret Sort
}

func (c chain) Level() int {
	//TODO implement me
	panic("implement me")
}

func (c chain) String() string {
	//TODO implement me
	panic("implement me")
}

func (c chain) Type() Sort {
	//TODO implement me
	panic("implement me")
}

func (c chain) Cast(sort Sort) adt.Option[Sort] {
	//TODO implement me
	panic("implement me")
}

func (c chain) Chain() adt.NonEmptySlice[Sort] {
	//TODO implement me
	panic("implement me")
}

func (c chain) le(dst Sort) bool {
	//TODO implement me
	panic("implement me")
}

func (c chain) prepend(param Sort) Sort {
	//TODO implement me
	panic("implement me")
}
