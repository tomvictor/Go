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

var in = bufio.NewReader(os.Stdin)

func Print() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for key, value := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", key, value)
		}
	}
}

func AddItem() {
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
}

func PrintInfo() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	fmt.Println("2) Add item")
	fmt.Println("q) quit")
}
