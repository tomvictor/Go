package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"main/schema/github.com/tomvictor/protobuf/schema"
	"os"
)

const data_file_name = "proto.data"

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

	//Write buffer to disk
	out, err := proto.Marshal(&p)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := os.WriteFile(data_file_name, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
	fmt.Printf("Write success: %v\n", p)

	// Read the existing user data.
	in, err := os.ReadFile(data_file_name)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	user := &schema.Person{}
	if err := proto.Unmarshal(in, user); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Printf("Read success: %v", user)
}
