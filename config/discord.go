package config

import "os"

type DiscordConfig struct {
	Token        string
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string
	RedirectURL  string
	Scopes       []string
}

func LoadDiscordConfig() DiscordConfig {
	return DiscordConfig{
		Token:        os.Getenv("DISCORD_TOKEN"),
		ClientID:     os.Getenv("DISCORD_CLIENT_ID"),
		ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
		AuthURL:      os.Getenv("DISCORD_AUTH_URL"),
		TokenURL:     os.Getenv("DISCORD_AUTH_TOKEN"),
		RedirectURL:  os.Getenv("DISCORD_URL_REDIRECT"),
		Scopes:       []string{"identify", "email"},
	}
}
