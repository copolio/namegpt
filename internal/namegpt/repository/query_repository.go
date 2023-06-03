package repository

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"gorm.io/gorm"
)

type QueryRepository interface {
	FindOrCreate(tx *gorm.DB, query *entity.Query) (*entity.Query, error)
}

func NewQueryRepository() QueryRepository {
	return &GormQueryRepository{}
}

type GormQueryRepository struct {
}

func (g GormQueryRepository) FindOrCreate(tx *gorm.DB, query *entity.Query) (*entity.Query, error) {
	err := tx.Transaction(func(tx2 *gorm.DB) error {
		result := tx2.Preload("DomainNames").Where(query).FirstOrCreate(query)
		return result.Error
	})
	return query, err
}
