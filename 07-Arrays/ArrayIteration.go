package main

import "fmt"

func main() {
	x := [5]int{0, 5, 10, 15, 20}
	//standard loop
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// getting indexx and value using range
	for index, element := range x {
		fmt.Println(index, "=>", element)
	}

	//only getting value using range
	for _, value := range x {
		fmt.Println(value)

	}
}
