package repository

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"gorm.io/gorm"
)

type DomainNameRepository interface {
	WithTransaction(transaction any)
	Save(domainName entity.DomainName) (*entity.DomainName, error)
}

func NewDomainNameRepository() DomainNameRepository {
	return DomainNameGormRepository{tx: config.GetGormDB()}
}

type DomainNameGormRepository struct {
	tx *gorm.DB
}

func (d DomainNameGormRepository) WithTransaction(transaction any) {
	d.tx = transaction.(*gorm.DB)
}

func (d DomainNameGormRepository) Save(domainName entity.DomainName) (*entity.DomainName, error) {
	result := d.tx.Save(&domainName)
	return &domainName, result.Error
}
