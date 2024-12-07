package service

import (
	"context"

	"github.com/FelipePn10/api-users-go/tree/main/internal/dto"
	userrepository "github.com/FelipePn10/api-users-go/tree/main/internal/repository/userepository"
)

type Service struct {
	repo userrepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
	UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error
}

func NewUserService(repo userrepository.UserRepository) UserService {
	return &Service{
		repo: repo,
	}
}
