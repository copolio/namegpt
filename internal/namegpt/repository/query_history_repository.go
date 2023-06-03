package repository

import (
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"gorm.io/gorm"
)

type QueryHistoryRepository interface {
	Save(tx *gorm.DB, queryHistory *entity.QueryHistory) (*entity.QueryHistory, error)
}

func NewQueryHistoryRepository() QueryHistoryRepository {
	return &GormQueryHistoryRepository{}
}

type GormQueryHistoryRepository struct {
}

func (g GormQueryHistoryRepository) Save(tx *gorm.DB, queryHistory *entity.QueryHistory) (*entity.QueryHistory, error) {
	err := tx.Transaction(func(tx2 *gorm.DB) error {
		result := tx2.Save(queryHistory)
		return result.Error
	})
	return queryHistory, err
}
