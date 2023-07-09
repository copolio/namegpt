package main

import (
	"fmt"
	"github.com/copolio/namegpt/api/swagger"
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/middleware"
	"github.com/copolio/namegpt/internal/namegpt/router"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func main() {
	swagger.SwaggerInfo.Title = "NameGPT API"
	swagger.SwaggerInfo.Description = "This is a NameGPT API server."
	swagger.SwaggerInfo.Version = "1.0"
	pwd, _ := os.Getwd()
	ginEngine := gin.Default()
	ginEngine.Use(middleware.Cors)
	ginEngine.Use(middleware.ErrorHandler)
	ginEngine.Use(static.Serve("/", static.LocalFile(pwd+"/web/namegpt-ui/out", true)))
	ginEngine.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.SetV0Routes(ginEngine)
	router.SetV1Routes(ginEngine)
	err := ginEngine.Run(fmt.Sprintf(":%d", config.NameGptAppConfig.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
