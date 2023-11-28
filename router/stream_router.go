package router

import (
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/handler"
	"github.com/jmoiron/sqlx"
)

func SetupStreamRoutes(db *sqlx.DB) {
	streamHandler := handler.NewStreamHandler(db)
	http.HandleFunc("/streams/create", streamHandler.CreateStream)
}
