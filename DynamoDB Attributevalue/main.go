package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

const data = `
	{
		"id": 1,
		"name": "Tom"
	}
`

func main() {

	var newUser User

	// load data
	err := json.Unmarshal([]byte(data), &newUser)
	if err != nil {
		log.Fatal("unable to unmarshal data")
	}

	fmt.Printf("New User: %+v\n", newUser)
	fmt.Println("Marshal to attribute value")
	av, err := attributevalue.MarshalMap(newUser)
	if err != nil {
		log.Fatalf("unable struct to marshal attribute value")
	}
	fmt.Printf("%s\n", av["Id"])
	fmt.Printf("%s\n", av["Name"])
}
