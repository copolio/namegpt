package router

import (
	"github.com/copolio/namegpt/internal/namegpt/controller"
	"github.com/gin-gonic/gin"
)

func SetV1Routes(router *gin.Engine) *gin.Engine {
	queryController := controller.NewQueryController()
	v1 := router.Group("/api/v1/domains")
	{
		v1.POST("/similar-names", queryController.GenerateSimilarDomains())
		v1.POST("/recommendations", queryController.GenerateRecommendedDomains())
	}
	return router
}
