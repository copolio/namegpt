package config

import (
	"github.com/copolio/namegpt/pkg/client"
	"github.com/copolio/namegpt/pkg/database"
	"github.com/spf13/viper"
	"log"
	"os"
	"regexp"
)

type Server struct {
	Port int `yaml:"port"`
}

type AppConfig struct {
	Server  Server              `yaml:"server"`
	Mysql   database.Datasource `yaml:"mysql"`
	Chatgpt client.ApiConfig    `yaml:"chatgpt"`
}

var NameGptAppConfig *AppConfig

func init() {
	log.Default().Println("Init AppConfig")
	projectDirName := "namegpt"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(string(rootPath))
	err := viper.ReadInConfig()
	if err != nil {
		log.Default().Fatalf("Error loading app configuration: \n%v", err)
	}
	err = viper.Unmarshal(&NameGptAppConfig)
	if err != nil {
		log.Default().Fatalf("Error unmarshalling app configuration: \n%v", err)
	}
}
