package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

type User struct {
    Name     string `validate:"required"`
    Age      int    `validate:"min=18"`
    Email    string `validate:"email"`
    Password string `validate:"min=8"`
}

func (u User) Validate() error {
    validate := validator.New()
    if err := validate.Struct(u); err != nil {
        return err
    }
    // Perform additional custom validation logic
    // Implement your own validation logic here
    return nil
}

func main() {
    user := User{
        Name:     "Alice",
        Age:      20,
        Email:    "alice@example.com",
        Password: "password",
    }

    if err := user.Validate(); err != nil {
        fmt.Println("Validation failed:", err)
    } else {
        fmt.Println("Validation passed!")
    }
}