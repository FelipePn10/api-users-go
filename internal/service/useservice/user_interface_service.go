package service

import userrepository "github.com/FelipePn10/api-users-go/tree/main/internal/repository/userepository"

type Service struct {
	repo userrepository.UserRepository
}

type UserService interface {
	CreateUser() error
}

func NewUserService(repo userrepository.UserRepository) UserService {
	return &Service{
		repo: repo,
	}
}
