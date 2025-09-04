package ts

const (
	// Initial - Initial can be cast into any type, it has a no value or a single zero value depends on the category
	Initial = "unit"
	// Final - every type can be cast into Final, it is like Any
	Final = "any"
)

type rule struct {
	src string
	dst string
}

var leMap = make(map[rule]struct{})

func AddRule(src string, dst string) {
	leMap[rule{src, dst}] = struct{}{}
}

func le(src string, dst string) bool {
	if src == Initial || dst == Final {
		return true
	}
	if src == dst {
		return true
	}
	if _, ok := leMap[rule{src, dst}]; ok {
		return true
	}
	return false
}
