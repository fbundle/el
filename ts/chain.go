package ts

import "github.com/fbundle/lab_public/lab/go_util/pkg/adt"

var _ Sort = chain{}

// chain - represent arrow type A -> B -> C
type chain struct {
	parm adt.NonEmptySlice[Sort]
	ret  Sort
}

func (c chain) Data() adt.Option[Data] {
	//TODO implement me
	panic("implement me")
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

func (c chain) Cast(parent Sort) adt.Option[Sort] {
	//TODO implement me
	panic("implement me")
}

func (c chain) Len() int {
	return len(c.parm.Repr()) + 1
}

func (c chain) le(dst Sort) bool {
	//TODO implement me
	panic("implement me")
}

func (c chain) prepend(param Sort) Sort {
	//TODO implement me
	panic("implement me")
}
