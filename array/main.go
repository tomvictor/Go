package main

import "fmt"

func main() {
	fmt.Println("Array samples in golang")

	// Fixed size collection of item
	// can't change the size after the declaraion

	// Eg 1

	var arr1 [3]string // initial value will be empty string

	arr1 = [3]string{"Coffee", "Tea", "Other"}
	fmt.Println("arr1: ", arr1)

	// // Eg 2

	arr2 := [3]string{"car", "Bus", "Tram"}
	fmt.Println("arr2: ", arr2)

	// // Eg 3
	arr3 := arr2 // this is a copy operation
	fmt.Println("arr3: ", arr3)

	arr2[1] = "Train"
	fmt.Println("arr2: ", arr2)
	fmt.Println("arr3: ", arr3)

	// integer array demo

	var intArray [3]int // declaration
	fmt.Println("intArray: ", intArray)

	// // Assign values

	intArray = [3]int{10, 1, 2}
	fmt.Println("intArray: ", intArray)
	fmt.Println("intArray: ", intArray[0])

	intArray[2] = 100
	fmt.Println("intArray: ", intArray)

	// // find length of array

	fmt.Println("length of intArray: ", len(intArray))
}
