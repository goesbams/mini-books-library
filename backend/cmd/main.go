package main

import (
	"fmt"
	"log"

	"github.com/goesbams/mini-books-library/backend/database"
)

func main() {
	// initialize database connection
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect the database:", err)
	}

	fmt.Println(db)
}
