package main

import "fmt"

func main() {
	fmt.Println("Array samples in golang")

	// Eg 1
	var arr1 [3]string
	arr1 = [3]string{"Coffee", "Tea", "Other"}
	fmt.Println(arr1)
	// Eg 2
	arr2 := [3]string{"car", "Bus", "Tram"}
	fmt.Println(arr2)
	// Eg 3
	arr3 := arr2
	fmt.Println(arr3)

	arr2[1] = "Train"
	fmt.Println(arr2)
}
