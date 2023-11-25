package router

import (
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/handler"
	"github.com/jmoiron/sqlx"
)

func SetupUserRoutes(db *sqlx.DB) {
	userHandler := handler.NewUserHandler(db)
	http.HandleFunc("/dashboard", userHandler.UserInfo)
}
