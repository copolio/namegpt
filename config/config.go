package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
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
	projectDirName := "namegpt"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	env := "/.env." + configMode.String()
	err := godotenv.Load(string(rootPath) + env)
	if err != nil {
		log.Fatalf("Error loading environmental variables: %v", err)
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
