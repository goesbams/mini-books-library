package repositories

import (
	"fmt"

	"github.com/goesbams/mini-books-library/backend/entities"
	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
	GetBooks(db *sqlx.DB) ([]entities.Book, error)
	AddBook(db *sqlx.DB, book *entities.Book) error
}

type BookRepositorySqlx struct{}

func NewBookRepository() BookRepository {
	return &BookRepositorySqlx{}
}

func (r *BookRepositorySqlx) GetBooks(db *sqlx.DB) ([]entities.Book, error) {
	var books []entities.Book
	err := db.Select(&books, "SELECT id, title, author, cover_image_url, description, publication_date, Isbn, number_of_pages FROM books")
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	if len(books) == 0 {
		return []entities.Book{}, nil
	}

	return books, nil
}

func (r *BookRepositorySqlx) AddBook(db *sqlx.DB, book *entities.Book) error {
	_, err := db.NamedExec(`
    INSERT INTO books (
        title, author, cover_image_url, description, publication_date, number_of_pages, isbn
    ) VALUES (:title, :author, :cover_image_url, :description, :publication_date, :number_of_pages, :isbn)
		`, book)
	if err != nil {
		return fmt.Errorf("failed to add book: %w", err)
	}

	return nil
}
