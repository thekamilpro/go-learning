package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"Coffee", "Espresso", "Cappucino"}
	fmt.Println(arr)

	fmt.Println(arr[1])

	arr[1] = "Chai tea"
	fmt.Println(arr)

	arr2 := arr //go copies by value

	arr2[2] = "Chai latte"
	fmt.Println(arr, arr2)
}
