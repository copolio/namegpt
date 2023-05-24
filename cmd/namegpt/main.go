package main

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/router"
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

	r := gin.Default()
	router.SetV0Routes(r)
	router.SetV1Routes(r)
	r.Run()
}
