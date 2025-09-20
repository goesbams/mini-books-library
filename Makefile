.PHONY: help migrate-up migrate-down docker-up-db docker-up-backend

help:
	@echo "Commands:"
	@echo "------------------------------------"
	@echo "  docker-up-db:      Compose up PostgreSQL database"
	@echo "  docker-up-backend: Compose up backend service"
	@echo "  migrate-up:        Apply migrations (add new table/columns)"
	@echo "  migrate-down:      Revert migrations (remove table/columns)"
	@echo "  docker-up:         Build all images"
	@echo "  docker-down:       Compose down all services"
	@echo "------------------------------------"

migrate-up:
	migrate -path ./backend/migrations -database postgres://user:password@localhost:5432/books_db?sslmode=disable up

migrate-down:
	migrate -path ./backend/migrations -database postgres://user:password@localhost:5432/books_db?sslmode=disable down

docker-up-db:
	docker-compose up postgres

docker-up-backend:
	docker-compose up backend

docker-down:
	docker-compose down