package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")[0])
	fmt.Println(strings.Contains("seafood", "fdsafdsa"))
}
