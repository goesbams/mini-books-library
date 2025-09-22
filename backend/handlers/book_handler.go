package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/goesbams/mini-books-library/backend/entities"
	"github.com/goesbams/mini-books-library/backend/services"
	"github.com/goesbams/mini-books-library/backend/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service services.BookService
}

func NewHandler(service services.BookService) *Handler {
	return &Handler{service: service}
}

// GetBooks fetches all books
// @Summary Get all books
// @Description Retrieve a list of all books in the library
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} entities.Book
// @Failure 500 {object} map[string]interface{}
// @Router /books [get]
func (h *Handler) GetBooks(c echo.Context) error {
	books, err := h.service.GetBooks()
	if err != nil {
		logrus.WithError(err).Error("failed to fetch books")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "internal server server",
			"message": "unable to fetch books",
		})
	}

	logrus.Info("fetched books successfully")
	return c.JSON(http.StatusOK, books)
}

// AddBook adds a new book
// @Summary Add a new book
// @Description Add a new book to the library
// @Tags books
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param title formData string true "Book Title"
// @Param author formData string true "Book Author"
// @Param cover_image_url formData string false "Cover Image URL"
// @Param description formData string false "Description"
// @Param publication_date formData string true "Publication Date (YYYY-MM-DD)"
// @Param number_of_pages formData int true "Number of Pages"
// @Param isbn formData string true "ISBN (13 digits)"
// @Success 201 {object} entities.Book
// @Failure 400 {object} map[string]interface{}
// @Router /books [post]
func (h *Handler) AddBook(c echo.Context) error {
	var book entities.Book
	if err := c.Bind(&book); err != nil {
		logrus.WithError(err).Error("failed to bind book data")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "bad request",
			"message": "invalid book data",
		})
	}

	if err := h.service.AddBook(&book); err != nil {
		switch e := err.(type) {
		case utils.ValidationError:
			logrus.WithError(err).Warn("validation failed")

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":   "bad_request",
				"message": e.Errors,
			})
		default:
			logrus.WithError(err).Error("failed to add book")

			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error":   "internal_server_error",
				"message": "unable to add book",
			})
		}
	}

	logrus.Info("added new book successfully", book)
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "book created successfully",
	})
}

// GetBookByID retrieves a book by its ID
// @Summary Get book by ID
// @Description Get detailed information about a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} entities.Book
// @Failure 404 {object} map[string]string "Book not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /books/{id} [get]
func (h *Handler) GetBookById(c echo.Context) error {
	id := c.Param("id")

	book, err := h.service.GetBookById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logrus.WithError(err).Warn(fmt.Sprintf("book with id:%s not found", id))
			return c.JSON(http.StatusNotFound, map[string]string{
				"error":   "not_found",
				"message": "book not found",
			})
		}

		logrus.WithError(err).Error(fmt.Sprintf("failed to fetch book by id: %s", id))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "internal_server_error",
			"message": "something went wrong while fetching book",
		})
	}

	logrus.Info(fmt.Sprintf("get book by id:%d title:%s successfully", book.ID, book.Title))
	return c.JSON(http.StatusOK, book)
}

// UpdateBook updates the information of an existing book
// @Summary Update a book by ID
// @Description Partially update a book's details by its ID (only provided fields will be updated)
// @Tags books
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path int true "Book ID"
// @Param title formData string false "Book Title"
// @Param author formData string false "Book Author"
// @Param cover_image_url formData string false "Cover Image URL"
// @Param description formData string false "Description"
// @Param publication_date formData string false "Publication Date (YYYY-MM-DD)"
// @Param number_of_pages formData int false "Number of Pages"
// @Param isbn formData string false "ISBN (13 digits)"
// @Success 200 {object} entities.Book
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /books/{id} [put]
func (h *Handler) UpdateBook(c echo.Context) error {
	id := c.Param("id")

	var book entities.Book
	if err := c.Bind(&book); err != nil {
		logrus.WithError(err).Error("failed to bind book data")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "bad_request",
			"message": "invalid book data",
		})
	}

	// call service
	if err := h.service.UpdateBook(id, &book); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error":   "not_found",
				"message": "book not found",
			})
		}

		if verr, ok := err.(utils.ValidationError); ok {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":   "bad_request",
				"message": verr.Errors,
			})
		}

		logrus.WithError(err).Error("failed to update book")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "internal_server_error",
			"message": "unable to update book",
		})
	}

	logrus.Infof("updated book id:%s title:%s successfully", id, book.Title)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "book updated successfully",
	})
}
