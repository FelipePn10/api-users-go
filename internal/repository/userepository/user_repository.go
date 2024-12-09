package userrepository

import (
	"context"
	"errors"
	//"github.com/FelipePn10/api-users-go/tree/main/internal/database"
	//"github.com/FelipePn10/api-users-go/tree/main/internal/database/sqlc"
	"github.com/FelipePn10/api-users-go/tree/main/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func (r *repository) CreateUser(ctx context.Context, u *entity.UserEntity) error {
	return nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	userExists, err := r.FindUserByEmail(ctx, email)
	if err != nil {
		slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
		return nil, err
	}

	if userExists != nil {
		slog.Error("user already exists", "err", err, slog.String("package", "userservice"))
		return nil, errors.New("user already exists")
	}
	return nil, nil
}

func (r *repository) FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error) {
	userExits, err := r.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to find user by id", "err", err, slog.String("package", "userservice"))
		return nil, err
	}
	if userExits != nil {
		slog.Error("user already exists", "err", err, slog.String("package", "userservice"))
		return nil, errors.New("user already exists")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userExits.Password), []byte(r.user.OldPassword))
	if err != nil {
		slog.Error("invalid password", "err", err, slog.String("package", "userservice"))
		return nil, errors.New("invalid password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userExits.Password), []byte(r.user.Password))
	if err == nil {
		slog.Error("new password is equal to old password", "err", err, slog.String("package", "userservice"))
		return nil, errors.New("new password is equal to old password")
	}
	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(r.user.Password), 12)
	if err != nil {
		slog.Error("error encrypt password", "err", err, slog.String("package", "userservice"))
		return nil, errors.New("error encrypt password")
	}
	err = r.UpdatePassword(ctx, string(passwordEncrypted), id)
	if err != nil {
		slog.Error("error update password", "err", err, slog.String("package", "userservice"))
		return nil, err
	}
	return &entity.UserEntity{Password: string(passwordEncrypted)}, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *entity.UserEntity) error {
	return nil
}

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (r *repository) FindManyUsers(ctx context.Context) ([]*entity.UserEntity, error) {
	users, err := r.queries.FindManyUsers(ctx)
	if err != nil {
		slog.Error("error retrieving users", "err", err, slog.String("package", "userrepository"))
		return nil, err
	}

	var userEntities []*entity.UserEntity
	for _, user := range users {
		userEntities = append(userEntities, &entity.UserEntity{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			//CreatedAt: user.CreatedAt,
			//UpdatedAt: user.UpdatedAt,
		})
	}

	return userEntities, nil
}

func (r *repository) UpdatePassword(ctx context.Context, pass, id string) error {
	return nil
}
