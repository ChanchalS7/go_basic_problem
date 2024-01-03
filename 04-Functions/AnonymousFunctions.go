package main

import "fmt"

//defining anonymous function
var area = func(l, b int) int {
	return l * b

}

func main() {
	area := area(2, 3)
	fmt.Println(area)
}
