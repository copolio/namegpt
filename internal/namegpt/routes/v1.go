package routes

import (
	"github.com/copolio/namegpt/internal/namegpt/controller"
	"github.com/gin-gonic/gin"
)

func SetV1Routes(router *gin.Engine) *gin.Engine {
	userController := controller.NewUserController()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userController.CreateUser)
		v1.GET("/users/:name", userController.GetUser)
	}
	return router
}
