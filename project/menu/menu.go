package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

type menu []menuItem

func (m menu) print() { //this is a method for type menu
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price) //%10s is formatting verb
		}
	}
}

func (m *menu) add() error {
	fmt.Println("Please enter a name of new item:")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)

	for _, item := range data {
		if item.name == name {
			return errors.New("menu item alrady exists")
		}
	}

	*m = append(*m, menuItem{name: name, prices: make(map[string]float64)})
	return nil //nil indicates no error was encountered
}

var in = bufio.NewReader(os.Stdin)

func Add() error {
	return data.add()
}

func Print() {
	data.print()
}
