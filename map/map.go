package main

import "fmt"

func main() {
	var m map[string][]string //hashtable of key = string, value = array of strings
	fmt.Println(m)

	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"},
	}

	fmt.Println(m)

	fmt.Println(m["coffee"])

	m["other"] = []string{"Hot Chocolate"}
	fmt.Println(m)

	delete(m, "tea")
	fmt.Println(m)

	fmt.Println(m["tea"])
	v, ok := m["tea"] //checking whether key tea exists
	fmt.Println(v, ok)

	// m2 := m maps are value by reference

}
