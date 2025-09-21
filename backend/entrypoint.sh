#!/bin/sh

# Run migrations
echo "Running migrations..."
migrate -path /app/backend/migrations -database postgres://$DATABASE_USER:$DATABASE_PASSWORD@$DATABASE_HOST:$DATABASE_PORT/$DATABASE_NAME?sslmode=$DATABASE_SSLMODE up

# Start the backend
echo "Starting backend..."
exec /usr/local/bin/backend
