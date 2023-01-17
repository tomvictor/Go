package main

import "fmt"

func main() {
	fmt.Println("Golang working with maps")

	var m map[string][]string
	var n map[string][]any

	fmt.Println(m)
	m = map[string][]string{
		"coffee": {"Coffee", "Coffee-Late"},
		"tea":    {"Chai", "Latte", "Chai-Latte"},
	}
	fmt.Println(m)
	fmt.Println(m["coffee"])
	// append
	m["others"] = []string{"milk", "water"} // but the ordering is determined by go, we have no control

	delete(m, "tea")
	fmt.Println(m)
	// stil we can query it
	v := m["tea"]
	fmt.Println(v) // this will be an empty arr
	// because it will send default
	// for more control we can fo like this
	v, ok := m["tea"] // ok will be false
	fmt.Println(v, ok)

	// Assignment
	c := m
	fmt.Println(c) // same as m
	c["others"] = []string{"hot-milk"}
	fmt.Println(m)
	fmt.Println(c)
	// both will be same
}
