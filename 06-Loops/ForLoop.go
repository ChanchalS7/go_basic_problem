package main

import "fmt"

func main() {
	//Basic for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//infinite loop
	i := 0
	for {
		fmt.Println("Hello world!")
		//breaks/stops the infinite loop
		if i == 10 {
			break
		}
		i++
	}
	//ranges
	strings := []string{"Hello", "World", "!"}
	//get the index and value while looping through the range
	for i, val := range strings {
		fmt.Printf("%d: %s \n", i, val)
	}

	// //only getting the index
	for i := range strings {
		fmt.Println(i)

	}
	for _, val := range strings {
		fmt.Println(val)
	}
}
