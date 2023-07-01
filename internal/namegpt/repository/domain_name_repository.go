package repository

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
)

type DomainNameRepository interface {
	CrudRepository[entity.DomainName, uint]
}

func NewDomainNameRepository() DomainNameRepository {
	return &DomainNameGormRepository{
		GormRepository: GormRepository[entity.DomainName, uint]{
			tx: config.GetGormDB(),
		},
	}
}

type DomainNameGormRepository struct {
	GormRepository[entity.DomainName, uint]
}
