package repository

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
)

type QueryHistoryRepository interface {
	CrudRepository[entity.QueryHistory, uint]
}

func NewQueryHistoryRepository() QueryHistoryRepository {
	return &QueryHistoryGormRepository{
		GormRepository: GormRepository[entity.QueryHistory, uint]{
			tx: config.GetGormDB(),
		},
	}
}

type QueryHistoryGormRepository struct {
	GormRepository[entity.QueryHistory, uint]
}
