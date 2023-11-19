package application

import (
	"log"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/router"
	"github.com/ericoliveiras/alert-bot-go/server"
)

func Start(config *config.Config) {
	app := server.NewServer(config)

	http.HandleFunc("/", router.HandleMain)
	http.HandleFunc("/login", router.HandleLogin)
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		router.HandleCallback(w, r, app.DB)
	})
	http.HandleFunc("/dashboard", router.GetUser)

	err := app.Start(config.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
