package main

import "fmt"

type Car struct {
	Name  string
	Model string
	Year  int
}

func (c Car) print() string {
	return c.Name + "-" + c.Model
}

func main() {
	fmt.Println("OOPs demo")

	vw := Car{
		Name:  "VW",
		Model: "Vento",
		Year:  2022,
	}
	fmt.Println(vw)
	fmt.Println(vw.print())
}
