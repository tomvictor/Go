package main

import "fmt"

func main() {
	fmt.Println("Golang panic and recover demo")
	x, y := 10, 2
	fmt.Printf("%v divide by %v is %v\n", x, y, divide(x, y))
	x, y = 10, 0
	fmt.Printf("%v divide by %v is %v\n", x, y, divide(x, y))
}

func divide(a, b int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return a / b
}
