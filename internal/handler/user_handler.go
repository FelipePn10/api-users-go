package userhandler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/FelipePn10/api-users-go/tree/main/internal/dto"
	"github.com/FelipePn10/api-users-go/tree/main/internal/handler/httperr"
	"github.com/FelipePn10/api-users-go/tree/main/internal/handler/validation"
)

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var req dto.CreateUserDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "userhandler"))
		return sendErrorResponse(w, httperr.NewBadRequestError("body is required"))
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", slog.Any("error", err), slog.String("package", "handler_user"))
		return sendErrorResponse(w, httperr.NewBadRequestError("invalid request body"))
	}

	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_user"))
		return sendErrorResponse(w, httpErr)
	}

	return nil
}

func sendErrorResponse(w http.ResponseWriter, err *httperr.RestErr) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
	return err
}
