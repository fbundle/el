package ts

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

// sortGE1 - representing a single sort with level >= 1
// level 1: Int, Bool, etc.
// level 2: Type, etc.
type sortGE1 struct {
	level int
	name  string
}

func (s sortGE1) Level() int {
	//TODO implement me
	panic("implement me")
}

func (s sortGE1) String() string {
	//TODO implement me
	panic("implement me")
}

func (s sortGE1) Type() Sort {
	//TODO implement me
	panic("implement me")
}

func (s sortGE1) Cast(sort Sort) adt.Option[Sort] {
	//TODO implement me
	panic("implement me")
}

func (s sortGE1) Chain() adt.NonEmptySlice[Sort] {
	//TODO implement me
	panic("implement me")
}

func (s sortGE1) le(dst Sort) bool {
	//TODO implement me
	panic("implement me")
}

func (s sortGE1) prepend(param Sort) Sort {
	//TODO implement me
	panic("implement me")
}

var _ Sort = sortGE1{}
