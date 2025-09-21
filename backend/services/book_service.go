package services

import (
	"github.com/goesbams/mini-books-library/backend/entities"
	"github.com/goesbams/mini-books-library/backend/repositories"
	"github.com/goesbams/mini-books-library/backend/utils"
	"github.com/jmoiron/sqlx"
)

type BookService interface {
	GetBooks() ([]entities.Book, error)
	AddBook(*entities.Book) error
	GetBookById(id string) (entities.Book, error)
}

type BookServiceSqlx struct {
	repo repositories.BookRepository
	db   *sqlx.DB
}

func NewBookService(repo repositories.BookRepository, db *sqlx.DB) BookService {
	return &BookServiceSqlx{repo: repo, db: db}
}

func (s *BookServiceSqlx) GetBooks() ([]entities.Book, error) {
	books, err := s.repo.GetBooks(s.db)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *BookServiceSqlx) AddBook(book *entities.Book) error {
	if err := book.Validate(); err != nil {
		return utils.FormatValidationError(err, book)
	}

	return s.repo.AddBook(s.db, book)
}

func (s *BookServiceSqlx) GetBookById(id string) (entities.Book, error) {
	return s.repo.GetBookById(s.db, id)
}
