package main

import (
	application "github.com/ericoliveiras/alert-bot-go"
	"github.com/ericoliveiras/alert-bot-go/config"
)

func main() {
	config := config.NewConfig()

	application.Start(config)
}
