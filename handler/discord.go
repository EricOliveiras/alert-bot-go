package handler

import (
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/config"
	"golang.org/x/oauth2"
)

var cfg = config.NewConfig()
var discordOauthConfig = GetDiscordOAuthConfig()

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

/*
Return a response according to the url passed. Ex: Users, Guilds...
*/
func GetInfo(w http.ResponseWriter, r *http.Request, token *oauth2.Token, url string) (*http.Response, error) {
	client := discordOauthConfig.Client(r.Context(), token)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
