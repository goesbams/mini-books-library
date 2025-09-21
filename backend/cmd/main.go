package main

import (
	"github.com/goesbams/mini-books-library/backend/config"
	"github.com/goesbams/mini-books-library/backend/database"
	_ "github.com/goesbams/mini-books-library/backend/docs"
	"github.com/goesbams/mini-books-library/backend/utils"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/goesbams/mini-books-library/backend/handlers"
	"github.com/goesbams/mini-books-library/backend/repositories"
	"github.com/goesbams/mini-books-library/backend/services"
	"github.com/labstack/echo/v4"
)

// @title Mini Books Library API
// @version 1.0
// @description This is a sample API for managing books in the library.
// @contact.name Bambang Handoko (Ando)
// @contact.email bambang.handoko12@gmail.com

func main() {
	// initialize logger
	logger := utils.InitializeLogger()

	// load configuration
	cfg, err := config.LoadConfig("config/config.dev.yaml")
	if err != nil {
		logger.Fatal("error loading config:", err)
	}

	// initialize database connection
	conn, err := database.ConnectDB(cfg)
	if err != nil {
		logger.Fatal("failed to connect the database:", err)
	}

	// setup repos & services
	bookRepo := repositories.NewBookRepository()
	bookService := services.NewBookService(bookRepo, conn)

	// setup handlers
	e := echo.New()
	handler := handlers.NewHandler(bookService)

	// Routes
	e.GET("/books", handler.GetBooks)
	e.POST("books", handler.AddBook)
	e.GET("books/:id", handler.GetBookById)

	// Swagger UI route
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	// start server
	e.Logger.Fatal(e.Start(":9000"))
}
