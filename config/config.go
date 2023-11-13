package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	DB      DBConfig
	HTTP    HTTPConfig
	Discord DiscordConfig
	Twitch  TwitchConfig
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		DB:      LoadDBConfig(),
		HTTP:    LoadHTTPConfig(),
		Discord: LoadDiscordConfig(),
		Twitch:  LoadTwitchConfig(),
	}
}
