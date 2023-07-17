package entity

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/pkg/database"
	"gorm.io/gorm"
	"log"
)

type DomainName struct {
	gorm.Model
	Name    string `gorm:"type:char(255)" json:"name"`
	QueryID uint
}

func init() {
	if config.NameGptAppConfig.Datasource.Ddl == database.CREATE {
		log.Default().Println("Start creating DOMAIN_NAME table")
		err := config.GetGormDB().AutoMigrate(&DomainName{})
		if err != nil {
			log.Default().Fatal("Error creating DOMAIN_NAME table: " + err.Error())
		}
		log.Default().Println("Created DOMAIN_NAME table")
	}
}

func (DomainName) TableName() string {
	return "domain_name"
}
