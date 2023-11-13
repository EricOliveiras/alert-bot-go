package application

import (
	"log"

	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/server"
)

func Start(config *config.Config) {
	app := server.NewServer(config)

	err := app.Start(config.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
