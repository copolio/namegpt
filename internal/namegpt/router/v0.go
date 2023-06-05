package router

import (
	"github.com/copolio/namegpt/internal/namegpt/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetV0Routes Example routes to practice Gin
func SetV0Routes(router *gin.Engine) *gin.Engine {
	userController := controller.NewUserController()
	v0 := router.Group("/api/v0")
	{
		v0.GET("/ping", ping())
		v0.POST("/users", userController.CreateUser())
		v0.GET("/users/:name", userController.GetUser())
	}
	return router
}

// requestPing
// @Summary Pong!
// @MetaData This is detail MetaData.
// @Tags v0
// @Accept  json
// @Produce  json
// @Router /v0/ping [get]
// @Success 200 {object} string
func ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
