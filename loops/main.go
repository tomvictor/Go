package main

import "fmt"

func main() {
	letters := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}

	for _, v := range letters {
		fmt.Println(v)
	}
}
