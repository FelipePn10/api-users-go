package routes

import (
	userhandler "github.com/FelipePn10/api-users-go/tree/main/internal/handler"
	"github.com/go-chi/chi/v5"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Patch("/{id}", h.UpdateUser)
		r.Get("/{id}", h.GetUserById)
		r.Delete("/{id}", h.DeleteUser)
		r.Get("/", h.FindManyUsers)
		r.Patch("/password/{id}", h.UpdateUserPassword)
	})
}
