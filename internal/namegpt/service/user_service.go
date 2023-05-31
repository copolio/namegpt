package service

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/repository"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(name string) (*entity.User, *gorm.DB)
	GetUser(name string) (*entity.User, *gorm.DB)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService() UserService {
	return &UserServiceImpl{
		userRepository: repository.NewUserRepository(),
	}
}

func (u UserServiceImpl) CreateUser(name string) (*entity.User, *gorm.DB) {
	return u.userRepository.Save(&entity.User{
		Name: name,
	})
}

func (u UserServiceImpl) GetUser(name string) (*entity.User, *gorm.DB) {
	return u.userRepository.FindByName(name)
}
