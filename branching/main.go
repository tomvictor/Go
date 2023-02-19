package main

import "fmt"

func main() {
	fmt.Println("Branching")

	i := 6

	if i < 6 {
		fmt.Println("i<6")
	} else if i > 6 {
		fmt.Println("i>6")
	} else {
		fmt.Println("Default")
	}

	switch {
	case i < 6, i < 7:
		fmt.Println("i<6")
	case i > 6:
		fmt.Println("i>6")
	default:
		fmt.Println("Default")
	}

	c := 8

	if c > 7 {
		fmt.Println("c is > 7")
		goto mylabel
	}

mylabel:
	if c > 7 {
		println("after label: c>7")
	}

}
