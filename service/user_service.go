package service

import (
	"context"
	"errors"
	"time"

	"github.com/ericoliveiras/alert-bot-go/builder"
	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/ericoliveiras/alert-bot-go/repository"
	"github.com/ericoliveiras/alert-bot-go/request"
	"github.com/google/uuid"
)

type UserServiceWrapper interface {
	Create(ctx context.Context, user *request.CreateUser) error
	GetById(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type UserService struct {
	Repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{Repository: repository}
}

func (us *UserService) Create(ctx context.Context, user *request.CreateUser) (*models.User, error) {
	existUser, err := us.Repository.GetByEmail(ctx, user.Email)
	if existUser != nil && err == nil {
		return existUser, nil
	}

	createUser := builder.NewUserBuilder().
		SetID(uuid.New()).
		SetDiscordID(user.DiscordID).
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetAvatar(user.Avatar).
		SetChannelLimit(1).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Build()

	err = us.Repository.Create(ctx, &createUser)
	if err != nil {
		return &models.User{}, err
	}

	return &createUser, nil
}

func (us *UserService) GetById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := us.Repository.GetById(ctx, id)
	if err != nil {
		return &models.User{}, errors.New("user not found")
	}

	return user, nil
}
