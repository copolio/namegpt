package entity

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/pkg/database"
	"gorm.io/gorm"
	"log"
)

type QueryHistory struct {
	gorm.Model
	QueryID uint
	Query   Query
	UserId  string
}

func init() {
	if config.NameGptAppConfig.Datasource.Ddl == database.CREATE {
		log.Default().Println("Start creating QUERY_HISTORY table")
		err := config.GetGormDB().AutoMigrate(&QueryHistory{})
		if err != nil {
			log.Default().Fatal("Error creating QUERY_HISTORY table: " + err.Error())
		}
		log.Default().Println("Created QUERY_HISTORY table")
	}
}

func (QueryHistory) TableName() string {
	return "query_history"
}
