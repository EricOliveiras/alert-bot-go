package migrations

import (
	"fmt"

	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(config *config.Config) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		panic(fmt.Errorf("failed to create migrate instance: %w", err))
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Errorf("failure to apply migrations: %w", err))
	}
}
