package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var connection *gorm.DB

func init() {
	dsn := NameGptAppConfig.Mysql.GetMysqlDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})
	if err != nil {
		log.Fatalf("Error connecting MySQL: %v\n", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error retrieving MySQL sql connection: %v\n", err)
	}
	sqlDB.SetMaxIdleConns(NameGptAppConfig.Mysql.MaxIdleConn)
	sqlDB.SetMaxOpenConns(NameGptAppConfig.Mysql.MaxOpenConn)
	connection = db
}

func GetGormDB() (db *gorm.DB) {
	return connection
}
