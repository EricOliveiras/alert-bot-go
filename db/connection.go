package db

import (
	"fmt"
	"log"

	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(config *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB.Host, config.DB.User, config.DB.Password, config.DB.Name, config.DB.Port)

	db, err := sqlx.Open(config.DB.Driver, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database. %v", err)
	}

	return db, nil
}
