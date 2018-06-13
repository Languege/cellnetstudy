package main

import (
	"errors"
	"fmt"
	"time"
)

func RPCClient(ch chan string, req string) (string, error) {

	ch <- req

	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("time out")
	}
}

func RPCServer(ch chan string) {

	for {

		data := <-ch

		fmt.Println("server received", data)

		time.Sleep(time.Millisecond)

		ch <- "ack"
	}
}

func main() {

	ch := make(chan string)

	go RPCServer(ch)

	response, err := RPCClient(ch, "mdzz")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
