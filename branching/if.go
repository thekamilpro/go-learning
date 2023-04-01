package main

import "fmt"

func main() {
	i := 9
	if i < 5 {
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10")
	} else {
		fmt.Println("i is at least 10")
	}
	fmt.Println("done")

	if j := 9; j < 5 { //can initialise variable on spot
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10")
	} else {
		fmt.Println("i is at least 10")
	}
	fmt.Println("done")
}
