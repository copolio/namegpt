package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DSN          string
	DdlAuto      DdlAutoMode
	ChatgptToken string
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

	chatgptToken := os.Getenv("CHATGPT_TOKEN")
	dsn := os.Getenv("MYSQL_DSN")
	ddl := MapToDdlAuto(os.Getenv("DDL_AUTO"))
	config = &Config{
		DSN:          dsn,
		DdlAuto:      ddl,
		ChatgptToken: chatgptToken,
	}
	return nil
}

func Get() *Config {
	return config
}
