package main

import "fmt"

func main() {
	fmt.Println("Infinite")
	i := 1
	for { //this is infinite
		i++
		fmt.Println(i)
		break
	}

	fmt.Println("While")
	i = 5
	for i < 10 {
		i++
		fmt.Println(i)
	}

	fmt.Println("For loop")
	for i := 50; i < 60; i++ {
		fmt.Println(i)
	}

	//what is i now?
	fmt.Println("I now is", i)
}
