package repository

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
)

type QueryRepository interface {
	CrudRepository[entity.Query, uint]
	FindOrCreate(query entity.Query) (*entity.Query, error)
}

func NewQueryRepository() QueryRepository {
	return &QueryGormRepository{
		GormRepository: GormRepository[entity.Query, uint]{
			tx: config.GetGormDB(),
		},
	}
}

type QueryGormRepository struct {
	GormRepository[entity.Query, uint]
}

func (r QueryGormRepository) FindOrCreate(query entity.Query) (*entity.Query, error) {
	result := r.tx.Preload("DomainNames").Where(&query).FirstOrCreate(&query)
	return &query, result.Error
}
