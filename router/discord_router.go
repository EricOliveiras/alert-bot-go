package router

import (
	"log"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/auth"
	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/controller"
	"github.com/ericoliveiras/alert-bot-go/handler"
	"github.com/ericoliveiras/alert-bot-go/middleware"
	"github.com/jmoiron/sqlx"
)

var cfg = config.NewConfig()

var discordOauthConfig = handler.GetDiscordOAuthConfig()
var oauthStateString = "random"

func HandleMain(w http.ResponseWriter, r *http.Request) {
	if middleware.IsAuthenticated(r) {
		http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
		return
	}

	var html = `<html><body><a href="/login">Login with Discord!</a></body></html>`
	_, err := w.Write([]byte(html))
	if err != nil {
		log.Fatalf("Error writing HTML response: %s", err.Error())
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := discordOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallback(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	state := r.FormValue("state")
	if state != oauthStateString {
		log.Printf("Invalid state!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := discordOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		log.Printf("Error exchanging code for token: %s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	auth.SaveCookie(w, token)

	resp, err := handler.GetInfo(w, r, token, cfg.Discord.GetUserInfoUrl)
	if err != nil {
		log.Printf("Error making request to get user information: %s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	defer resp.Body.Close()

	userController := controller.NewUserController(db)

	userController.Create(w, r, resp)
}
