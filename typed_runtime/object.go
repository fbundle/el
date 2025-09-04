package runtime

import (
	"el/sorts"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Sort = sorts.Sort

type Data interface {
	String() string
}

type Object interface {
	Data() Data
	String() string
	Type() Object
	Cast(dtype Object) adt.Option[Object]
}

func makeType(name string) object {
	const typeLevel = 1
	return object{
		data:   nil,
		sort:   sorts.MustAtom(typeLevel, name, nil),
		parent: nil, // default parent
	}
}

func makeData(data Data, dtype Sort) object {
	const dataLevel = 0
	return object{
		data:   data,
		sort:   sorts.MustAtom(dataLevel, data.String(), dtype),
		parent: dtype,
	}
}

type object struct {
	data   Data // nullable
	sort   Sort
	parent Sort
}

func (o object) Data() Data {
	return o.data
}

func (o object) String() string {
	return o.sort.String()
}

func (o object) Type() Object {
	return object{
		data:   nil,
		sort:   o.parent,
		parent: o.parent.Parent(),
	}
}
func (o object) Cast(newParent Object) adt.Option[Object] {
	var newParentObject object
	if ok := adt.Cast[object](newParent).Unwrap(&newParentObject); !ok {
		return adt.None[Object]()
	}
	if ok := o.parent.LessEqual(newParentObject.sort); !ok {
		return adt.None[Object]()
	}
	return adt.Some[Object](object{
		data:   o.data,
		sort:   o.sort,
		parent: o.parent.Parent(),
	})
}
