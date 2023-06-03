package repository

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/pkg/database/mysql"
	"gorm.io/gorm"
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
		db: mysql.GetGormDB(),
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
	return &user, result.Error
}
