package main

import (
	"fmt"
)

func main() {
	
	ch := make(chan int)

	go func() {
		fmt.Println("goroutine start")

		ch <- 0

		fmt.Println("goroutine end")
	}()

	fmt.Println("wait goroutine")

	<-ch

	fmt.Println("all done")
}
