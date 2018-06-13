package main

import (
	"fmt"
	"reflect"
)

type Dog struct {
	LegCount int
}

func main() {

	// 声明整型变量a并赋初值
	var a int = 1024

	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(&a).Elem()

	// 尝试将a修改为1(此处会发生崩溃)
	valueOfA.SetInt(1)

	fmt.Println(a)

	dog := Dog{LegCount: 1}

	// 获取legCount字段的反射值对象
	vLegCount := reflect.ValueOf(&dog).Elem().FieldByName("LegCount")

	vLegCount.SetInt(4)

	fmt.Printf("%+v", dog)
}
