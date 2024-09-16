package main

import (
	"fmt"
)

/*
 */
func main() {

	messages := make(chan string)

	go func() {
		fmt.Print("working...")
		messages <- "test"
		messages <- "........" + (<-messages)
		fmt.Println("done with inner function")

	}()

	go task1(messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func task1(messages chan string) {
	messages <- "doing some work"
	fmt.Print("working on task1...")
	fmt.Println("done with outer function")
}
