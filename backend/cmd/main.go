package main

import (
	"fmt"
	"log"

	"github.com/goesbams/mini-books-library/backend/config"
	"github.com/goesbams/mini-books-library/backend/database"
	"github.com/goesbams/mini-books-library/backend/migrations"
)

func main() {
	// load configuration
	cfg, err := config.LoadConfig("config/config.dev.yaml")
	if err != nil {
		log.Fatal("error loading config:", err)
	}

	// initialize database connection
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("failed to connect the database:", err)
	}

	// run migrations
	err = migrations.RunMigrations(cfg)
	if err != nil {
		log.Fatal("failed to run migrations:", err)
	}

	fmt.Println(db)
}
