package userrepository

import (
	"context"
	"database/sql"
	"github.com/FelipePn10/api-users-go/tree/main/internal/database/sqlc"
	"github.com/FelipePn10/api-users-go/tree/main/internal/dto"
	"github.com/FelipePn10/api-users-go/tree/main/internal/entity"
)

func NewUserRepository(db *sql.DB, q *sqlc.Queries, user *dto.UpdateUserPasswordDto) UserRepository {
	return &repository{
		db,
		q,
		user,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
	user    *dto.UpdateUserPasswordDto
}

type UserRepository interface {
	CreateUser(ctx context.Context, u *entity.UserEntity) error
	FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error)
	UpdateUser(ctx context.Context, u *entity.UserEntity) error
	DeleteUser(ctx context.Context, id string) error
	FindManyUsers(ctx context.Context) ([]*entity.UserEntity, error)
	UpdatePassword(ctx context.Context, pass, id string) error
}
