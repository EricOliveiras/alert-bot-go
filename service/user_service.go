package service

import (
	"context"
	"time"

	"github.com/ericoliveiras/alert-bot-go/builder"
	"github.com/ericoliveiras/alert-bot-go/repository"
	"github.com/ericoliveiras/alert-bot-go/request"
	"github.com/google/uuid"
)

type UserServiceWrapper interface {
	Create(ctx context.Context, user *request.CreateUser) error
}

type UserService struct {
	Repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{Repository: repository}
}

func (us *UserService) Create(ctx context.Context, user *request.CreateUser) error {
	createUser := builder.NewUserBuilder().
		SetID(uuid.New()).
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetChannelLimit(1).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Build()

	err := us.Repository.Create(ctx, &createUser)
	if err != nil {
		return err
	}

	return nil
}
