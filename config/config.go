package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DSN     string
	DdlAuto DdlAutoMode
}

var config *Config

func init() {
	log.Default().Println("Init Config")
	err := loadConfig(DEVELOPMENT)
	if err != nil {
		log.Fatal("Failed to load configuration.")
	}
}

func loadConfig(configMode ConfigMode) error {
	env := ".env." + configMode.String()
	err := godotenv.Load(env)
	if err != nil {
		log.Fatal("Error loading environmental variables")
		return err
	}

	dsn := os.Getenv("MYSQL_DSN")
	ddl := MapToDdlAuto(os.Getenv("DDL_AUTO"))
	config = &Config{
		DSN:     dsn,
		DdlAuto: ddl,
	}
	return nil
}

func Get() *Config {
	return config
}
