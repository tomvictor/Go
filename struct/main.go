package main

import "fmt"

// var Car struct {
// 	id   int
// 	name string
// }

type Car struct {
	id   int
	name string
	// power  float32
}

func main() {
	fmt.Println("Struct")

	vw := Car{
		id:   1,
		name: "vw",
	}
	fmt.Println(vw)

	mz := Car{
		id:   2,
		name: "Mazda",
	}
	fmt.Println(mz.name)
	fmt.Println(mz)

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
