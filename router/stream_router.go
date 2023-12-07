package router

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/ericoliveiras/alert-bot-go/handler"
	"github.com/jmoiron/sqlx"
)

func SetupStreamRoutes(db *sqlx.DB, discordSesseion *discordgo.Session) {
	streamHandler := handler.NewStreamHandler(db, discordSesseion)
	http.HandleFunc("/streams/create", streamHandler.CreateStream)
}
