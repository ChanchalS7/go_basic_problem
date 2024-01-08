package main

import "fmt"

func main() {
	l := 10
	b := 20
	//closure functions are a special case of anyonymous function where you access outside function
	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()

}
