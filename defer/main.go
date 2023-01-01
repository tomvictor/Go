package main

import (
	"database/sql"
	"fmt"
)

// Output: in the reverse order/Last in First Out
// Main 1
// Main 2
// Defer 2
// Defer 1
func defer_demo() {
	fmt.Println("Main 1")
	defer fmt.Println("Defer 1")
	fmt.Println("Main 2")
	defer fmt.Println("Defer 2")

}

// More useful Example of defer statement in golang
func database_demo() {
	db, _ := sql.Open("driver", "connection string")
	defer db.Close()

	rows, _ := db.Query("some valid sql statement")
	defer rows.Close()
}

func main() {
	defer_demo()
	// database_demo()
}
