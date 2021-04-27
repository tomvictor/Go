package main

import "fmt"

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Tom"
	fmt.Println(firstName)

	b := true

	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	var name *string = new(string)

	*name = "Jessie"

	fmt.Println(name)
	fmt.Println(*name)

	lastname := "Vithayathil"

	ptr := &lastname
	fmt.Println(ptr, *ptr)

	lastname = "Victor"
	fmt.Println(ptr, *ptr)
}
