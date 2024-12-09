package service

import (
	"context"
	"fmt"
	"github.com/FelipePn10/api-users-go/tree/main/internal/dto"
	"github.com/FelipePn10/api-users-go/tree/main/internal/handler/response"
	//"time"
)

func (s *Service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	return nil
}

func (s *Service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
	return nil
}

func (s *Service) GetUserByID(ctx context.Context, id string) (*response.UserResponse, error) {
	userFake := response.UserResponse{
		ID:    "123",
		Name:  "John Doe",
		Email: "john.doe@gmail.com",
		//CreatedAt: time.Now(),
		//UpdatedAt: time.Now(),
	}
	return &userFake, nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (s *Service) FindManyUsers(ctx context.Context) (response.ManyUsersResponse, error) {
	usersFake := response.ManyUsersResponse{}
	for i := 0; i < 5; i++ {
		userFake := response.UserResponse{
			ID:    "123",
			Name:  "John Doe",
			Email: fmt.Sprintf("jonh.doe-%v@gmail.com", i),
			//CreatedAt: time.Now(),
			//UpdatedAt: time.Now(),
		}
		usersFake.Users = append(usersFake.Users, userFake)
	}
	return usersFake, nil
}

func (s *Service) UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error {
	fmt.Println("new password: ", u.Password)
	fmt.Println("old password: ", u.OldPassword)
	return nil
}
