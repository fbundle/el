package main

import (
	"el/ts"
	"fmt"
)

const (
	typeLevel = 1
)

func makeNameType(typeName string) ts.Sort {
	return ts.MustSingle(typeLevel, typeName)
}

func printSorts(sorts ...ts.Sort) {
	for _, sort := range sorts {
		fmt.Printf("[%s] is of type [%s]\n", sort.String(), sort.Parent().String())
	}
}
func printCast(type1 ts.Sort, type2 ts.Sort) {
	ok := type1.LessEqual(type2)

	if ok {
		fmt.Printf("type [%s] CAN be cast into [%s]\n", type1, type2)
	} else {
		fmt.Printf("type [%s] CANNOT be cast into [%s]\n", type1, type2)
	}
}

func strongestType(length int) ts.Sort {
	// this type can be cast into every type of this length
	if length < 1 {
		panic("type_error")
	}
	var sorts []ts.Sort
	for i := 0; i < length-1; i++ {
		sorts = append(sorts, ts.MustSingle(typeLevel, ts.Initial))
	}
	sorts = append(sorts, ts.MustSingle(typeLevel, ts.Final))

	return ts.MustChain(sorts...)
}

func weakestType(length int) ts.Sort {
	// every type of this length can be cast into this type
	if length < 1 {
		panic("type_error")
	}
	var sorts []ts.Sort
	for i := 0; i < length-1; i++ {
		sorts = append(sorts, ts.MustSingle(typeLevel, ts.Final))
	}
	sorts = append(sorts, ts.MustSingle(typeLevel, ts.Initial))

	return ts.MustChain(sorts...)
}

func main() {
	fmt.Printf("anything can be cast into [%s]\n", ts.Final)
	fmt.Printf("[%s] can be cast into anything\n", ts.Initial)

	intType := makeNameType("int")
	boolType := makeNameType("bool")
	stringType := makeNameType("string")
	intIntType := ts.MustChain(intType, intType)
	intIntIntType := ts.MustChain(intType, intType, intType)
	weak1 := weakestType(1)
	strong1 := strongestType(1)
	weak3 := weakestType(3)
	strong3 := strongestType(3)

	ts.AddRule("bool", "int") // cast bool -> int
	fmt.Println("[bool] can be cast into [int]")

	printSorts(stringType, intIntType, intIntIntType, weak1, strong1, weak3, strong3)

	printCast(boolType, intType)
	printCast(intType, boolType)
	printCast(weak3, intIntType)
	printCast(strong3, intIntType)
}
