package main

import "fmt"

const (
	PUBLISHED = iota
	DRAFT
	TRASH
)

func main() {
	// Int
	var i int
	i = 42
	fmt.Println(i)

	// float32
	var f float32 = 3.14
	fmt.Println(f)

	// string
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

	// Pointers

	ptr := &lastname
	fmt.Println(ptr, *ptr)

	lastname = "Victor"
	fmt.Println(ptr, *ptr)

	const pi float32 = 3.14
	fmt.Println(pi)
	fmt.Println(pi + 1.2)

	fmt.Println(PUBLISHED, DRAFT, TRASH)

	// Array

	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr2)

	// Slice
	slice := []int{1, 2, 0}

	fmt.Println(slice)
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice)

	// Map
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m)

	// Struct
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Tom"
	u.LastName = "Victor"

	fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user{
		ID:        2,
		FirstName: "Don",
		LastName:  "Victor",
	}

	fmt.Println(u2)
}
