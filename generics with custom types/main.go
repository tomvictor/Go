package main

import (
	"fmt"
)

func main() {
	fmt.Println("Generics with custom types")

	value1 := []int{1, 2, 3, 4}
	sum1 := Add(value1)
	fmt.Printf("Sum of %v is %v\n", value1, sum1)

	value2 := []float32{1.1, 2.2, 3.3, 4.4}
	sum2 := Add(value2)
	fmt.Printf("Sum of %v is %v\n", value2, sum2)

	value3 := []string{"Homo", "-Sapiens"}
	sum3 := Add(value3)
	fmt.Printf("Sum of %v is %v\n", value3, sum3)
}

type Adder interface {
	int | float32 | string
}

func Add[V Adder](s []V) V {
	var result V
	for _, value := range s {
		result += value
	}
	return result
}
