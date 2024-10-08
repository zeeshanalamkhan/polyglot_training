package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	//deferstack()
}

// Deferred function calls are pushed onto a stack.
// When a function returns, its deferred calls are executed in last-in-first-out order.
func deferstack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
