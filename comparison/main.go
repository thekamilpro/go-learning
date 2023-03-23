package main

import "fmt"

func main() {
	a, b := 5, 10
	c := a == b
	fmt.Println(c)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(float32(a) / float32(b))
}
