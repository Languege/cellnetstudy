package main

import (
	"fmt"
	"time"
)

func statusUpdate() string {
	promt := "hahaha"

	return promt
}

func main() {
	c := time.Tick(3 * time.Second)
	for now := range c {
		fmt.Printf("%+v %s\n", now, statusUpdate())
	}
}
