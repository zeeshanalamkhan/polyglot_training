package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func evenOdd(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			println(i, " Even")
		} else {
			println(i, " Odd")
		}
	}
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	println("--------")
	evenOdd(10)
	println("--------")
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// If with a short statement
func pow(x, n, lim float64) float64 {
	//Like for, the if statement can start with a short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if.
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
