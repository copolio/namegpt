package controller

import (
	_ "github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/service"
	"github.com/copolio/namegpt/pkg/dto/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type QueryController struct {
	queryUseCase service.QueryUseCase
}

func NewQueryController() *QueryController {
	return &QueryController{queryUseCase: service.NewQueryUseCase()}
}

// SearchDomainNames
// @Summary Get similar domain names given input
// @Description Creates a user in database.
// @Tags v1
// @Param q query request.SearchDomainNames true "Similar domain name query"
// @Produce  json
// @Router /v1/search [get]
// @Success 200 {array} entity.DomainName "domain names"
func (q QueryController) SearchDomainNames(c *gin.Context) {
	query := request.SearchDomainNames{}
	if err := c.Bind(&query); err != nil {
		log.Default().Printf("Request binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	domainNames, err := q.queryUseCase.Handle(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, domainNames)
	}
}
