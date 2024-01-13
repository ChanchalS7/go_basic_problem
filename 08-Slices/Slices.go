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

	//initialize the slices with values using slice literal
	var z = []int{10, 20, 30, 40}
	fmt.Printf("z \tLen: %v \tCap: %v\n", len(z), cap(z))
	fmt.Println(z)

	//creating a slice using the new keyword
	var a = new([50]int)[2:10]
	fmt.Printf("a \tLen:%v \tCap: %v\n", len(a), cap(a))
	fmt.Println(a)

	//add items using append function
	var b = make([]int, 1, 100)
	fmt.Println("b:", b)
	b = append(b, 20)
	fmt.Println(b)

	//access slice items
	var c = []int{10, 20, 30, 40}
	fmt.Println("C", c[0])
	fmt.Println(c[0:3])

	//change items values
	var d = []int{10, 20, 30, 40}
	fmt.Println("D:", d)
	d[1] = 35
	fmt.Println(d)

	//copy slice into another slice
	var e = []int{10, 20, 30, 40}
	var f = []int{50, 60, 70, 80}

	copy(e, f)
	fmt.Println("E:", e)

	//Append a slice to an existing one
	var g = []int{10, 20, 30, 40}
	var h = []int{50, 60, 70, 80}
	g = append(g, h...)
	fmt.Println("G:", g)

}
