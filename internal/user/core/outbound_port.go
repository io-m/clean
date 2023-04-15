package core

import "github.com/google/uuid"

type IUserRepository interface {
	FindOne(id uuid.UUID) (*UserDAO, error)
	FetchMany() ([]*UserDAO, error)
	Update(user *User) (*UserDAO, error)
	Create(user *User) (*UserDAO, error)
	Delete(id uuid.UUID) error
}
