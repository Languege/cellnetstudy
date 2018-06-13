package main

import (
	"fmt"
	"github.com/hyperOx/goc2p/src/basic/set"
)

func main() {
	fmt.Print("这是用来测试hash_set的\n")

	set1 := set.NewHashSet()

	set1.Add("string")
	set1.Add(1)
	set1.Add(true)

	fmt.Printf("\n%s\n", set1.String())

	fmt.Printf("\n%v\n", set1.Elements())

	//set1.Remove(1)

	fmt.Printf("\n%v\n", set1.Elements())

	fmt.Printf("\n%t\n", set1.Contains("string"))

	fmt.Printf("\n%d\n", set1.Len())

	set2 := set.NewHashSet()
	set2.Add(true)
	set2.Add("string")

	fmt.Printf("\n%t\n", set2.Same(set1))

	fmt.Printf("\n%t\n", set.IsSuperset(set1, set2))

	set3 := set.Union(set1, set2)

	fmt.Printf("\n%v\n", set3.Elements())

	set4 := set.Difference(set1, set2)

	fmt.Printf("\n%v\n", set4.Elements())

	set5 := set.SymmetricDifference(set1, set2)

	fmt.Printf("\n%v\n", set5.Elements())

	set6 := set.Intersect(set1, set2)

	fmt.Printf("\n%v\n", set6.Elements())
}
