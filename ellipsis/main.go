// Golang program to show
// how to use Ellipsis (…)
package main

import "fmt"

func main() {
	sayHello()
	sayHello("Rahul")
	sayHello("Mohit", "Rahul", "Rohit", "Johny")
}

// using Ellipsis
func sayHello(names ...string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}
