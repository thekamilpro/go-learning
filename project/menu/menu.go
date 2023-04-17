package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

type menu []menuItem

func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price) //%10s is formatting verb
		}
	}
}

func (m *menu) add() {
	fmt.Println("Please enter a name of new item:")
	name, _ := in.ReadString('\n')
	*m = append(*m, menuItem{name: name, prices: make(map[string]float64)})
}

var in = bufio.NewReader(os.Stdin)

func Add() {
	data.add()
}

func Print() {
	data.print()
}
