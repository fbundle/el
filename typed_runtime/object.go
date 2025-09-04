package runtime

import (
	"el/ts"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Sort = ts.Sort

type Data interface {
	String() string
}

type Object interface {
	Data() Data
	String() string
	Type() Object
	Cast(parentType Object) adt.Option[Object]
}

func makeDataObject(data Data, parent Sort) object {
	const dataLevel = 0
	return object{
		data:   data,
		sort:   ts.MustSingleName(dataLevel, data.String()),
		parent: parent,
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
