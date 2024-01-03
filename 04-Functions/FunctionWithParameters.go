package main

import "fmt"

func main() {
	hello("golang")
	add(2, 3)

}
func hello(x string) {
	fmt.Printf("Hello %s\n", x)
}
func add(x int, y int) {
	fmt.Println(x + y)
}
