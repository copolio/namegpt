package service

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (u UserService) CreateUser(name string) *entity.User {
	return u.userRepository.Save(&entity.User{
		Name: name,
	})
}

func (u UserService) GetUser(name string) *entity.User {
	return u.userRepository.FindByName(name)
}
