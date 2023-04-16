package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/io-m/clean/internal/user/models"
)

func ToUserFromUserRequest(u *models.UserRequest) *models.User {
	user := models.User{
		Id:       uuid.New(),
		Email:    u.Email,
		Password: u.Password,
	}
	return &user
}

func ToUserDaoFromUser(u *models.User) *models.UserDAO {
	ud := models.UserDAO{
		User:           u,
		HashedPassword: "need to have hashing function",
		CreatedAt:      time.Now(),
	}
	return &ud
}

func ToUserResponseFromUserDao(ud *models.UserDAO) *models.UserResponse {
	ur := models.UserResponse{
		User:      ud.User,
		CreatedAt: ud.CreatedAt,
		UpdatedAt: ud.UpdatedAt,
		DeletedAt: ud.DeletedAt,
	}
	return &ur
}
