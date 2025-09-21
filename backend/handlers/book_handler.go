package handlers

import (
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
