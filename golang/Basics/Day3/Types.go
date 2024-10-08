package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint16     = 1<<8 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var msg string = "Testing"

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println("Message : " + msg)
	// Example of un-initialized variable
	fmt.Println("----------------")
	variable_not_init()
	// Example of constant
	fmt.Println("----------------")
	constantTest()
}

/*
Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
*/

func variable_not_init() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("Default values: %v %v %v %q\n", i, f, b, s)
}

/*
Constants are declared like variables, but with the const keyword.
Constants can be character, string, boolean, or numeric values.
Constants cannot be declared using the := syntax.
*/

func constantTest() {
	const Pi = 3.14
	const World = "PS_LBG"
	fmt.Println("Hello ", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
