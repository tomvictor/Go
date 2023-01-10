package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println("Array samples in golang")

	s := []string{"a", "b", "c"}
	fmt.Println(s)
	s = append(s, "d", "e")
	fmt.Println(s)

	s = slices.Delete(s, 1, 2)
	fmt.Println(s)

}
