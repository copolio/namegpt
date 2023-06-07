package repository

import (
	"errors"
	"fmt"
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/middleware"
	"gorm.io/gorm"
	"net/http"
)

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindByName(name string) (*entity.User, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &GormUserRepository{
		db: config.GetGormDB(),
	}
}

func (u GormUserRepository) Save(user *entity.User) (*entity.User, error) {
	err := u.db.Transaction(func(tx2 *gorm.DB) error {
		result := tx2.Save(user)
		return result.Error
	})
	return user, err
}

func (u GormUserRepository) FindByName(name string) (*entity.User, error) {
	var user entity.User
	result := u.db.Where(&entity.User{Name: name}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user repository error: %w", &middleware.ResponseStatusError{
			Code:     http.StatusNotFound,
			Message:  fmt.Sprintf("user %s does not exist", name),
			MetaData: result.Error.Error(),
		})
	}
	return &user, result.Error
}
