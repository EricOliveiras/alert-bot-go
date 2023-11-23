package application

import (
	"log"

	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/router"
	"github.com/ericoliveiras/alert-bot-go/server"
)

func Start(config *config.Config) {
	app := server.NewServer(config)

	router.SetupOauthRoutes(app.DB)
	router.SetupUserRoutes()
	router.SetupDiscordRoutes(app.DB)

	err := app.Start(config.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
