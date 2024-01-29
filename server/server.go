package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/ericoliveiras/alert-bot-go/db"
	"github.com/ericoliveiras/alert-bot-go/discord"
	"github.com/ericoliveiras/alert-bot-go/middleware"
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

	disc, err := discord.InitDiscord(config.Discord.Token)
	if err != nil {
		log.Fatalf("Error initializing Discord: %v", err)
	}

	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			discord.Bot(context.Background(), disc, conn)
		}
	}()

	return &Server{
		DB:      conn,
		Config:  config,
		Discord: disc,
	}
}

func (server *Server) Start(addr string) error {
	log.Printf("Server running at :%s\n", addr)
	return http.ListenAndServe(":"+addr, middleware.CorsConfig())
}
