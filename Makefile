.PHONY: help migrate-up migrate-down docker-build docker-up docker-down swag-up seed

# Default service value for docker commands (override with `make docker-up service=frontend`)
service ?= backend

help:
	@echo "------------------------------------------------------------------------------------------------------------"
	@echo "Commands:"
	@echo "------------------------------------------------------------------------------------------------------------"
	@echo "  help                          Show available commands on Makefile"
	@echo "  migrate-up                    Apply database migrations (e.g., add new tables/columns)"
	@echo "  migrate-down                  Revert database migrations (e.g., remove tables/columns)"
	@echo "  docker-build service=<name>   Build the specified service (e.g., backend, postgres)"
	@echo "  docker-up service=<name>      Start the specified service (e.g., backend, postgres) and its dependencies"
	@echo "  docker-down                   Stop and remove all running containers and networks"
	@echo "  swag-up                       Generate Swagger API documentation"
	@echo "  seed                          Seed books data into the database"
	@echo "------------------------------------------------------------------------------------------------------------"

# Migration commands
migrate-up:
	@echo "Running migrations to update the database..."
	migrate -path ./backend/migrations -database "postgres://user:password@localhost:5432/books_db?sslmode=disable" up

migrate-down:
	@echo "Reverting migrations to rollback the database..."
	migrate -path ./backend/migrations -database "postgres://user:password@localhost:5432/books_db?sslmode=disable" down

# Docker commands
docker-build:
	@echo "Building the $(service) service image..."
	docker-compose build --no-cache $(service)

docker-up:
	@echo "Starting the $(service) service..."
	docker-compose up $(service)

docker-down-all:
	@echo "Stopping and removing all running containers and networks..."
	docker-compose down

docker-down:
	@echo "Stopping and removing service"
	docker-compose stop $(service)
	docker-compose rm $(service)

# Swagger docs
swag-up:
	@echo "Generating Swagger documentation..."
	swag init -g backend/cmd/main.go

# Seed database
seed:
	@echo "Seeding initial book data into the database..."
	cd backend && go run seeds/main.go

web-start:
	@echo " Building frontend image and starting frontend container"
	docker build -t mini-books-frontend ./frontend
	docker run --rm -p 3000:3000 --name mini-frontend mini-books-frontend

web-stop:
	@echo "Stopping FE instance..."
	docker stop mini-frontend || true
	docker rm mini-frontend || true
