package application

import (
	"log"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/auth"
	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/server"
)

func Start(config *config.Config) {
	app := server.NewServer(config)

	http.HandleFunc("/", auth.HandleMain)
	http.HandleFunc("/login", auth.HandleLogin)
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		auth.HandleCallback(w, r, app.DB)
	})

	err := app.Start(config.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
