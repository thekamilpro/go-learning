package main

import "fmt"

func main() {
	s := "Hello there!"
	p := &s
	fmt.Println(p)  //print memory address
	fmt.Println(*p) //dereference, show actual value

	*p = "Hello Gophers!" //change s through p
	fmt.Println(s)

	p = new(string)
	fmt.Println(p)  //new address
	fmt.Println(*p) //empty
}
