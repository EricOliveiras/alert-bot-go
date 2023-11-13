package config

import "os"

type TwitchConfig struct {
	ClientID       string
	ClientSecret   string
	TokenURL       string
	SearchStramURL string
}

func LoadTwitchConfig() TwitchConfig {
	return TwitchConfig{
		ClientID:       os.Getenv("CLIENT_ID"),
		ClientSecret:   os.Getenv("CLIENT_SECRET"),
		TokenURL:       os.Getenv("TOKEN_URL"),
		SearchStramURL: os.Getenv("SEARCH_STREAM_URL"),
	}
}
