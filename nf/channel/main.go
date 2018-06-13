package main

import (
	"fmt"
)

func sum1(values []int, resultChan chan int, startSignal chan int) {
	<-startSignal
	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Println("sum1")
	//将计算结果发到channel中
	resultChan <- sum
}

func sum2(values []int, resultChan chan int, startSignal chan int) {
	<-startSignal
	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Println("sum2")
	//将计算结果发到channel中
	resultChan <- sum
}

func sum3(values []int, resultChan chan int, startSignal chan int) {
	<-startSignal
	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Println("sum3")
	//将计算结果发到channel中
	resultChan <- sum
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 3)
	startSignal := make(chan int)
	// var resultChan chan int
	go sum1(values[:len(values)/2], resultChan, startSignal)
	go sum2(values[len(values)/2:], resultChan, startSignal)
	go sum3(values[len(values)/3:], resultChan, startSignal)
	// sum1, sum2, sum3 := <-resultChan, <-resultChan, <-resultChan
	close(startSignal)

	sum := <-resultChan
	fmt.Println("Result:", sum)

	sum2 := <-resultChan
	fmt.Println("Result:", sum2)

	sum3 := <-resultChan
	fmt.Println("Result:", sum3)
}
