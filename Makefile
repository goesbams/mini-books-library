.PHONY: help migrate-up migrate-down docker-build docker-up docker-down start-up

# default service value
service ?= backend

help:
	@echo "------------------------------------------------------------------------------------------------------------"
	@echo "Commands:"
	@echo "------------------------------------------------------------------------------------------------------------"
	@echo -e "    migrate-up: \t\t Apply database migrations (e.g., add new tables/columns)"
	@echo -e "    migrate-down: \t\t Revert database migrations (e.g., remove tables/columns)"
	@echo -e "    docker-build service=<service_name>: \t Build the specified service (e.g., backend, postgres)"
	@echo -e "    docker-up service=<service_name>: \t Start the specified service (e.g., backend, postgres) and its dependencies"
	@echo -e "    docker-down: \t\t Stop and remove all running containers and networks"
	@echo "------------------------------------------------------------------------------------------------------------"

# migrate up from local to docker
migrate-up:
	@echo "Running migrations to update the database..."
	migrate -path ./backend/migrations -database postgres://user:password@localhost:5432/books_db?sslmode=disable up

migrate-down:
	@echo "Reverting migrations to rollback the database..."
	migrate -path ./backend/migrations -database postgres://user:password@localhost:5432/books_db?sslmode=disable down

docker-build:
	@echo "Building the $(service) service image..."
	docker-compose build --no-cache $(service)

docker-up:
	@echo "Starting up the $(service) service..."
	docker-compose up $(service)

docker-down:
	@echo "Stopping and removing all running containers and networks..."
	docker-compose down

swag-up:
	@echo "Generating swagger document"
	swag init -g backend/cmd/main.go