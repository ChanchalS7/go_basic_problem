package main

import (
	"fmt"
	"reflect"
)

func main() {
	printMultipleString("Hello", "Golang", "!")
	printMultipleVariables(1, "green", false, 1.314, []string{"foo", "bar", "baz"})
}

//passing multiple attributes using variadic function

func printMultipleString(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func printMultipleVariables(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}
