package main

import "fmt"

func first() {
	fmt.Println("First function")
}
func second() {
	fmt.Println("Second function")
}
func main() {
	defer second()
	first()
}
