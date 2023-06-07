package entity

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/pkg/database"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name string `gorm:"unique"`
}

func init() {
	if config.NameGptAppConfig.Mysql.Ddl == database.CREATE {
		log.Default().Println("Start creating USER table")
		err := config.GetGormDB().AutoMigrate(&User{})
		if err != nil {
			log.Default().Fatal("Error creating USER table: " + err.Error())
		}
		log.Default().Println("Created USER table")
	}
}

func (User) TableName() string {
	return "user"
}
