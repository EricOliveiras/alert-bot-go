package router

import (
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/handler"
	"github.com/jmoiron/sqlx"
)

func SetupOauthRoutes(db *sqlx.DB) {
	http.HandleFunc("/", handler.HandleMain)
	http.HandleFunc("/login", handler.HandleLogin)
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleCallback(w, r, db)
	})
}
