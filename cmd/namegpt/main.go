package main

import (
	"github.com/copolio/namegpt/internal/namegpt/router"
	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()
	router.SetV0Routes(ginEngine)
	router.SetV1Routes(ginEngine)
	ginEngine.Run()
}
