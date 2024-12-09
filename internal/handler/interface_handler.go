package userhandler

import (
	"net/http"

	userservice "github.com/FelipePn10/api-users-go/tree/main/internal/service/useservice"
)

func NewUserHandler(service userservice.UserService) UserHandler {
	return &handler{
		service: service,
	}
}

type handler struct {
	service userservice.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindManyUsers(w http.ResponseWriter, r *http.Request)
	UpdateUserPassword(w http.ResponseWriter, r *http.Request)
}
