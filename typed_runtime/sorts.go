package runtime

import (
	"el/sorts"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Sort = sorts.Sort

var objectDict = &_objectDict{data: make(map[string]Object)}

type _objectDict struct {
	data map[string]Object
}

func (d *_objectDict) Iter(yield func(name string, dtype Object) bool) {
	for name, obj := range d.data {
		if ok := yield(name, obj); !ok {
			return
		}
	}
}

func (d *_objectDict) MakeType(name string) Object {
	const typeLevel = 1
	o := _object{
		data:   nil,
		sort:   sorts.MustAtom(typeLevel, name, nil),
		parent: nil, // default parent
	}
	d.data[name] = o
	return o
}

func (d *_objectDict) MakeData(data Data, parent Object) Object {
	const dataLevel = 0
	return _object{
		data:   data,
		sort:   sorts.MustAtom(dataLevel, data.String(), parent.Sort()),
		parent: parent.Sort(),
	}
}

// _object - unorder-score means private, even in the same package
type _object struct {
	data   Data // nullable
	sort   Sort
	parent Sort
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
		data:   nil,
		sort:   o.parent,
		parent: o.parent.Parent(),
	}
}
func (o _object) Cast(newParent Object) adt.Option[Object] {
	newParentObject := newParent.(_object) // must cast
	if ok := o.parent.LessEqual(newParentObject.sort); !ok {
		return adt.None[Object]()
	}
	return adt.Some[Object](_object{
		data:   o.data,
		sort:   o.sort,
		parent: o.parent.Parent(),
	})
}
