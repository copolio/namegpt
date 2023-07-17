package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var connection *gorm.DB

func init() {
	dsn := NameGptAppConfig.Datasource.GetMysqlDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting MySQL: %v\n", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error retrieving MySQL sql connection: %v\n", err)
	}
	sqlDB.SetMaxIdleConns(NameGptAppConfig.Datasource.MaxIdleConn)
	sqlDB.SetMaxOpenConns(NameGptAppConfig.Datasource.MaxOpenConn)
	connection = db
}

func GetGormDB() (db *gorm.DB) {
	return connection
}
