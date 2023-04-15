package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/io-m/clean/internal/core"
)

// TODO: Make readJSON, writeJSON, writeErrorJSON methods

type authHandler struct {
	authService IAuthService
}

func NewAuthHandler(as IAuthService) core.IAuthHttpHandler {
	return &authHandler{
		authService: as,
	}
}

func (ah *authHandler) Verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := chi.URLParam(r, "token")
	if ok := ah.authService.VerifyJWT(token); !ok {
		http.Error(w, fmt.Errorf("find user handler: %w", errors.New("token is not valid")).Error(), http.StatusBadRequest)
		return
	}
	response := make(map[string]string)
	response["token"] = token
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}
