package main

import (
	"el/ts"
	"fmt"
)

func newData(v any) ts.Data {
	return data{value: v}
}

type data struct {
	value any
}

func (d data) String() string {
	return fmt.Sprint(d.value)
}

const (
	dataLevel = 0
	typeLevel = 1
)

func makeNameType(typeName string) ts.Sort {
	return ts.MustSingleName(typeLevel, typeName)
}

func makeData(data ts.Data, dtype ts.Sort) ts.Sort {
	return ts.MustSingleData(dataLevel, data, dtype)
}

func printSorts(sorts ...ts.Sort) {
	for _, sort := range sorts {
		fmt.Printf("[%s] is of type [%s]\n", sort.String(), sort.Type().String())
	}
}
func printCast(value1 ts.Sort, type2 ts.Sort) {
	ok := value1.Cast(type2).Ok

	if ok {
		fmt.Printf("value [%s] of type [%s] CAN be cast into [%s]\n", value1, value1.Type(), type2)
	} else {
		fmt.Printf("value [%s] of type [%s] CANNOT be cast into [%s]\n", value1, value1.Type(), type2)
	}
}

func strongestType(length int) ts.Sort {
	// this type can be cast into every type of this length
	if length < 1 {
		panic("type_error")
	}
	var sorts []ts.Sort
	for i := 0; i < length-1; i++ {
		sorts = append(sorts, ts.MustSingleName(typeLevel, ts.Initial))
	}
	sorts = append(sorts, ts.MustSingleName(typeLevel, ts.Final))

	return ts.MustChain(sorts...)
}

func weakestType(length int) ts.Sort {
	// every type of this length can be cast into this type
	if length < 1 {
		panic("type_error")
	}
	var sorts []ts.Sort
	for i := 0; i < length-1; i++ {
		sorts = append(sorts, ts.MustSingleName(typeLevel, ts.Final))
	}
	sorts = append(sorts, ts.MustSingleName(typeLevel, ts.Initial))

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

	weak3Sort := makeData(newData(nil), weak3)
	strong3Sort := makeData(newData(nil), strong3)

	ts.AddRule("bool", "int") // cast bool -> int
	fmt.Println("[bool] can be cast into [int]")

	oneSort := makeData(newData(1), intType)
	trueSort := makeData(newData(true), boolType)
	helloSort := makeData(newData("hello"), stringType)
	add1 := makeData(newData(func(i int) int {
		return i + 1
	}), intIntType)
	add := makeData(newData(func(i int) func(int) int {
		return func(j int) int {
			return i + j
		}
	}), intIntIntType)

	printSorts(oneSort, trueSort, helloSort, add1, add, weak1, strong1, weak3, strong3)

	printCast(trueSort, intType)
	printCast(oneSort, boolType)
	printCast(weak3Sort, add.Type())
	printCast(strong3Sort, add.Type())
}
