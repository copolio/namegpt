package service

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/repository"
)

type UserUseCase interface {
	CreateUser(name string) (user *entity.User, err error)
	GetUser(name string) (user *entity.User, err error)
}

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserUseCase() UserUseCase {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (u UserService) CreateUser(name string) (*entity.User, error) {
	return u.userRepository.Save(entity.User{
		Name: name,
	})
}

func (u UserService) GetUser(name string) (*entity.User, error) {
	return u.userRepository.FindByName(name)
}
