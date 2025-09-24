package main

import (
	"github.com/goesbams/mini-books-library/backend/config"
	"github.com/goesbams/mini-books-library/backend/database"
	_ "github.com/goesbams/mini-books-library/backend/docs"
	"github.com/goesbams/mini-books-library/backend/middleware"
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
	urlService := services.NewUrlService()

	// create echo instance
	e := echo.New()

	// CORS middleware
	e.Use(middleware.CORS())

	// define handlers
	bookHandler := handlers.NewBookHandler(bookService)
	urlHandler := handlers.NewUrlHandler(urlService)

	// Routes
	e.GET("/books", bookHandler.GetBooks)
	e.POST("books", bookHandler.AddBook)
	e.GET("books/:id", bookHandler.GetBookById)
	e.PUT("books/:id", bookHandler.UpdateBook)
	e.DELETE("books/:id", bookHandler.DeleteBook)

	e.POST("/urls/process", urlHandler.ProcessUrl)

	// Swagger UI route
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	// start server
	e.Logger.Fatal(e.Start(":9000"))
}
