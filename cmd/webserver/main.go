package main

import (
	"fmt"
	"github.com/FelipePn10/api-users-go/tree/main/internal/dto"
	"log/slog"
	"net/http"

	logger "github.com/FelipePn10/api-users-go/tree/main/config"
	"github.com/FelipePn10/api-users-go/tree/main/config/env"
	"github.com/FelipePn10/api-users-go/tree/main/internal/database"
	"github.com/FelipePn10/api-users-go/tree/main/internal/database/sqlc"
	userhandler "github.com/FelipePn10/api-users-go/tree/main/internal/handler"
	"github.com/FelipePn10/api-users-go/tree/main/internal/handler/routes"
	userrepository "github.com/FelipePn10/api-users-go/tree/main/internal/repository/userepository"
	userservice "github.com/FelipePn10/api-users-go/tree/main/internal/service/useservice"
	"github.com/go-chi/chi/v5"
)

func main() {
	logger.InitLogger()
	slog.Info("starting api")

	_, err := env.LoadingConfig("C:\\Users\\Felipe\\OneDrive\\Desktop\\api-users-go\\cmd\\webserver")
	if err != nil {
		slog.Error("failed to load environment variables", err, slog.String("package", "main"))
		return
	}

	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
		return
	}

	queries := sqlc.New(dbConnection)

	userDto := &dto.UpdateUserPasswordDto{
		BasicUserDto: dto.BasicUserDto{
			Name:     "John Doe",
			Email:    "johndoe@example.com",
			Password: "StrongPassword@123",
		},
		OldPassword: "OldPassword@123",
	}

	userRepo := userrepository.NewUserRepository(dbConnection, queries, userDto)
	newUserService := userservice.NewUserService(userRepo)
	newUserHandler := userhandler.NewUserHandler(newUserService)

	router := chi.NewRouter()
	routes.InitUserRoutes(router, newUserHandler)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}
	slog.Info("server stopped")
	return
}
