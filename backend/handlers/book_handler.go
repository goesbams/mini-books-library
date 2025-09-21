package handlers

import (
	"net/http"

	"github.com/goesbams/mini-books-library/backend/services"
	"github.com/labstack/echo/v4"
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
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "internal server server",
			"message": "unable to fetch books",
		})
	}

	return c.JSON(http.StatusOK, books)
}
