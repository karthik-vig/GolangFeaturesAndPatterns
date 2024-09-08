package generics

import (
	"fmt"
)

type intTypeAlias int

type unionTypeAliasInterface interface {
	intTypeAlias | ~int32
	int32
}

type GenericStruct[T string | rune, K ~int | ~int32, X unionTypeAliasInterface] struct {
	someValue1    K
	someValue     T
	someStringVal string
}

func (s *GenericStruct[T, K, X]) CompareWithStructVal(val T) bool {
	return val == s.someValue
}

func (s *GenericStruct[T, K, X]) PrintStructValues() {
	fmt.Println("The string value is: ", s.someStringVal)
	switch any(s.someValue).(type) {
	case string:
		fmt.Println("It is string")
	case rune:
		fmt.Println("It is an rune")
	}
	fmt.Println("The generic values is: ", s.someValue)
}

func PrintGenericValue[T any](val T) {
	fmt.Println(val)
}

func greatFunc[T rune | byte, V []T](val T) {
	// code here...
}

func TestGenerics() {
	structVar1 := GenericStruct[string, intTypeAlias, int32]{someValue: "hello world", someStringVal: "hey there", someValue1: 0}
	structVar2 := GenericStruct[rune, intTypeAlias, int32]{someValue: 'a', someStringVal: "hey there 2", someValue1: 1}
	PrintGenericValue("hey man") //type has been inferred
	fmt.Println("struct var 1 compare: ", structVar1.CompareWithStructVal("hello world"))
	fmt.Println("struct var 2 compare: ", structVar2.CompareWithStructVal('a'))
	structVar1.PrintStructValues()
	structVar2.PrintStructValues()
	greatFunc[rune]('a')
}
