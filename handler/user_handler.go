package handler

import (
	"encoding/json"
	"net/http"

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
	UserService *service.UserService
}

func NewUserHandler(db *sqlx.DB) *UserHandler {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := UserHandler{UserService: userService}

	return &userHandler
}

func (uc *UserHandler) Create(w http.ResponseWriter, r *http.Request, resp *http.Response) {
	var user request.CreateUser

	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	err = uc.UserService.Create(r.Context(), &user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if !middleware.IsAuthenticated(r) {
		http.Error(w, "Missing auth header", http.StatusUnauthorized)
		return
	}

	cookie, _ := r.Cookie("access_token")
	token := &oauth2.Token{
		AccessToken: cookie.Value,
	}

	resp, err := utils.GetInfo(w, r, token, cfg.Discord.GetUserInfoUrl)
	if err != nil {
		http.Error(w, "Error getting user information", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user request.CreateUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding JSON response", http.StatusInternalServerError)
		return
	}

	userGuilds, err := GetGuilds(w, r, token)
	if err != nil {
		http.Error(w, "Error fetching user's guilds", http.StatusInternalServerError)
		return
	}

	response := response.UserGuildsResponse{
		User:   user,
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
