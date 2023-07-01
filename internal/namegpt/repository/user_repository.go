package repository

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
)

type UserRepository interface {
	CrudRepository[entity.User, uint]
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
	return &user, result.Error
}
