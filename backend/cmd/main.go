package main

import (
	"log"

	"github.com/goesbams/mini-books-library/backend/config"
	"github.com/goesbams/mini-books-library/backend/database"
	"github.com/goesbams/mini-books-library/backend/handlers"
	"github.com/goesbams/mini-books-library/backend/repositories"
	"github.com/goesbams/mini-books-library/backend/services"
	"github.com/labstack/echo/v4"
)

func main() {
	// load configuration
	cfg, err := config.LoadConfig("config/config.dev.yaml")
	if err != nil {
		log.Fatal("error loading config:", err)
	}

	// initialize database connection
	conn, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("failed to connect the database:", err)
	}

	// setup repos & services
	bookRepo := repositories.NewBookRepository()
	bookService := services.NewBookService(bookRepo, conn)

	// setup handlers
	e := echo.New()
	handler := handlers.NewHandler(bookService)

	// Routes
	e.GET("/books", handler.GetBooks)

	// start server
	e.Logger.Fatal(e.Start(":9000"))
}
