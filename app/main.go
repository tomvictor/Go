package main

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

var menu = []menuItem{
	{
		name: "Coffe",
		prices: map[string]float64{
			"regular": 12.5,
			"large":   20,
		},
	},
	{
		name: "tea",
		prices: map[string]float64{
			"single": 10,
			"double": 15,
			"triple": 20,
		},
	},
}

var in = bufio.NewReader(os.Stdin)

func main() {

loop:

	for {
		printInfo()

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			printMenu()
		case "2":
			addItem()
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}

}

func printMenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for key, value := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", key, value)
		}
	}
}

func addItem() {
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
}

func printInfo() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	fmt.Println("2) Add item")
	fmt.Println("q) quit")
}
