package main

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/routes"
	"github.com/copolio/namegpt/pkg/database/mysql"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := config.Load(config.DEVELOPMENT)
	if err != nil {
		log.Fatal("Failed to load configuration.")
	}
	mysql.InitConnection()

	router := gin.Default()
	routes.SetV0Routes(router)
	router.Run()
}
