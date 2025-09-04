package runtime

import (
	"el/sorts"
	"sync"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Sort = sorts.Sort

type ObjectPool interface {
	Iter(yield func(name string, dtype Object) bool)
	MakeType(name string) Object
	MakeData(data Data, parent Object) Object
}

var Pool ObjectPool = &_sortDict{objectDict: make(map[string]Object)}

type _sortDict struct {
	mu         sync.RWMutex
	objectDict map[string]Object
}

func (d *_sortDict) Iter(yield func(name string, dtype Object) bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	for name, obj := range d.objectDict {
		if ok := yield(name, obj); !ok {
			return
		}
	}
}

func (d *_sortDict) MakeType(name string) Object {
	const typeLevel = 1
	o := _object{
		data:   nil,
		sort:   sorts.MustAtom(typeLevel, name, nil),
		parent: nil, // default parent
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	d.objectDict[o.String()] = o
	return o
}

func (d *_sortDict) MakeData(data Data, parent Object) Object {
	const dataLevel = 0
	o := _object{
		data:   data,
		sort:   sorts.MustAtom(dataLevel, data.String(), parent.Sort()),
		parent: parent.Sort(),
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	d.objectDict[o.String()] = o
	return o
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
