package entity

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/pkg/database/mysql"
	"gorm.io/gorm"
	"log"
)

type Query struct {
	gorm.Model
	Keyword     string `gorm:"uniqueIndex;type:char(255)"`
	DomainNames []DomainName
}

func init() {
	curConfig := config.Get()
	if curConfig.DdlAuto == config.CREATE {
		log.Default().Println("Start creating QUERY table")
		err := mysql.GetGormDB().AutoMigrate(&Query{})
		if err != nil {
			log.Default().Fatal("Error creating QUERY table: " + err.Error())
		}
		log.Default().Println("Created QUERY table")
	}
}

func (Query) TableName() string {
	return "query"
}
