package main

import (
	"el/pkg/el/persistent"
	"fmt"
)

func testPersistent() {
	l := persistent.EmptyList[int]()
	l = l.Push(1).Push(2).Push(3)
	fmt.Println(l.Repr())

	m := persistent.EmptyMap[int, string]()
	m = m.Set(1, "one").Set(2, "two").Set(3, "three")
	fmt.Println(m.Repr())
}

func main() {
	testPersistent()
}
