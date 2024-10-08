package main

import "fmt"

// variable constant
var greet string
var name = "learning@lbg"

func main() {
	fmt.Println("greet is " + greet)
	greet = "hello "
	// declare local variable
	msg := greet + name
	fmt.Println(msg)
}
