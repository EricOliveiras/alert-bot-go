package config

import "os"

type DiscordConfig struct {
	Token string
}

func LoadDiscordConfig() DiscordConfig {
	return DiscordConfig{
		Token: os.Getenv("DISCORD_TOKEN"),
	}
}
