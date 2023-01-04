package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Demo: Convert Panics to Errors")
	a, b := 10, 0
	// result, err := divideNormal(a, b)
	result, err := dividePanic(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Divide %v by %v is %v\n", a, b, result)

}

// Normal Method using simple errors
func divideNormal(divident, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Divisor must not be zero")
	}

	return divident / divisor, nil
}

// Method 2 using panic recovery, note the name return values
func dividePanic(divident, divisor int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
			return
		}
	}()

	return divident / divisor, nil
}
