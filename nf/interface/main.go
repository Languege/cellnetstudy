package main

import "fmt"

type Animal interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string {
	return "cat"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "dog"
}

func Test(params interface{}) {
	fmt.Println(params)
}

func main() {
	animals := []Animal{Cat{}, &Dog{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	var dog Animal = new(Dog)
	animal, ok := dog.(Animal)
	fmt.Println(animal, ok)

	switch dog.(type) {

	case *Cat:
		fmt.Println("*Cat")
	case *Dog:
		fmt.Println("*Dog")
	}

}
