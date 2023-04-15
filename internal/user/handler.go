package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/io-m/clean/internal/core"
)

// TODO: Make readJSON, writeJSON, writeErrorJSON methods

type userHandler struct {
	userService IUserService
}

func NewUserHandler(us IUserService) core.IUserHttpHandler {
	return &userHandler{
		userService: us,
	}
}

func (uh *userHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := uuid.MustParse(chi.URLParam(r, "id"))
	fmt.Println("ID --> ", id)
	user, err := uh.userService.FindOne(id)
	if err != nil {
		// w.WriteHeader(http.StatusNotFound)
		http.Error(w, fmt.Errorf("find user handler: %w", err).Error(), http.StatusNotFound)
		return
		// json.NewEncoder(w).Encode(util_error.NewError(err, util_error.Driver, "User not found"))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// func (cr *chiHandler) FetchMany() http.HandlerFunc {
// 	cr.router.Get(path, f)
// }

// func (cr *chiHandler) Update() http.HandlerFunc {
// 	cr.router.Put(path, f)
// }

// func (cr *chiHandler) Create() http.HandlerFunc {
// 	cr.router.Post(path, f)
// }

// func (cr *chiHandler) Delete() http.HandlerFunc {
// 	cr.router.Delete(path, f)
// }
