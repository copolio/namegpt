package repository

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
)

type QueryHistoryRepository interface {
	WithTransaction(transaction any)
	Save(queryHistory entity.QueryHistory) (*entity.QueryHistory, error)
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
