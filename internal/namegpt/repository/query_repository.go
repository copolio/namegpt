package repository

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"gorm.io/gorm"
)

type QueryRepository interface {
	WithTransaction(transaction any)
	FindOrCreate(query entity.Query) (*entity.Query, error)
}

func NewQueryRepository() QueryRepository {
	return &QueryGormRepository{
		tx: config.GetGormDB(),
	}
}

type QueryGormRepository struct {
	tx *gorm.DB
}

func (g QueryGormRepository) WithTransaction(transaction any) {
	g.tx = transaction.(*gorm.DB)
}

func (g QueryGormRepository) FindOrCreate(query entity.Query) (*entity.Query, error) {
	result := g.tx.Preload("DomainNames").Where(&query).FirstOrCreate(&query)
	return &query, result.Error
}
