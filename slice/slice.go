package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	var s []string
	fmt.Println(s)

	s = []string{"Coffee", "Espresso", "Cappucino"}
	fmt.Println(s)

	s = append(s, "Hot Chocolate", "Hot Tea")
	fmt.Println(s)

	slices.Delete(s, 1, 2)
	fmt.Println(s)

	// s[1] = "Chai tea"
	// fmt.Println(s)

	// s2 := s //slices are copy by reference

	// s2[2] = "Chai latte"
	// fmt.Println(s, s2) //these point to the same object in memory
}
