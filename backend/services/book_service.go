package services

import (
	"github.com/goesbams/mini-books-library/backend/entities"
	"github.com/goesbams/mini-books-library/backend/repositories"
	"github.com/jmoiron/sqlx"
)

type BookService interface {
	GetBooks() ([]entities.Book, error)
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

	if len(books) == 0 {
		return []entities.Book{}, nil
	}
	return books, nil
}
