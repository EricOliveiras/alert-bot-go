package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/handler"
	"github.com/ericoliveiras/alert-bot-go/middleware"
	"github.com/ericoliveiras/alert-bot-go/request"
	"github.com/ericoliveiras/alert-bot-go/response"
	"golang.org/x/oauth2"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	if !middleware.IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	cookie, _ := r.Cookie("access_token")
	token := &oauth2.Token{
		AccessToken: cookie.Value,
	}

	resp, err := handler.GetInfo(w, r, token, cfg.Discord.GetUserInfoUrl)
	if err != nil {
		log.Printf("Error getting user information: %s", err.Error())
		http.Error(w, "Error getting user information", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user request.CreateUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding JSON response: %s", err.Error())
		http.Error(w, "Error decoding JSON response", http.StatusInternalServerError)
		return
	}

	userGuilds, err := GetGuilds(w, r, token)
	if err != nil {
		log.Printf("Error fetching user's guilds: %s", err.Error())
		http.Error(w, "Error fetching user's guilds", http.StatusInternalServerError)
		return
	}

	response := response.UserGuildsResponse{
		User:   user,
		Guilds: userGuilds,
	}

	responseJSON, err := json.Marshal(&response)
	if err != nil {
		log.Printf("Error encoding JSON response: %s", err.Error())
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		log.Printf("Error writing JSON response: %s", err.Error())
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		return
	}
}
