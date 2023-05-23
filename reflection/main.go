package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name" customtag:"custom value"`
	Age  int    `json:"age" customtag:"another value"`
}

func main() {
	p := Person{Name: "John", Age: 30}

	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s, JSON Tag: %s, Custom Tag: %s\n", field.Name, field.Tag.Get("json"), field.Tag.Get("customtag"))
	}
}
