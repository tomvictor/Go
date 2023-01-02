package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Generic Programming")
	fmt.Print(strings.Repeat("-", 20))
	fmt.Printf("\n")
	src := map[string]float32{
		"Tom":    200,
		"Don":    201,
		"Jessie": 202,
	}
	fmt.Println(src)
	cln := GenCopy(src)
	fmt.Println(cln)
}

func GenCopy[K comparable, V any](src map[K]V) map[K]V {
	clone := make(map[K]V, len(src))
	for key, value := range src {
		clone[key] = value
	}
	return clone
}
