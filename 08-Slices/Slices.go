package main

import (
	"fmt"
	"reflect"
)

func main() {
	//create empty slices
	var x []int

	fmt.Println(reflect.ValueOf(x).Kind())

	//creating slice using make function
	var y = make([]string, 10, 20)
	fmt.Printf("y \tLen: %v \tCap:%v\n", len(y), cap(y))
}
