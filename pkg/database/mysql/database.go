package mysql

import (
	"github.com/copolio/namegpt/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func InitConnection() {
	conf := config.Get()
	dsn := conf.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Error connecting database: " + err.Error())
	}
	database = db
}

func Get() *gorm.DB {
	return database
}
