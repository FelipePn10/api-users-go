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

// handler implementa a interface UserHandler
type handler struct {
	service userservice.UserService
}

// UserHandler define os métodos que nosso handler deve implementar
type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}
