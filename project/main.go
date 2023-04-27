package main

import (
	"bufio"
	"fmt"
	"os"
	"project/menu"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {

loop: //this is label, it's what allows break statement to break out of this loop
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			err := menu.Add()
			if err != nil { // true if an error occurs
				fmt.Println(fmt.Errorf("invalid input: %w", err))
			}
		case "q":
			break loop
		default:
			fmt.Println("unknown option")
		}
	}
}
