package mysql

import (
	"github.com/copolio/namegpt/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var connection *gorm.DB

func init() {
	conf := config.Get()
	dsn := conf.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})
	if err != nil {
		log.Fatal("Error connecting connection: " + err.Error())
	}
	connection = db
}

func GetGormDB() (db *gorm.DB) {
	return connection
}
