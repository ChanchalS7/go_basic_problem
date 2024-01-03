package main

import (
	"fmt"
)

// returning a single value of type int
func add(x, y int) int {
	return x + y
}

// Named return value
func getArea(l int, b int) (area int) {
	area = l * b
	return
}

// returning multiple name values
func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return
}

// passing address to a function
func addValue(x *int, y *string) {
	*x = *x + 5
	*y = *y + "World!"
	return
}

func main() {
	//Accepting return value in variable
	sum := add(20, 30)
	fmt.Println("Sum:", sum)
	//Accepting a named return value
	area := getArea(4, 7)
	fmt.Println("Area:", area)

	//Accepting multiple return value
	area, paramteres := rectangle(10, 10)
	fmt.Println("Area:", area)
	fmt.Println("Paramteres:", paramteres)

	var number = 20
	var text = "Hey"
	fmt.Println("Before:", number, text)
	addValue(&number, &text)
	fmt.Println("After:", &number, &text)

}
