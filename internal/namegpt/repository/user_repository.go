package repository

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/pkg/database/mysql"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.User) *entity.User
	FindByName(name string) (*entity.User, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &GormUserRepository{
		db: mysql.Get(),
	}
}

func (u GormUserRepository) Save(user *entity.User) *entity.User {
	u.db.Save(user)
	return user
}

func (u GormUserRepository) FindByName(name string) (*entity.User, error) {
	var user entity.User
	result := u.db.Where(&entity.User{Name: name}).First(&user)
	return &user, result.Error
}
