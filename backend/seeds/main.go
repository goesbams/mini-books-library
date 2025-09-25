package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func SeedBooks(db *sql.DB) {
	_, err := db.Exec(`
		INSERT INTO books (title, author, cover_image_url, description, publication_date, number_of_pages, isbn) VALUES
		('Clean Code: A Handbook of Agile Software Craftsmanship',
		 'Robert C. Martin',
		 'https://www.oreilly.com/covers/urn:orm:book:9780136083238/400w/',
		 'A handbook of agile software craftsmanship focusing on best practices for writing clean, maintainable code.',
		 '2008-08-01',
		 464,
		 '9780136083238'),

		('The Pragmatic Programmer: Your Journey to Mastery',
		 'Andrew Hunt, David Thomas',
		 'https://covers.openlibrary.org/b/isbn/9780201616224-L.jpg',
		 'A classic book on software development and pragmatic thinking.',
		 '1999-10-30',
		 352,
		 '9780201616224'),

		('Design Patterns: Elements of Reusable Object-Oriented Software',
		 'Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides',
		 'https://covers.openlibrary.org/b/isbn/9780201633610-L.jpg',
		 'The famous “Gang of Four” book introducing design patterns in software engineering.',
		 '1994-10-21',
		 395,
		 '9780201633610'),

		('Refactoring: Improving the Design of Existing Code',
		 'Martin Fowler',
		 'https://covers.openlibrary.org/b/isbn/9780201485677-L.jpg',
		 'Guidance on how to refactor code to improve readability and maintainability.',
		 '1999-07-08',
		 431,
		 '9780201485677'),

		('Working Effectively with Legacy Code',
		 'Michael Feathers',
		 'https://covers.openlibrary.org/b/isbn/9780131177055-L.jpg',
		 'A practical guide to improving and modifying legacy codebases safely.',
		 '2004-09-22',
		 456,
		 '9780131177055'),

		('Domain-Driven Design: Tackling Complexity in the Heart of Software',
		 'Eric Evans',
		 'https://covers.openlibrary.org/b/isbn/9780321125217-L.jpg',
		 'The foundational book on domain-driven design (DDD).',
		 '2003-08-30',
		 560,
		 '9780321125217'),

		('Continuous Delivery: Reliable Software Releases through Build, Test, and Deployment Automation',
		 'Jez Humble, David Farley',
		 'https://covers.openlibrary.org/b/isbn/9780321601919-L.jpg',
		 'Best practices for continuous integration and continuous delivery (CI/CD).',
		 '2010-07-27',
		 512,
		 '9780321601919'),

		('The Mythical Man-Month: Essays on Software Engineering',
		 'Frederick P. Brooks Jr.',
		 'https://covers.openlibrary.org/b/isbn/9780201835953-L.jpg',
		 'Classic essays on software project management and engineering.',
		 '1995-08-12',
		 336,
		 '9780201835953'),

		('Patterns of Enterprise Application Architecture',
		 'Martin Fowler',
		 'https://covers.openlibrary.org/b/isbn/9780321127426-L.jpg',
		 'Comprehensive catalog of enterprise application architecture patterns.',
		 '2002-11-15',
		 533,
		 '9780321127426'),

		('Introduction to Algorithms',
		 'Thomas H. Cormen, Charles E. Leiserson, Ronald L. Rivest, Clifford Stein',
		 'https://covers.openlibrary.org/b/isbn/9780262033848-L.jpg',
		 'Widely used reference on algorithms (CLRS).',
		 '2009-07-31',
		 1312,
		 '9780262033848');
	`)
	if err != nil {
		log.Fatalf("❌ Failed to seed books: %v", err)
	}
	log.Println("✅ Seeded 10 books")
}

func main() {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "user"
	dbPass := "password"
	dbName := "books_db"

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect DB: %v", err)
	}
	defer db.Close()

	SeedBooks(db)
}
