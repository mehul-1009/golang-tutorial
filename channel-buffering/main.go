package main

import "fmt"

func main() {

	messages := make([]int, 0, 5)

	messages <- 1
	messages <- 2

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
