package main

import (
	"fmt"
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
	fmt.Println(menu)
}
