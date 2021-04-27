package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	user   []*User
	nextID int32 = 1
)
