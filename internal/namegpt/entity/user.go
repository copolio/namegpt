package entity

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/pkg/database/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name string
}

func init() {
	curConfig := config.Get()
	if curConfig.DdlAuto == config.CREATE {
		log.Default().Println("Start creating USER table")
		err := mysql.Get().AutoMigrate(&User{})
		if err != nil {
			log.Default().Fatal("Error creating USER table: " + err.Error())
		}
		log.Default().Println("Created USER table")
	}
}
