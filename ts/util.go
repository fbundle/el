package ts

const (
	// Initial - Initial can be cast into any type, it has a no value or a single zero value depends on the category
	Initial = "unit"
	// Final - every type can be cast into Final, it is like Any
	Final = "any"
)

type rule struct {
	srcName string
	dstName string
}

var leMap = make(map[rule]struct{})

func AddRule(srcName string, dstName string) {
	leMap[rule{srcName, dstName}] = struct{}{}
}

func leName(srcName string, dstName string) bool {
	if srcName == Initial || dstName == Final {
		return true
	}
	if srcName == dstName {
		return true
	}
	if _, ok := leMap[rule{srcName, dstName}]; ok {
		return true
	}
	return false
}

type Name string

func (n Name) String() string {
	return string(n)
}
