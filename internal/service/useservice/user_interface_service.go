package service

import (
	"context"
	"github.com/FelipePn10/api-users-go/tree/main/internal/dto"
	"github.com/FelipePn10/api-users-go/tree/main/internal/handler/response"
	userrepository "github.com/FelipePn10/api-users-go/tree/main/internal/repository/userepository"
)

type Service struct {
	repo userrepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
	UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error
	GetUserByID(ctx context.Context, id string) (*response.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
	FindManyUsers(ctx context.Context) ([]*response.UserResponse, error)
	UpdateUserPassword(ctx context.Context, u dto.UpdateUserPasswordDto, id string) error
}

func NewUserService(repo userrepository.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}
