package main

import (
	"el/sorts"
	"fmt"
)

const (
	typeLevel = 1
)

func makeNameType(typeName string) sorts.Sort {
	return sorts.MustAtom(typeLevel, typeName, nil)
}

func printSorts(sorts ...sorts.Sort) {
	for _, sort := range sorts {
		fmt.Printf("[%s] is of type [%s]\n", sort.String(), sort.Parent().String())
	}
}
func printCast(type1 sorts.Sort, type2 sorts.Sort) {
	ok := type1.LessEqual(type2)

	if ok {
		fmt.Printf("type [%s] CAN be cast into [%s]\n", type1, type2)
	} else {
		fmt.Printf("type [%s] CANNOT be cast into [%s]\n", type1, type2)
	}
}

func strongestType(length int) sorts.Sort {
	// this type can be cast into every type of this length
	if length < 1 {
		panic("type_error")
	}
	var ss []sorts.Sort
	for i := 0; i < length-1; i++ {
		ss = append(ss, sorts.MustAtom(typeLevel, sorts.Initial, nil))
	}
	ss = append(ss, sorts.MustAtom(typeLevel, sorts.Terminal, nil))

	return sorts.MustArrow(ss...)
}

func weakestType(length int) sorts.Sort {
	// every type of this length can be cast into this type
	if length < 1 {
		panic("type_error")
	}
	var ss []sorts.Sort
	for i := 0; i < length-1; i++ {
		ss = append(ss, sorts.MustAtom(typeLevel, sorts.Terminal, nil))
	}
	ss = append(ss, sorts.MustAtom(typeLevel, sorts.Initial, nil))

	return sorts.MustArrow(ss...)
}

func main() {
	fmt.Printf("anything can be cast into [%s]\n", sorts.Terminal)
	fmt.Printf("[%s] can be cast into anything\n", sorts.Initial)

	intType := makeNameType("int")
	boolType := makeNameType("bool")
	stringType := makeNameType("string")
	intIntType := sorts.MustArrow(intType, intType)
	intIntIntType := sorts.MustArrow(intType, intType, intType)
	weak1 := weakestType(1)
	strong1 := strongestType(1)
	weak3 := weakestType(3)
	strong3 := strongestType(3)

	sorts.AddRule("bool", "int") // cast bool -> int
	fmt.Println("[bool] can be cast into [int]")

	printSorts(stringType, intIntType, intIntIntType, weak1, strong1, weak3, strong3)

	printCast(boolType, intType)
	printCast(intType, boolType)
	printCast(weak3, intIntType)
	printCast(strong3, intIntType)
}
