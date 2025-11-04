package main

import "fmt"

func call(m1 string) { messages <- m1 }

var messages = make(chan string)

func main() {

	go call("by")

	msg := <-messages
	fmt.Println(msg)

	messages1 := make(chan string)

	go func() { messages1 <- "ping" }()

	msg1 := <-messages1
	fmt.Println(msg1)
}
