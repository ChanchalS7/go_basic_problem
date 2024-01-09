package main

import "fmt"

func main() {
	//basic while loop
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	//do while loop
	num := 0
	for {
		//work
		fmt.Println(num)
		if num == 10 {
			break
		}
		num++
	}
}
