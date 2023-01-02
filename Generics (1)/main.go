package main

import "fmt"

func main() {
	fmt.Println("Generic Programming Demo")
	src := []float64{
		1.1,
		2.3,
		36.9,
	}
	fmt.Println(src)
	cln := normalCopy(src)
	fmt.Println(cln)

	// float32 type
	gsrc1 := []float32{
		1.1,
		2.3,
		36.9,
	}
	gcln1 := genericCopy(gsrc1)
	fmt.Println(gcln1)

	// float64 type
	gsrc2 := []float64{
		1.1,
		2.3,
		36.9,
	}
	gcln2 := genericCopy(gsrc2)
	fmt.Println(gcln2)
}

func normalCopy(src []float64) []float64 {
	clone := make([]float64, len(src))
	for index, value := range src {
		clone[index] = value
	}
	return clone
}

func genericCopy[V any](src []V) []V {
	clone := make([]V, len(src))
	for index, value := range src {
		clone[index] = value
	}
	return clone
}
