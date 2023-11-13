package server

import (
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/db"
	"github.com/ericoliveiras/alert-bot-go/discord"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	DB      *sqlx.DB
	Config  *config.Config
	Discord *discordgo.Session
}

func NewServer(config *config.Config) *Server {
	conn, err := db.Init(config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	discord, err := discord.InitDiscord(config.Discord.Token)
	if err != nil {
		log.Fatalf("Error initializing Discord: %v", err)
	}

	return &Server{
		DB:      conn,
		Config:  config,
		Discord: discord,
	}
}

func (server *Server) Start(addr string) error {
	log.Printf("Server running at :%s\n", addr)
	return http.ListenAndServe(":"+addr, nil)
}
