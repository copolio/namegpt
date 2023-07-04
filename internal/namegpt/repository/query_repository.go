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
	return &queryGormRepository{
		GormRepository: GormRepository[entity.Query, uint]{
			tx: config.GetGormDB(),
		},
	}
}

type queryGormRepository struct {
	GormRepository[entity.Query, uint]
}

func (r queryGormRepository) FindOrCreate(query entity.Query) (*entity.Query, error) {
	result := r.tx.Preload("DomainNames").Where(&query).FirstOrCreate(&query)
	return &query, result.Error
}
