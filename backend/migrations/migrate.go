package migrations

import (
	"fmt"
	"log"

	"github.com/goesbams/mini-books-library/backend/config"
	"github.com/golang-migrate/migrate/v4"
)

func RunMigrations(cfg *config.Config) error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Dbname,
		cfg.Database.Sslmode,
	)

	m, err := migrate.New("file://backend/migrations", connStr)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		log.Fatalf("failed to apply migrations: %v", err)
		return err
	}

	log.Println("migrations applied successfully")
	return nil
}
