package models

import (
	"net/http"

	"github.com/google/uuid"
)

type IUserRepository interface {
	FindOne(id uuid.UUID) (*UserDAO, error)
	FetchMany() ([]*UserDAO, error)
	Update(user *User) (*UserDAO, error)
	Create(user *User) (*UserDAO, error)
	Delete(id uuid.UUID) error
}

type IUserHttpHandler interface {
	FindOne(http.ResponseWriter, *http.Request)
	// FetchMany(http.ResponseWriter, *http.Request)
	// Update(http.ResponseWriter, *http.Request)
	// Create(http.ResponseWriter, *http.Request)
	// Delete(http.ResponseWriter, *http.Request)
}
