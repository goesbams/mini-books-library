package repositories

import (
	"github.com/goesbams/mini-books-library/backend/entities"
	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
	GetBooks(db *sqlx.DB) ([]entities.Book, error)
	AddBook(db *sqlx.DB, book *entities.Book) error
	GetBookById(db *sqlx.DB, id int) (entities.Book, error)
	UpdateBook(db *sqlx.DB, id int, book *entities.Book) error
	DeleteBook(db *sqlx.DB, in int)
}

type BookRepositorySqlx struct{}
