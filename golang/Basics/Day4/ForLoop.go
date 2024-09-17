package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println("-----")
	fmt.Println(sum)
	fmt.Println("-----")
	loop1()
}

// The init and post statements are optional.
func loop1() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
