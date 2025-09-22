package repositories

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/goesbams/mini-books-library/backend/entities"
	"github.com/jmoiron/sqlx"
)

type BookRepositoryInterface interface {
	GetBooks(db *sqlx.DB) ([]entities.Book, error)
	AddBook(db *sqlx.DB, book *entities.Book) error
	GetBookById(db *sqlx.DB, id string) (entities.Book, error)
	UpdateBook(db *sqlx.DB, id string, book *entities.Book) error
	DeleteBook(db *sqlx.DB, id string) error
}

type BookRepository struct{}

func NewBookRepository() BookRepositoryInterface {
	return &BookRepository{}
}

func (r *BookRepository) GetBooks(db *sqlx.DB) ([]entities.Book, error) {
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

func (r *BookRepository) AddBook(db *sqlx.DB, book *entities.Book) error {
	_, err := db.NamedExec(`
    INSERT INTO books (
        title, author, cover_image_url, description, publication_date, number_of_pages, isbn
    ) VALUES (:title, :author, :cover_image_url, :description, :publication_date, :number_of_pages, :isbn)
		`, book)
	if err != nil {
		return fmt.Errorf("database error: %w", err)
	}

	return nil
}

func (r *BookRepository) GetBookById(db *sqlx.DB, id string) (entities.Book, error) {
	var book entities.Book
	err := db.Get(&book, `
		SELECT id, title, author, cover_image_url, description, publication_date, number_of_pages, isbn
		FROM books
		WHERE id = $1
	`, id)
	if err != nil {
		return entities.Book{}, fmt.Errorf("database error: %w", err)
	}

	return book, nil
}

func (r *BookRepository) UpdateBook(db *sqlx.DB, id string, book *entities.Book) error {
	book.ID, _ = strconv.Atoi(id)

	updates := []string{}
	args := map[string]interface{}{
		"id": book.ID,
	}

	if book.Title != "" {
		updates = append(updates, "title = :title")
		args["title"] = book.Title
	}
	if book.Author != "" {
		updates = append(updates, "author = :author")
		args["author"] = book.Author
	}
	if book.CoverImageUrl != "" {
		updates = append(updates, "cover_image_url = :cover_image_url")
		args["cover_image_url"] = book.CoverImageUrl
	}
	if book.Description != "" {
		updates = append(updates, "description = :description")
		args["description"] = book.Description
	}
	if book.PublicationDate != "" {
		updates = append(updates, "publication_date = :publication_date")
		args["publication_date"] = book.PublicationDate
	}
	if book.NumberOfPages != 0 {
		updates = append(updates, "number_of_pages = :number_of_pages")
		args["number_of_pages"] = book.NumberOfPages
	}
	if book.Isbn != "" {
		updates = append(updates, "isbn = :isbn")
		args["isbn"] = book.Isbn
	}

	// no updates provided
	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf("UPDATE books SET %s WHERE id = :id", strings.Join(updates, ", "))

	res, err := db.NamedExec(query, args)
	if err != nil {
		return fmt.Errorf("database error: %w", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *BookRepository) DeleteBook(db *sqlx.DB, id string) error {
	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("database error: %w", err)
	}
	return nil
}
