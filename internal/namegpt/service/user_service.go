package service

import (
	"errors"
	"fmt"
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/middleware"
	"github.com/copolio/namegpt/internal/namegpt/repository"
	"gorm.io/gorm"
	"net/http"
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
	user, err := u.userRepository.FindByName(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user service error: %w", &middleware.ResponseStatusError{
			Code:     http.StatusNotFound,
			Message:  fmt.Sprintf("user %s does not exist", name),
			MetaData: err.Error(),
		})
	}
	return user, err
}
