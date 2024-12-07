package routes

import (
	"net/http"

	userhandler "github.com/FelipePn10/api-users-go/tree/main/internal/handler"
	"github.com/go-chi/chi/v5"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Route("/users", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			err := h.CreateUser(w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	})
}
