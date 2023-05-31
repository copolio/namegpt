package service

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/repository"
)

type UserService interface {
	CreateUser(name string) *entity.User
	GetUser(name string) (*entity.User, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService() UserService {
	return &UserServiceImpl{
		userRepository: repository.NewUserRepository(),
	}
}

func (u UserServiceImpl) CreateUser(name string) *entity.User {
	return u.userRepository.Save(&entity.User{
		Name: name,
	})
}

func (u UserServiceImpl) GetUser(name string) (*entity.User, error) {
	return u.userRepository.FindByName(name)
}
