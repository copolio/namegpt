package main

import (
	"github.com/copolio/namegpt/docs"
	"github.com/copolio/namegpt/internal/namegpt/router"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	docs.SwaggerInfo.BasePath = "/api"
	ginEngine := gin.Default()
	router.SetV0Routes(ginEngine)
	router.SetV1Routes(ginEngine)
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	ginEngine.Run()
}
