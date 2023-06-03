package service

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/repository"
	"gorm.io/gorm"
)

type UserUseCase interface {
	CreateUser(name string) (*entity.User, *gorm.DB)
	GetUser(name string) (*entity.User, *gorm.DB)
}

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserUseCase() UserUseCase {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (u UserService) CreateUser(name string) (*entity.User, *gorm.DB) {
	return u.userRepository.Save(&entity.User{
		Name: name,
	})
}

func (u UserService) GetUser(name string) (*entity.User, *gorm.DB) {
	return u.userRepository.FindByName(name)
}
