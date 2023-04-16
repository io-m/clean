package auth

import (
	"github.com/io-m/clean/internal/user/models"
)

type IAuthService interface {
	VerifyJWT(string) bool
}

type authService struct {
	userService models.IUserService
}

func NewAuthService(us models.IUserService) IAuthService {
	return &authService{
		userService: us,
	}
}

func (as *authService) VerifyJWT(token string) bool {
	return token == "abcd"
}
