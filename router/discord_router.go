package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/auth"
	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/controller"
	"github.com/ericoliveiras/alert-bot-go/handler"
	"github.com/ericoliveiras/alert-bot-go/middleware"
	"github.com/ericoliveiras/alert-bot-go/response"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
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
	url := cfg.Discord.DiscordBotInvite + oauthStateString
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

func GetGuilds(w http.ResponseWriter, r *http.Request, token *oauth2.Token) ([]response.GuildResponse, error) {
	resp, err := handler.GetInfo(w, r, token, cfg.Discord.GetGuildsInfoUrl)
	if err != nil {
		log.Printf("Error getting guilds information: %s", err.Error())
		http.Error(w, "Error getting guilds information", http.StatusInternalServerError)
		return nil, err
	}
	defer resp.Body.Close()

	var userGuilds []response.GuildResponse
	err = json.NewDecoder(resp.Body).Decode(&userGuilds)
	if err != nil {
		log.Printf("Error decoding user's guilds JSON response: %s", err.Error())
		return nil, err
	}

	var userGuildsOwner []response.GuildResponse
	for _, guild := range userGuilds {
		if guild.Owner {
			userGuildsOwner = append(userGuildsOwner, guild)
		}
	}

	var guildBotIsPresent []response.GuildResponse
	for _, guild := range userGuildsOwner {
		botIsPresent := handler.BotIsInGuild(cfg.Discord.Token, guild.ID, cfg.Discord.ClientID)
		if botIsPresent {
			guildBotIsPresent = append(guildBotIsPresent, guild)
		}
	}

	guildsWithChannels := handler.FetchGuildTextChannels(guildBotIsPresent, cfg.Discord.Token)

	return guildsWithChannels, nil
}
