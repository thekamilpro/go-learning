package main

import "fmt"

func main() {
	divident, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", divident, divisor, divide(divident, divisor))

	divident, divisor = 10, 0 //this errors
	fmt.Printf("%v divided by %v is %v\n", divident, divisor, divide(divident, divisor))
}

func divide(divident, divisor int) int {
	defer func() { //anonymous function, this handles the error
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return divident / divisor
}
