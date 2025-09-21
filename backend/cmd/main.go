package main

import (
	"fmt"
	"log"

	"github.com/go-chi/chi/middleware"
	"github.com/goesbams/mini-books-library/backend/config"
	"github.com/goesbams/mini-books-library/backend/database"
	"github.com/labstack/echo/v4"
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

	// create echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	fmt.Println(db)
}
