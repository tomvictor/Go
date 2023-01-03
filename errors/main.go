package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Golang errors demo")
	err := errors.New("New sample error")
	fmt.Printf("%w\n", err)

}
