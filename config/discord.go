package config

import "os"

type DiscordConfig struct {
	Token            string
	ClientID         string
	ClientSecret     string
	AuthURL          string
	TokenURL         string
	RedirectURL      string
	GetUserInfoUrl   string
	GetGuildsInfoUrl string
	DiscordBotInvite string
	Scopes           []string
}

func LoadDiscordConfig() DiscordConfig {
	return DiscordConfig{
		Token:            os.Getenv("DISCORD_TOKEN"),
		ClientID:         os.Getenv("DISCORD_CLIENT_ID"),
		ClientSecret:     os.Getenv("DISCORD_CLIENT_SECRET"),
		AuthURL:          os.Getenv("DISCORD_AUTH_URL"),
		TokenURL:         os.Getenv("DISCORD_AUTH_TOKEN"),
		RedirectURL:      os.Getenv("DISCORD_URL_REDIRECT"),
		GetUserInfoUrl:   os.Getenv("USER_INFO_URL"),
		GetGuildsInfoUrl: os.Getenv("GUILD_INFO_URL"),
		DiscordBotInvite: os.Getenv("DISCORD_BOT_INVITE"),
		Scopes:           []string{"identify", "email", "guilds"},
	}
}
