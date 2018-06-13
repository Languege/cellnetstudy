package main

import (
	"fmt"
	"reflect"
)

type I interface {
	A()
}

type S struct {
	id string
}

func (s *S) A() {
	println("A")
}

func main() {

	iType := reflect.TypeOf((*I)(nil)).Elem()
	fmt.Println(iType)

	aType := reflect.TypeOf((*S)(nil))
	fmt.Println(aType)

	fmt.Println(aType.Implements(iType))
}
