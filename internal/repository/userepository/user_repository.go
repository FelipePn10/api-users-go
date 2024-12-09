package userrepository

import (
	"context"
	"errors"
	"github.com/FelipePn10/api-users-go/tree/main/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func (r *repository) CreateUser(ctx context.Context, u *entity.UserEntity) error {
	return nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	userExists, err := s.repo.FindUserByEmail(ctx, u.Email)
	if err != nil {
		slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
		return nil, err
	}

	if userExists != nil {
		slog.Error("user already exists", "err", err, slog.String("package", "userservice"))
		return nil, errors.New("user already exists")
	}

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
		return nil, errors.New("error to encrypt password")
	}
	return &entity.UserEntity{Password: string(passwordEncrypted)}, nil
}

func (r *repository) FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *entity.UserEntity) error {
	return nil
}

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (r *repository) FindManyUsers(ctx context.Context) ([]entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) UpdatePassword(ctx context.Context, pass, id string) error {
	return nil
}
