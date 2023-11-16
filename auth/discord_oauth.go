package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/repository"
	"github.com/ericoliveiras/alert-bot-go/request"
	"github.com/ericoliveiras/alert-bot-go/service"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
)

var cfg = config.NewConfig()

func GetDiscordOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.Discord.ClientID,
		ClientSecret: cfg.Discord.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  cfg.Discord.AuthURL,
			TokenURL: cfg.Discord.TokenURL,
		},
		RedirectURL: cfg.Discord.RedirectURL,
		Scopes:      cfg.Discord.Scopes,
	}
}

var discordOauthConfig = GetDiscordOAuthConfig()
var oauthStateString = "random"

func HandleMain(w http.ResponseWriter, r *http.Request) {
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

	SaveCookie(w, token)

	client := discordOauthConfig.Client(r.Context(), token)
	resp, err := client.Get("https://discord.com/api/users/@me")
	if err != nil {
		log.Printf("Error making request to get user information: %s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	var user request.CreateUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding JSON: %s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	createUser(w, r, db, &user)
}

func createUser(w http.ResponseWriter, r *http.Request, db *sqlx.DB, user *request.CreateUser) {
	userService := service.NewUserService(repository.NewUserRepository(db))
	err := userService.Create(context.Background(), user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
}
