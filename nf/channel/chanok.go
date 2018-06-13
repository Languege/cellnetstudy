//testrange.go
package main

import "fmt"

func send(strings chan string) {
	strings <- "Five hour's New York jet lag"
	// strings <- "and Cayce Pollard wakes in Camden Town"
	// strings <- "to the dire and ever-decreasing circles"
	// strings <- "of disrupted circadian rhythm."
	close(strings)
}

func rev(strings chan string) {

	s, ok := <-strings
	fmt.Print(s, ok)
}

func main() {
	strings := make(chan string)
	go send(strings)
	s, ok := <-strings

	fmt.Print(s, ok)

	s2, ok2 := <-strings

	fmt.Print(s2, ok2)
}
