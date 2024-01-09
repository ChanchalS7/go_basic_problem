package main

import "fmt"

func main() {
	var i int = 10
	var k int = 20
	var z int = 30
	//arithmetic operation
	fmt.Printf("i+k=%d\n", i+k)
	fmt.Printf("i-k=%d\n", i-k)
	fmt.Printf("i*k=%d\n", i*k)
	fmt.Printf("i/k=%d\n", i/k)
	fmt.Printf("i mod k=%d\n", i%k)

	//compariso operator
	fmt.Println(i == k)
	fmt.Println(i != k)
	fmt.Println(i < k)
	fmt.Println(i > k)
	fmt.Println(i <= k)
	fmt.Println(i >= k)

	//logical operator
	fmt.Println(i < z && i > k)
	fmt.Println(i < z || i > k)
	fmt.Println(!(i == z && i > k))

	//Assignment operator
	var x, y = 15, 25
	x = y
	fmt.Println("=", x)
	x = 15
	x += y
	fmt.Println("+=", x)
	x = 50
	x -= y
	fmt.Println("-=", x)

	x = 2
	y *= x
	fmt.Println("*=", x)

	x = 100
	x /= y
	fmt.Println("/=", x)

	x = 40
	x %= y
	fmt.Println("%=", x)

}
