package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// 将结构体序列化为JSON格式并输出到缓冲中
func writeStruct(buff *bytes.Buffer, value reflect.Value) error {

	// 取值的类型对象
	valueType := value.Type()

	// 写入结构体左大括号
	buff.WriteString("{")

	// 遍历结构体的所有值
	for i := 0; i < value.NumField(); i++ {

		// 获取每个字段的字段值(reflect.Value)
		fieldValue := value.Field(i)

		// 获取每个字段的类型(reflect.StructField)
		fieldType := valueType.Field(i)

		// 写入字段名左双引号
		buff.WriteString("\"")

		// 写入字段名
		buff.WriteString(fieldType.Name)

		// 写入字段名右双引号和冒号
		buff.WriteString("\":")

		// 写入每个字段值
		writeAny(buff, fieldValue)

		//每个字段尾部写入逗号, 最后一个字段不添加
		if i < value.NumField()-1 {
			buff.WriteString(",")
		}
	}

	// 写入结构体右大括号
	buff.WriteString("}")

	return nil
}

// 将切片转换为JSON格式并输出到缓冲中
func writeSlice(buff *bytes.Buffer, value reflect.Value) error {

	// 写入切片开始标记
	buff.WriteString("[")

	// 遍历每个切片元素
	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)

		// 写入每个切片元素
		writeAny(buff, sliceValue)

		//每个元素尾部写入逗号, 最后一个字段不添加
		if s < value.Len()-1 {
			buff.WriteString(",")
		}
	}

	// 写入切片结束标记
	buff.WriteString("]")

	return nil
}

// 将任意值转换为JSON格式并输出到缓冲中
func writeAny(buff *bytes.Buffer, value reflect.Value) error {

	switch value.Kind() {
	case reflect.String:
		// 写入带有双引号括起来的字符串
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		// 将整型转换为字符串并写入缓冲中
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buff, value)
	case reflect.Struct:
		return writeStruct(buff, value)
	default:
		// 遇到不认识的种类, 返回错误
		return errors.New("unsupport kind: " + value.Kind().String())
	}

	return nil
}

// 给外部使用序列化值为JSON的接口
func MarshalJson(v interface{}) (string, error) {
	// 准备一个缓冲
	var b bytes.Buffer

	// 将任意值转换为JSON并输出到缓冲中
	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {
		return b.String(), nil
	} else {
		return "", err
	}
}

// 声明技能结构
type Skill struct {
	Name  string
	Level int
}

// 声明角色结构
type Actor struct {
	Name string
	Age  int

	Skills []Skill
}

func main() {

	// 填充基本角色数据
	a := Actor{
		Name: "cow boy",
		Age:  37,

		Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}

	if result, err := MarshalJson(a); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
