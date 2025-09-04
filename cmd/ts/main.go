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
		fmt.Printf("%s \t--- type %s\n", sort.String(), sort.Type().String())
	}
}

func main() {
	intType := makeNameType("int")
	boolType := makeNameType("bool")
	stringType := makeNameType("string")
	intIntType := ts.MustChain(intType, intType)
	intIntIntType := ts.MustChain(intType, intType, intType)

	ts.AddRule("bool", "int") // cast bool -> int

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

	printSorts(oneSort, trueSort, helloSort, add1, add)

}
