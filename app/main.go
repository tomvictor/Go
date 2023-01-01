package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
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
loop:

	for {

		in := bufio.NewReader(os.Stdin)

		fmt.Println("Please select an option")
		fmt.Println("1) Display menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for key, value := range item.prices {
					fmt.Printf("\t%10s%10.2f\n", key, value)
				}
			}
		case "2":
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}

}
