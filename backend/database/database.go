package database

import (
	"fmt"
	"log"

	"github.com/goesbams/mini-books-library/backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg *config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Dbname,
		cfg.Database.Sslmode,
	)

	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("unable to connect to database:", err)
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("failed to ping database")
		return nil, err
	}

	log.Println("successfully connected to the database")
	return conn, nil
}
