package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/io-m/clean/internal/core"
)

func ToUserFromUserRequest(u *core.UserRequest) *core.User {
	user := core.User{
		Id:       uuid.New(),
		Email:    u.Email,
		Password: u.Password,
	}
	return &user
}

func ToUserDaoFromUser(u *core.User) *core.UserDAO {
	ud := core.UserDAO{
		User:           u,
		HashedPassword: "need to have hashing function",
		CreatedAt:      time.Now(),
	}
	return &ud
}

func ToUserResponseFromUserDao(ud *core.UserDAO) *core.UserResponse {
	ur := core.UserResponse{
		User:      ud.User,
		CreatedAt: ud.CreatedAt,
		UpdatedAt: ud.UpdatedAt,
		DeletedAt: ud.DeletedAt,
	}
	return &ur
}
