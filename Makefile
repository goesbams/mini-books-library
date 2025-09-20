.PHONY: help migrate-up migrate-down docker-up

help:
	@echo "Commands:"
	@echo "  docker-up:    Build all images"
	@echo "  migrate-up:      Apply migrations (add new table/columns)"
	@echo "  migrate-down:    Revert migrations (remove table/columns)"

migrate-up:
	migrate -path ./backend/migrations -database postgres://user:password@localhost:5432/books_db?sslmode=disable up

migrate-down:
	migrate -path ./backend/migrations -database postgres://user:password@localhost:5432/books_db?sslmode=disable down

docker-up:
	docker-compose up

