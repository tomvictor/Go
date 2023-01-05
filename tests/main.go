package main

import "fmt"

func main() {
	fmt.Println("Golang Testing demo code")
	a, b := 10, 5
	sum := Adder(a, b)
	fmt.Printf("a + b = %v\n", sum)
}

func Adder(a, b int) int {
	return a + b
}

// go test . 	 - Test module
// go test ./... - Test full
