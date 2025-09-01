package main

import (
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/ordered_map"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/seq"
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/stack"
)

var Seq = seq.Empty[int]()
var OrderedMap = ordered_map.EmptyOrderedMap[int, int]()
var Stack = stack.Empty[int]()

func main() {
	fmt.Println("Hello, World!")
}
