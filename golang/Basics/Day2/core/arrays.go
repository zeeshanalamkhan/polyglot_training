package main

import "fmt"

func main() {
	var cities [3]string
	cities[0] = "hyderabad"
	cities[1] = "Bengaluru"
	cities[2] = "Delhi"
	fmt.Println(cities[2])

	cities[2] = "Gurgav"
	fmt.Println(cities[2])

	fmt.Println(len(cities))

	cityNames := [...]string{"Kolkata", "Pune", "Mumbai", "Delhi"}
	fmt.Println(len(cityNames))

	for index, value := range cities {
		fmt.Println(index, "-->", value)
	}
}
