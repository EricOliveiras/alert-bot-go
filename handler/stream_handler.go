package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/middleware"
	"github.com/ericoliveiras/alert-bot-go/repository"
	"github.com/ericoliveiras/alert-bot-go/request"
	"github.com/ericoliveiras/alert-bot-go/service"
	"github.com/jmoiron/sqlx"
)

type StreamHandler struct {
	StreamService *service.StreamService
}

func NewStreamHandler(db *sqlx.DB) *StreamHandler {
	streamRepository := repository.NewStreamRepository(db)
	discordRepository := repository.NewDiscordRepository(db)
	discordStreamRepository := repository.NewDiscordStreamRepository(db)
	streamService := service.NewStreamService(
		streamRepository,
		discordRepository,
		discordStreamRepository,
	)
	streamHandler := StreamHandler{
		streamService,
	}
	return &streamHandler
}

func (sh *StreamHandler) CreateStream(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !middleware.IsAuthenticated(r) {
		http.Error(w, "Missing auth header", http.StatusUnauthorized)
		return
	}

	var streamRequest request.StreamRequest
	if err := json.NewDecoder(r.Body).Decode(&streamRequest); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := sh.StreamService.Create(r.Context(), &streamRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
