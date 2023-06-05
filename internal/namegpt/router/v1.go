package router

import (
	"github.com/copolio/namegpt/internal/namegpt/controller"
	"github.com/gin-gonic/gin"
)

func SetV1Routes(router *gin.Engine) *gin.Engine {
	queryController := controller.NewQueryController()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/search", queryController.SearchDomainNames())
	}
	return router
}
