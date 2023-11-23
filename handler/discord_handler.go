package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/auth"
	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/middleware"
	"github.com/ericoliveiras/alert-bot-go/repository"
	"github.com/ericoliveiras/alert-bot-go/request"
	"github.com/ericoliveiras/alert-bot-go/response"
	"github.com/ericoliveiras/alert-bot-go/service"
	"github.com/ericoliveiras/alert-bot-go/utils"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
)

type DiscordHandler struct {
	DiscordService *service.DiscordChannelService
}

func NewDiscordHandler(db *sqlx.DB) *DiscordHandler {
	discordRepository := repository.NewDiscordRepository(db)
	userRepository := repository.NewUserRepository(db)
	discordService := service.NewDiscordChannelService(discordRepository, userRepository)
	discordHandler := DiscordHandler{DiscordService: discordService}

	return &discordHandler
}

func (dc *DiscordHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !middleware.IsAuthenticated(r) {
		http.Error(w, "Missing auth header", http.StatusUnauthorized)
		return
	}

	var newChannelInfo request.CreateDiscordChannel
	err := json.NewDecoder(r.Body).Decode(&newChannelInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dc.DiscordService.Create(r.Context(), &newChannelInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

var cfg = config.NewConfig()

var discordOauthConfig = utils.GetDiscordOAuthConfig()
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
		http.Error(w, "Invalid state!", http.StatusInternalServerError)
		return
	}

	code := r.FormValue("code")
	token, err := discordOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Error exchanging code for token", http.StatusInternalServerError)
		return
	}

	auth.SaveCookie(w, token)

	resp, err := utils.GetInfo(w, r, token, cfg.Discord.GetUserInfoUrl)
	if err != nil {
		http.Error(w, "Error exchanging code for token", http.StatusBadRequest)
	}
	defer resp.Body.Close()

	userController := NewUserHandler(db)

	userController.Create(w, r, resp)
}

func GetGuilds(w http.ResponseWriter, r *http.Request, token *oauth2.Token) ([]response.GuildResponse, error) {
	resp, err := utils.GetInfo(w, r, token, cfg.Discord.GetGuildsInfoUrl)
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
		botIsPresent := utils.BotIsInGuild(cfg.Discord.Token, guild.ID, cfg.Discord.ClientID)
		if botIsPresent {
			guildBotIsPresent = append(guildBotIsPresent, guild)
		}
	}

	guildsWithChannels := utils.FetchGuildTextChannels(guildBotIsPresent, cfg.Discord.Token)

	return guildsWithChannels, nil
}
