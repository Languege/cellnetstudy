package main

import (
	"fmt"
)

func main() {
	sliceA := []string{"Notepad", "UltraEdit", "Eclipse"}
	sliceB := []string{"Vim", "Emacs", "LiteIDE", "IDEA"}
	n1 := copy(sliceA, sliceB)
	fmt.Println()
	fmt.Print(sliceA)
	fmt.Printf("成功复制%d个元素", n1)

	sliceC := []string{"Notepad", "UltraEdit", "Eclipse"}
	sliceD := []string{"Vim", "Emacs", "LiteIDE", "IDEA"}
	n2 := copy(sliceD, sliceC)
	fmt.Println()
	fmt.Print(sliceD)
	fmt.Printf("成功复制%d个元素", n2)
}
