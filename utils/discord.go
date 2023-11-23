package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/response"
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

func BotIsInGuild(botToken, guildID, userID string) bool {
	url := fmt.Sprintf("https://discord.com/api/v10/guilds/%s/members/%s", guildID, userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %s", err.Error())
		return false
	}

	req.Header.Set("Authorization", "Bot "+botToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error checking bot membership: %s", err.Error())
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func GetGuildChannels(botToken, guildID string) []response.GuildChannelsResponse {
	url := fmt.Sprintf("https://discord.com/api/v10//guilds/%s/channels", guildID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %s", err.Error())
		return nil
	}

	req.Header.Set("Authorization", "Bot "+botToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error checking bot membership: %s", err.Error())
		return nil
	}
	defer resp.Body.Close()

	var guildChats []response.GuildChannelsResponse
	err = json.NewDecoder(resp.Body).Decode(&guildChats)
	if err != nil {
		log.Printf("Error decoding guild channels: %s", err.Error())
		return nil
	}

	var textChannels []response.GuildChannelsResponse
	for _, chat := range guildChats {
		if chat.Type == 0 {
			textChannels = append(textChannels, chat)
		}
	}

	return textChannels
}

func FetchGuildTextChannels(guilds []response.GuildResponse, botToken string) []response.GuildResponse {
	var guildsWithTextChannels []response.GuildResponse

	for _, guild := range guilds {
		channels := GetGuildChannels(botToken, guild.ID)

		var textChannels []response.GuildChannelsResponse
		for _, channel := range channels {
			if channel.Type == 0 {
				textChannels = append(textChannels, channel)
			}
		}

		guildWithTextChannels := response.GuildResponse{
			Icon:     guild.Icon,
			ID:       guild.ID,
			Name:     guild.Name,
			Owner:    guild.Owner,
			Channels: textChannels,
		}

		guildsWithTextChannels = append(guildsWithTextChannels, guildWithTextChannels)
	}

	return guildsWithTextChannels
}
