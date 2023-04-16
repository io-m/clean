package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/io-m/clean/internal/user/models"
	"github.com/io-m/clean/pkg/utils"
)

type IUserService interface {
	FindOne(uuid.UUID) (*models.UserResponse, error)
	FetchMany() ([]*models.UserResponse, error)
	Update(*models.UserRequest) (*models.UserResponse, error)
	Create(*models.UserRequest) (*models.UserResponse, error)
	Delete(uuid.UUID) error
}

type userService struct {
	userRepository models.IUserRepository
}

func NewUserService(ur models.IUserRepository) IUserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) FindOne(id uuid.UUID) (*models.UserResponse, error) {
	ud, err := us.userRepository.FindOne(id)
	if err != nil {
		return nil, fmt.Errorf("find user service: %w", err)
	}
	return utils.ToUserResponseFromUserDao(ud), nil
}

func (us *userService) FetchMany() ([]*models.UserResponse, error) {
	uds, err := us.userRepository.FetchMany()
	if err != nil {
		return nil, fmt.Errorf("user_service:fetching_many: %w", err)
	}
	var urs []*models.UserResponse
	for _, ud := range uds {
		ur := utils.ToUserResponseFromUserDao(ud)
		urs = append(urs, ur)
	}
	return urs, nil
}

func (us *userService) Create(ur *models.UserRequest) (*models.UserResponse, error) {
	user := utils.ToUserFromUserRequest(ur)
	ud, err := us.userRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("user_service:creating_user: %w", err)
	}
	return utils.ToUserResponseFromUserDao(ud), nil
}

func (us *userService) Update(ur *models.UserRequest) (*models.UserResponse, error) {
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
