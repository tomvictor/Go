package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"main/schema/github.com/tomvictor/protobuf/schema"
	"os"
)

func main() {

	//Create Proto Object
	p := schema.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*schema.Person_PhoneNumber{
			{Number: "bv", Type: schema.Person_MOBILE},
		},
	}

	//Print as string
	fmt.Println(p.String())

	//Write to disk
	out, err := proto.Marshal(&p)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := os.WriteFile("proto.data", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
