package repository

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"gorm.io/gorm"
)

type QueryHistoryRepository interface {
	WithTransaction(transaction any)
	Save(queryHistory entity.QueryHistory) (*entity.QueryHistory, error)
}

func NewQueryHistoryRepository() QueryHistoryRepository {
	return &QueryHistoryGormRepository{
		tx: config.GetGormDB(),
	}
}

type QueryHistoryGormRepository struct {
	tx *gorm.DB
}

func (g QueryHistoryGormRepository) WithTransaction(transaction any) {
	g.tx = transaction.(*gorm.DB)
}

func (g QueryHistoryGormRepository) Save(queryHistory entity.QueryHistory) (*entity.QueryHistory, error) {
	result := g.tx.Save(&queryHistory)
	return &queryHistory, result.Error
}
