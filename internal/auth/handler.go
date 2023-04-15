package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/io-m/clean/internal/core"
)

func (h *authHandler) HandleRoutes() {
	h.mux.Route("/auth", func(r chi.Router) {
		r.Get("/{token}", h.Verify)
	})
}

// TODO: Make readJSON, writeJSON, writeErrorJSON methods

type authHandler struct {
	authService IAuthService
	mux         *chi.Mux
}

func NewAuthHandler(mux *chi.Mux, as IAuthService) core.IHandleRoutes {
	return &authHandler{
		authService: as,
		mux:         mux,
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
