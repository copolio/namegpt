package entity

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/pkg/database/mysql"
	"gorm.io/gorm"
	"log"
)

type DomainName struct {
	gorm.Model
	Name    string `gorm:"type:char(255)"`
	QueryID uint
}

func init() {
	curConfig := config.Get()
	if curConfig.DdlAuto == config.CREATE {
		log.Default().Println("Start creating DOMAIN_NAME table")
		err := mysql.GetGormDB().AutoMigrate(&DomainName{})
		if err != nil {
			log.Default().Fatal("Error creating DOMAIN_NAME table: " + err.Error())
		}
		log.Default().Println("Created DOMAIN_NAME table")
	}
}

func (DomainName) TableName() string {
	return "domain_name"
}
