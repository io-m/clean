package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/io-m/clean/internal/user/core"
	"github.com/io-m/clean/pkg/utils"
)

type IUserService interface {
	FindOne(uuid.UUID) (*core.UserResponse, error)
	FetchMany() ([]*core.UserResponse, error)
	Update(*core.UserRequest) (*core.UserResponse, error)
	Create(*core.UserRequest) (*core.UserResponse, error)
	Delete(uuid.UUID) error
}

type userService struct {
	userRepository core.IUserRepository
}

func NewUserService(ur core.IUserRepository) IUserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) FindOne(id uuid.UUID) (*core.UserResponse, error) {
	ud, err := us.userRepository.FindOne(id)
	if err != nil {
		return nil, fmt.Errorf("find user service: %w", err)
	}
	return utils.ToUserResponseFromUserDao(ud), nil
}

func (us *userService) FetchMany() ([]*core.UserResponse, error) {
	uds, err := us.userRepository.FetchMany()
	if err != nil {
		return nil, fmt.Errorf("user_service:fetching_many: %w", err)
	}
	var urs []*core.UserResponse
	for _, ud := range uds {
		ur := utils.ToUserResponseFromUserDao(ud)
		urs = append(urs, ur)
	}
	return urs, nil
}

func (us *userService) Create(ur *core.UserRequest) (*core.UserResponse, error) {
	user := utils.ToUserFromUserRequest(ur)
	ud, err := us.userRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("user_service:creating_user: %w", err)
	}
	return utils.ToUserResponseFromUserDao(ud), nil
}

func (us *userService) Update(ur *core.UserRequest) (*core.UserResponse, error) {
	user := utils.ToUserFromUserRequest(ur)
	ud, err := us.userRepository.Update(user)
	if err != nil {
		return nil, fmt.Errorf("user_service:updating_user: %w", err)
	}
	return utils.ToUserResponseFromUserDao(ud), nil
}

func (us *userService) Delete(id uuid.UUID) error {
	err := us.userRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("user_service:deleting_user: %w", err)
	}
	return nil
}
