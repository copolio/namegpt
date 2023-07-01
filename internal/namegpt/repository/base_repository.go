package repository

import (
	"errors"
	"gorm.io/gorm"
)

type CrudRepository[T any, K any] interface {
	WithTransaction(transaction any)
	Save(entity T) (*T, error)
	FindById(id K) (*T, error)
	Delete(entity T) (err error)
}

type GormRepository[T any, K any] struct {
	tx *gorm.DB
}

func (r *GormRepository[T, K]) WithTransaction(transaction any) {
	r.tx = transaction.(*gorm.DB)
}

func (r *GormRepository[T, K]) Save(entity T) (*T, error) {
	result := r.tx.Save(&entity)
	return &entity, result.Error
}

func (r *GormRepository[T, K]) FindById(id K) (*T, error) {
	var entity T
	result := r.tx.Take(&entity)
	if result.RowsAffected > 1 {
		return nil, errors.New("more than single row exists")
	}

	return &entity, result.Error
}

func (r *GormRepository[T, K]) Delete(entity T) (err error) {
	result := r.tx.Delete(&entity)
	if result.RowsAffected > 1 {
		return errors.New("more than single row deleted")
	}
	return result.Error
}
