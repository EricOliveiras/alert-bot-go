package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ericoliveiras/alert-bot-go/repository"
	"github.com/ericoliveiras/alert-bot-go/request"
	"github.com/ericoliveiras/alert-bot-go/service"
	"github.com/jmoiron/sqlx"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(db *sqlx.DB) *UserController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := UserController{UserService: userService}

	return &userController
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request, resp *http.Response) {
	var user request.CreateUser

	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding JSON: %s", err.Error())
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	err = uc.UserService.Create(r.Context(), &user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
}
