package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DSN string
}

var config *Config

func Load(configMode ConfigMode) error {
	env := ".env." + configMode.String()
	err := godotenv.Load(env)
	if err != nil {
		log.Println("Error loading environmental variables")
		return err
	}

	dsn := os.Getenv("MYSQL_DSN")
	config = new(Config)
	config.DSN = dsn
	return nil
}

func Get() *Config {
	return config
}
