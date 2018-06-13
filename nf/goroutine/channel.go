package main

import "fmt"

//import "time"

func sum(a []int, c chan int, biaoshi string) {
	fmt.Printf("biaoshi:%s\n", biaoshi)
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c, "befor")
	go sum(a[len(a)/2:], c, "after")

	//time.Sleep(1)
	x := <-c
	y := <-c // receive from c

	fmt.Println(x, y, x+y)
}
