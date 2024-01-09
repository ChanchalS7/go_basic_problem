package main

import "fmt"

func main() {
	//if statement
	x := true
	if x {
		fmt.Println("True")
	}
	// if else
	y := 100
	if y > 80 {
		fmt.Println("Greater than 80")
	} else {
		fmt.Println("Lesser than 80")
	}
	//if-elseif statement
	grade := 5
	if grade == 1 {
		fmt.Println("You have A")
	} else if grade > 1 && grade < 5 {
		fmt.Println("You have no A but you are positive")
	} else {
		fmt.Println("Your grade is negative")
	}

	//if statement
	if a := 10; a == 10 {
		fmt.Println(a)
	}
}
