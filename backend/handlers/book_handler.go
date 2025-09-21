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

func (h *Handler) GetBooks(c echo.Context) error {
	books, err := h.service.GetBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, books)
	}

	return c.JSON(http.StatusOK, books)
}
