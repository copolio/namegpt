package repository

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/pkg/database/mysql"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.User) (*entity.User, *gorm.DB)
	FindByName(name string) (*entity.User, *gorm.DB)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &GormUserRepository{
		db: mysql.Get(),
	}
}

func (u GormUserRepository) Save(user *entity.User) (*entity.User, *gorm.DB) {
	result := u.db.Save(user)
	return user, result
}

func (u GormUserRepository) FindByName(name string) (*entity.User, *gorm.DB) {
	var user entity.User
	result := u.db.Where(&entity.User{Name: name}).First(&user)
	return &user, result
}
