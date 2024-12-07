package service

import (
	"context"
	"github.com/FelipePn10/api-users-go/tree/main/internal/dto"
	"github.com/FelipePn10/api-users-go/tree/main/internal/handler/response"
)

func (s *Service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	return nil
}

func (s *Service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
	return nil
}

func (s *Service) GetUserbyID(ctx context.Context, id string) (*response.UserResponse, error) {
	userFake := response.UserResponse{
		ID:    "123",
		Name:  "John Doe",
		Email: "john.doe@gmail.com",
		//CreatedAt: time.Now(),
		//UpdatedAt: time.Now(),
	}
	return &userFake, nil
}
