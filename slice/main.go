package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println("golang slice sample")

	var s []string
	fmt.Println(s)
	s = []string{"Chai", "Tea", "Cofee"}
	fmt.Println(s)

	v := s
	fmt.Println(v)
	v[1] = "Chai"
	fmt.Println(s, v) //[Chai Chai Cofee] [Chai Chai Cofee]

	// append
	s = append(s, "Hot-Chocolate", "Hot-Milk")
	fmt.Println(s)
	// Deletion
	slices.Delete(s, 1, 2)
	fmt.Println(s)

}
