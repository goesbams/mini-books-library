package repositories

import (
	"github.com/goesbams/mini-books-library/backend/entities"
	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
	GetBooks(db *sqlx.DB) ([]entities.Book, error)
}

type BookRepositorySqlx struct{}

func NewBookRepository() BookRepository {
	return &BookRepositorySqlx{}
}

func (r *BookRepositorySqlx) GetBooks(db *sqlx.DB) ([]entities.Book, error) {
	var books []entities.Book
	err := db.Select(&books, "SELECT id, title, author, cover_image_url, publication_date FROM books")
	return books, err
}
