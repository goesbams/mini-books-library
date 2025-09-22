package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/goesbams/mini-books-library/backend/entities"
	"github.com/goesbams/mini-books-library/backend/repositories"
	"github.com/goesbams/mini-books-library/backend/utils"
	"github.com/jmoiron/sqlx"
)

type BookService interface {
	GetBooks() ([]entities.Book, error)
	AddBook(*entities.Book) error
	GetBookById(id string) (entities.Book, error)
	UpdateBook(id string, book *entities.Book) error
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

func (s *BookServiceSqlx) UpdateBook(id string, book *entities.Book) error {
	validate := validator.New()

	var validationErrors []utils.FieldError

	if book.Title != "" {
		if err := validate.Var(book.Title, "min=2,max=255"); err != nil {
			validationErrors = append(validationErrors, utils.FieldError{Field: "title", Rule: "min=2,max=255"})
		}
	}

	if book.Author != "" {
		if err := validate.Var(book.Author, "min=2,max=255"); err != nil {
			validationErrors = append(validationErrors, utils.FieldError{Field: "author", Rule: "min=2,max=255"})
		}
	}

	if book.CoverImageUrl != "" {
		if err := validate.Var(book.CoverImageUrl, "omitempty,url"); err != nil {
			validationErrors = append(validationErrors, utils.FieldError{Field: "cover_image_url", Rule: "url"})
		}
	}

	if book.Description != "" {
		if err := validate.Var(book.Description, "max=1000"); err != nil {
			validationErrors = append(validationErrors, utils.FieldError{Field: "description", Rule: "max=1000"})
		}
	}

	if book.PublicationDate != "" {
		if err := validate.Var(book.PublicationDate, "datetime=2006-01-02"); err != nil {
			validationErrors = append(validationErrors, utils.FieldError{Field: "publication_date", Rule: "datetime"})
		}
	}

	if book.NumberOfPages != 0 {
		if err := validate.Var(book.NumberOfPages, "gt=0"); err != nil {
			validationErrors = append(validationErrors, utils.FieldError{Field: "number_of_pages", Rule: "gt=0"})
		}
	}

	if book.Isbn != "" {
		if err := validate.Var(book.Isbn, "len=13,numeric"); err != nil {
			validationErrors = append(validationErrors, utils.FieldError{Field: "isbn", Rule: "len=13,numeric"})
		}
	}

	if book.Title == "" && book.Author == "" && book.CoverImageUrl == "" &&
		book.Description == "" && book.PublicationDate == "" &&
		book.NumberOfPages == 0 && book.Isbn == "" {
		return utils.ValidationError{
			Errors: []utils.FieldError{
				{Field: "body", Rule: "at_least_one_field_required"},
			},
		}
	}

	if len(validationErrors) > 0 {
		return utils.ValidationError{Errors: validationErrors}
	}

	return s.repo.UpdateBook(s.db, id, book)
}
