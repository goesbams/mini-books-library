package database

import (
	"context"
	"fmt"
	"log"

	"github.com/goesbams/mini-books-library/backend/config"
	"github.com/jackc/pgx/v5"
)

func ConnectDB(cfg *config.Config) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Dbname,
		cfg.Database.Sslmode,
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("unable to connect to database:", err)
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("failed to ping database")
		return nil, err
	}

	log.Println("successfully connected to the database")
	return conn, nil
}
