package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetV0Routes(router *gin.Engine) *gin.Engine {
	v0 := router.Group("/api/v0")
	{
		v0.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	return router
}
