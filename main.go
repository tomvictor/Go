package main

import (
	"fmt"

	"github.com/tomvictor/try_go/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Tom",
		LastName:  "Victor",
	}
	fmt.Println(u)
}
