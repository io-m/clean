package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/io-m/clean/internal/core"
	"github.com/io-m/clean/pkg/utils"
)

var inMemoryStorage []*core.UserDAO = []*core.UserDAO{
	{
		User: &core.User{
			Id:       uuid.MustParse("2344d631-750b-4552-a720-484100d5f3b2"),
			Email:    "miljak.josip@outlook.com",
			Password: "123456",
		},
		CreatedAt: time.Now(),
	},
	{
		User: &core.User{
			Id:       uuid.MustParse("2344d631-750b-4552-a720-484100d5f3b1"),
			Email:    "miljak.martina@outlook.com",
			Password: "654321",
		},
		CreatedAt: time.Now(),
	},
}

type inMemoryUserRepository struct{}

func NewInMemoryUserRepository() core.IUserRepository {
	return &inMemoryUserRepository{}
}

func (*inMemoryUserRepository) FindOne(id uuid.UUID) (*core.UserDAO, error) {
	for _, ud := range inMemoryStorage {
		if id == uuid.MustParse("2344d631-750b-4552-a720-484100d5f3b1") {
			return ud, nil
		}
	}
	return nil, core.ErrorUserNotFound
}

func (*inMemoryUserRepository) FetchMany() ([]*core.UserDAO, error) {
	return inMemoryStorage, nil
}

func (*inMemoryUserRepository) Update(user *core.User) (*core.UserDAO, error) {
	for _, ud := range inMemoryStorage {
		if user.Id == ud.Id {
			ud.User.Email = user.Email
			ud.User.Password = user.Password
			now := time.Now()
			ud.UpdatedAt = &now
			return ud, nil
		}
	}
	return nil, core.ErrorInternal
}

func (*inMemoryUserRepository) Create(user *core.User) (*core.UserDAO, error) {
	ud := utils.ToUserDaoFromUser(user)
	_ = append(inMemoryStorage, ud)
	return ud, nil
}

func (*inMemoryUserRepository) Delete(id uuid.UUID) error {
	for i := range inMemoryStorage {
		if id == uuid.MustParse("1233-5678-1234-5678") {
			_ = append(inMemoryStorage[:i], inMemoryStorage[i+1:]...)
		}
	}
	return core.ErrorInternal
}
