package runtime

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
	"github.com/fbundle/sorts/sorts_v1"
)

const _typeLevel = 1
const _dataLevel = 0

type Sort = sorts.Sort

const (
	Unit = sorts.Unit
	Any  = sorts.Any
)

var Arrow = sorts.Arrow

func MakeType(name string) Object {
	o := _object{
		data: nil,
		sort: sorts.MustAtom(_typeLevel, name, nil),
	}
	return o
}

func MakeData(data Data, parent Object) Object {
	return _object{
		data: data,
		sort: sorts.MustAtom(_dataLevel, data.String(), parent.Sort()),
	}
}

func MakeSort(sort Sort) Object {
	return _object{
		data: nil, // no data
		sort: sort,
	}
}

func makeWeakestType(numParams int) Object {
	// every type of this length can be cast into this type
	if numParams < 0 {
		panic("type_error")
	}
	var ss []sorts.Sort
	for i := 0; i < numParams; i++ {
		ss = append(ss, sorts.MustAtom(_typeLevel, sorts.Any, nil))
	}
	ss = append(ss, sorts.MustAtom(_typeLevel, sorts.Unit, nil))
	s := sorts.MustArrow(ss...)

	// convert sort to object
	return _object{
		data: nil,
		sort: s,
	}

}

// _object - unorder-score means private, even in the same package
type _object struct {
	data Data // nullable
	sort Sort // hold the sort of object
}

func (o _object) Data() Data {
	return o.data
}

func (o _object) String() string {
	return o.sort.String()
}
func (o _object) Sort() Sort {
	return o.sort
}
func (o _object) Type() Object {
	return _object{
		data: nil,
		sort: o.sort.Parent(),
	}
}
func (o _object) Cast(newParent Object) adt.Option[Object] {
	newParentObject := newParent.(_object) // must cast
	if ok := o.sort.Parent().LessEqual(newParentObject.sort); !ok {
		return adt.None[Object]()
	}

	return adt.Some[Object](MakeData(o.data, newParentObject))
}
