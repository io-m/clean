package models

import (
	"time"

	"github.com/google/uuid"
)

// Base User entity
type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
}

// UserDAO is extended user entity used to save in DB
type UserDAO struct {
	*User
	HashedPassword string     `json:"-"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

// UserRequest is entity received from driver actors
type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// UserResponse is entity returned back as response
type UserResponse struct {
	*User
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
