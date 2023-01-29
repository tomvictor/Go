package helper

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func Calc(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func PrintNames(names ...string) {
	fmt.Println(names[0])
	fmt.Println(names[1])
}
