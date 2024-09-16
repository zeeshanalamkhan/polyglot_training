/*
Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.
To invoke this function in a goroutine, use go f(s).
This new goroutine will execute concurrently with the calling one.
You can also start a goroutine for an anonymous function call.
Our two function calls are running asynchronously in separate goroutines now.
Wait for them to finish (for a more robust approach, use a WaitGroup).
When we run this program, we see the output of the blocking call first, then the output of the two goroutines.
The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.
*/
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 4; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		for i := 0; i < 5; i++ {
			fmt.Println(msg, ":", i)
			time.Sleep(time.Second)
		}

	}("goroutine inner f1")

	time.Sleep(time.Second * 5)

	/**
	QUIZ : Why below code is not showing any output while
	it is exact same as above function
	**/
	go func(msg string) {
		for i := 0; i < 5; i++ {
			fmt.Println(msg, ":", i)
			time.Sleep(time.Second)
		}

	}("goroutine inner f2")
	fmt.Println("done")

}
