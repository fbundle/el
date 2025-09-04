package sorts

const (
	// Initial - Initial can be cast into any type, it has a no value or a atom zero value depends on the category
	Initial = "unit"
	// Terminal - every type can be cast into Terminal, it is like Any
	Terminal = "any"
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
	if src == Initial || dst == Terminal {
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
