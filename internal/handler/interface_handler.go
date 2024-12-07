package userhandler

import (
	"net/http"

	userservice "github.com/FelipePn10/api-users-go/tree/main/internal/service/useservice"
)

func NewUserHandler(service userservice.UserService) UserHandler {
	return &handler{
		service,
	}
}

type handler struct {
	service userservice.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request) error
}
