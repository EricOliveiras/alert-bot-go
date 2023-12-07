package router

import (
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/handler"
	"github.com/jmoiron/sqlx"
)

func SetupDiscordRoutes(db *sqlx.DB) {
	discordHandler := handler.NewDiscordHandler(db)
	http.HandleFunc("/discord-channel/create", discordHandler.Create)
	http.HandleFunc("/discord-channel/delete", discordHandler.HandleDeleteChannel)
}
