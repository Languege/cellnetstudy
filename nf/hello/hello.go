package main

import (
	"fmt"
	"github.com/hyperOx/goc2p/src/basic"
	// "reflect"
	// "sort"
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

const (
	HELLO string = "malijun"
)

// type Sortable interface {
// 	sort.Interface
// 	Sort()
// }
type SortableStrings struct {
	List []string
}

func (s *SortableStrings) Add(value string) {
	s.List = append(s.List, value)
}

// func (a SortableStrings) Len() int           { return len(a) }
// func (a SortableStrings) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortableStrings) Less(i, j int) bool { return a[i] < a[j] }
// func (a SortableStrings) Sort()              { sort.Sort(a) }

// type Sequence struct { //有点继承的味道
// 	Sortable //匿名字段，Sequence结构体变成可以直接调用Sortable中的方法和字段
// 	sorted   bool
// }

// func (self *Sequence) name() {
// 	self.Sortable.Sort()
// 	self.sorted = true
// }

// type Person struct {
// 	Name    string `json:"name"`
// 	Age     uint8  `json:"age"`
// 	Address string `json:"address"`
// }

// type ByAge []Person

// //ByAge implements sort.Interface for []Person based on
// // the Age field.
// func (a ByAge) Len() int           { return len(a) }
// func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func modify(n int) (number int) {
	defer func() {
		number += n
	}()
	number++
	return
}

/** 以下函数用于测试 panic recover 和 defer的运行 start 功能和php中的try catch类似*/
func fetchDemo() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Recovered a panic. [index=%d]\n", v)
		}
	}()
	ss := []string{"A", "B", "C"}
	fmt.Printf("fetch the elements in %v one by one..\n", ss)
	fetchElement(ss, 0)
	fmt.Println("The elements fetching is done.")
}

func fetchElement(ss []string, index int) (element string) {
	if index >= len(ss) {
		fmt.Printf("Occur a panic![index=%d]\n", index)
		panic(index)
	}
	fmt.Printf("Fetching the element... [index=%d]\n", index)
	element = ss[index]
	defer fmt.Printf("The element is %s . [index=%d]\n", element, index)
	fetchElement(ss, index+1)
	return
}

/** 以上函数用于测试 panic recover 和 defer的运行 end */
func main() {
	fmt.Println(modify(2))
	// name := "liugaoyun"
	// fmt.Printf("\nHello, %s, %s", name, HELLO)
	// //切片
	// array1 := [6]string{"Go", "Python", "Java", "C", "C++", "PHP"}
	// slice1 := array1[:4:4]
	// fmt.Printf("\nslice1.len:%d, cap:%d", len(slice1), cap(slice1))
	// fmt.Print(array1)

	// slice1 = append(slice1, "Ruby", "Erlang")
	// fmt.Printf("\nslice1.len:%d, cap:%d", len(slice1), cap(slice1))
	// fmt.Print(array1)

	// sliceA := []string{"Notepad", "UltraEdit", "Eclipse"}
	// sliceB := []string{"Vim", "Emacs", "LiteIDE", "IDEA"}
	// n1 := copy(sliceA, sliceB)
	// fmt.Println()
	// fmt.Print(sliceA)
	// fmt.Printf("成功复制%d个元素", n1)
	// n2 := copy(sliceB, sliceA)
	// fmt.Println()
	// fmt.Print(sliceB)
	// fmt.Printf("成功复制%d个元素", n2)

	//结构体
	// var sequenece Sequence
	// fmt.Print(sequenece)

	// anonym := struct {
	// 	a int
	// 	b string
	// }{a: 1, b: "mdzz"}

	// fmt.Println()
	// fmt.Print(anonym)
	// person := Person{Name: "liugaoyun", Age: 24, Address: "118757456@qq.com"}

	// fmt.Println()
	// fmt.Print(person)
	// for i := 0; i < reflect.TypeOf(person).NumField(); i++ {
	// 	fmt.Println()
	// 	fmt.Print(reflect.TypeOf(person).Field(i))
	// }
	// seq := Sequence{Sortable: SortableStrings{"3", "2", "1"}, sorted: false}
	// seq.Sort()
	// fmt.Println()
	// fmt.Print(seq)

	// people := []Person{
	// 	{"Bob", 31, "1"},
	// 	{"John", 31, "3"},
	// 	{"Michael", 17, "2"},
	// 	{"Jenny", 26, "3"},
	// }

	// fmt.Println(people)
	// sort.Sort(ByAge(people))
	// fmt.Println(people)
	//seq_value := basic.StringSeq{basic.StringSeq.value: []string{"i", "am", "a", "good", "student"}}
	//seq := basic.Sequence{GenericSeq: string_seq, sorted: false}
	var seq basic.Sequence
	var seq_string_value basic.StringSeq
	// seq_string_value.Append("1")
	// seq_string_value.Append("2")
	// seq_string_value.Append("3")

	seq.GenericSeq = &seq_string_value

	seq.Init()
	seq.Append("i am a good student")
	seq.Append("he is a good student")
	seq.Append("she is a good student")
	seq.Append("you are a good student")
	fmt.Println(seq.ElemType())
	fmt.Println(seq.Sorted())
	seq.Sort()
	fmt.Println(seq.GenericSeq)

	//接口
	//1.验证basic.StringSeq 实现了 basic.GenericSeq接口
	_, ok := interface{}(&seq_string_value).(basic.Sortable)
	fmt.Println(ok)

	//函数
	test := SortableStrings{List: []string{"1", "2", "3"}}
	test.Add("4")
	fmt.Println(test.List)

	//4.2.1 error
	var err error = errors.New("A normal error.")
	fmt.Println(err)

	file, err3 := os.Open("test.txt")
	if err3 != nil {
		if pe, ok := err3.(*os.PathError); ok {
			fmt.Printf("Path Error:%s (op=%s, path=%s)\n", pe.Err, pe.Op, pe.Path)
		} else {
			fmt.Printf("Unknown Error:%s\n", err)
		}
	}
	r := bufio.NewReader(file)
	var buf bytes.Buffer
	for {
		byteArray, _, err4 := r.ReadLine()
		if err4 != nil {
			if err4 == io.EOF {
				break
			} else {
				fmt.Printf("Read Error:%s\n", err4)
				break
			}
		} else {
			buf.Write(byteArray)
		}
	}
	//panic("玛丽 转过去 屁股太高点 让我草死个死贱逼")
	fmt.Println(buf.String())

	fetchDemo()
}
