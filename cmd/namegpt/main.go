package main

import (
	"github.com/copolio/namegpt/api/swagger"
	"github.com/copolio/namegpt/internal/namegpt/middleware"
	"github.com/copolio/namegpt/internal/namegpt/router"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	swagger.SwaggerInfo.Title = "NameGPT API"
	swagger.SwaggerInfo.Description = "This is a NameGPT API server."
	swagger.SwaggerInfo.Version = "1.0"
	swagger.SwaggerInfo.BasePath = "/api"
	ginEngine := gin.Default()
	ginEngine.Use(middleware.ErrorHandler())
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.SetV0Routes(ginEngine)
	router.SetV1Routes(ginEngine)
	err := ginEngine.Run()
	if err != nil {
		log.Fatal(err)
	}
}
