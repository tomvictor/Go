package main

import (
	"errors"
	"fmt"
)

type MyType struct {
	value string
}

func main() {
	fmt.Println("Return pointer")
	// Return Type

	// Calling function will never fail
	ans, err := withoutErr()
	fmt.Println(ans, err)
	// Calling a function can fail
	ans, err = withErr()
	fmt.Println(ans, err)

	// Return Pointer

	// Calling function will never fail
	res, err := withoutErrP()
	fmt.Println(res, err)
	// Calling a function can fail
	res, err = withErrP()
	fmt.Println(res, err)

}

func withoutErr() (MyType, error) {
	return MyType{"India"}, nil
}

func withErr() (MyType, error) {
	return MyType{}, errors.New("error occured")
}

func withoutErrP() (*MyType, error) {
	return &MyType{"India"}, nil
}

func withErrP() (*MyType, error) {
	return nil, errors.New("error occured")
}
