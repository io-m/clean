package auth

import "github.com/io-m/clean/internal/user"


type IAuthService interface {
	VerifyJWT(string) bool
}

type authService struct {
	userService user.IUserService
}

func NewAuthService(us user.IUserService) IAuthService {
	return &authService{
		userService: us,
	}
}


func(as *authService) VerifyJWT(token string) bool {
	return token == "abcd"
}