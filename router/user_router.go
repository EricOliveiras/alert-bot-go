package router

import (
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/handler"
)

func SetupUserRoutes() {
	http.HandleFunc("/dashboard", handler.GetUser)
}
