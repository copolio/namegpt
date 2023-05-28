package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetV0Routes(router *gin.Engine) *gin.Engine {
	v0 := router.Group("/api/v0")
	{
		v0.GET("/ping", ping())
	}
	return router
}

// requestPing
// @Summary This is simple Summary.
// @Description This is detail Description.
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
