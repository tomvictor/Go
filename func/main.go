package main

import (
	"fmt"
	"fun/helper"
)

func main() {
	fmt.Println("Func Demo")
	helper.PrintNames("vw", "bmw")
	result := helper.Sum(1, 2)
	fmt.Println(result)
	a, b := helper.Calc(1, 2)
	fmt.Println(a, b)
}
