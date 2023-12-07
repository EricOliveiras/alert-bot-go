package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/auth"
	"github.com/ericoliveiras/alert-bot-go/middleware"
	"github.com/ericoliveiras/alert-bot-go/repository"
	"github.com/ericoliveiras/alert-bot-go/request"
	"github.com/ericoliveiras/alert-bot-go/response"
	"github.com/ericoliveiras/alert-bot-go/service"
	"github.com/ericoliveiras/alert-bot-go/utils"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
)

type UserHandler struct {
	UserService   *service.UserService
	DiscordServer *service.DiscordChannelService
}

func NewUserHandler(db *sqlx.DB) *UserHandler {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	channelRepository := repository.NewDiscordRepository(db)
	discordService := service.NewDiscordChannelService(channelRepository, userRepository)
	userHandler := UserHandler{UserService: userService, DiscordServer: discordService}

	return &userHandler
}

func (uc *UserHandler) Create(w http.ResponseWriter, r *http.Request, resp *http.Response, discordToken *oauth2.Token) {
	var user request.CreateUser

	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	returnedUser, err := uc.UserService.Create(r.Context(), &user)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	jwtToken, err := utils.GenerateToken(returnedUser.ID, returnedUser.DiscordID, discordToken)
	if err != nil {
		http.Error(w, "Failed to create JWT token", http.StatusInternalServerError)
		return
	}

	auth.SaveTokenCookie(w, jwtToken)

	http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
}

func (uc *UserHandler) UserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !middleware.IsAuthenticated(r) {
		http.Error(w, "Missing auth header", http.StatusUnauthorized)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	cookie, _ := r.Cookie("jwt_token")

	id, _, token, err := utils.GetIDsFromToken(cookie.Value)
	if err != nil {
		http.Error(w, "Error getting cookie value", http.StatusInternalServerError)
		return
	}

	user, err := uc.UserService.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userResponse := response.UserResponse{
		ID:           user.ID,
		DiscordID:    user.DiscordID,
		Username:     user.Username,
		Email:        user.Email,
		Avatar:       user.Avatar,
		ChannelLimit: user.ChannelLimit,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	if user.ChannelLimit == 0 {
		channel, err := uc.DiscordServer.GetChannelByUserID(r.Context(), user.ID)
		if err != nil {
			http.Error(w, "Error getting channel", http.StatusInternalServerError)
			return
		}

		channelFormat := response.ChannelResponse{
			ID:          channel.ID,
			Name:        channel.Name,
			ChannelId:   channel.ChannelId,
			ServerId:    channel.ServerId,
			StreamLimit: channel.StreamLimit,
			UserID:      channel.UserID,
			CreatedAt:   channel.CreatedAt,
			UpdatedAt:   channel.UpdatedAt,
		}

		userChannelResponse := response.UserChannelResponse{
			User:    userResponse,
			Channel: channelFormat,
		}

		responseJSON, err := json.Marshal(&userChannelResponse)
		if err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(responseJSON)
		if err != nil {
			http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
			return
		}

		return
	}

	userGuilds, err := GetGuilds(w, r, token)
	if err != nil {
		http.Error(w, "Error fetching user's guilds", http.StatusInternalServerError)
		return
	}

	response := response.UserGuildsResponse{
		User:   userResponse,
		Guilds: userGuilds,
	}

	responseJSON, err := json.Marshal(&response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		return
	}
}
