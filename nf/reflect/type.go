package main

import (
	"fmt"
	"reflect"
	"strings"
)

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

// 声明一个空结构体
type Cat struct {
	Name string
}

func main() {

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(Cat{Name: "hellokitty"})

	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)

	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	if strings.Compare(typeOfCat.Kind().String(), "struct") == 0 {
		num_feild := typeOfCat.NumField()

		for i := 0; i < num_feild; i++ {
			fmt.Print(typeOfCat.Field(i))
		}

	}

}
