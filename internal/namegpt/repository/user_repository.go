package repository

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/pkg/database/mysql"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: mysql.Get(),
	}
}

func (u UserRepository) Save(user *entity.User) *entity.User {
	u.db.Save(user)
	return user
}

func (u UserRepository) FindByName(name string) *entity.User {
	var user entity.User
	u.db.Where(&entity.User{Name: name}).First(&user)
	return &user
}
