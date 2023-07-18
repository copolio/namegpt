package repository

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"time"
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

var _ QueryRepository = &queryGormRepository{}

type queryGormRepository struct {
	GormRepository[entity.Query, uint]
}

func (r queryGormRepository) FindOrCreate(query entity.Query) (*entity.Query, error) {
	result := r.tx.Preload("DomainNames", "created_at > ?", time.Now().AddDate(0, -1, 0)).Where(&query).FirstOrCreate(&query)
	return &query, result.Error
}
