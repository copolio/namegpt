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
	Save(user entity.User) (*entity.User, error)
	FindByName(name string) (*entity.User, error)
}

type UserGormRepository struct {
	GormRepository[entity.User, uint]
}

func NewUserRepository() UserRepository {
	return &UserGormRepository{
		GormRepository: GormRepository[entity.User, uint]{
			tx: config.GetGormDB(),
		},
	}
}

func (r UserGormRepository) FindByName(name string) (*entity.User, error) {
	var user entity.User
	result := r.tx.Where(&entity.User{Name: name}).Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user repository error: %w", &middleware.ResponseStatusError{
			Code:     http.StatusNotFound,
			Message:  fmt.Sprintf("user %s does not exist", name),
			MetaData: result.Error.Error(),
		})
	}
	return &user, result.Error
}
