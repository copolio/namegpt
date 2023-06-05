package controller

import (
	_ "github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/service"
	"github.com/copolio/namegpt/pkg/dto/request"
	"github.com/gin-gonic/gin"
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
// @MetaData Creates a user in database.
// @Tags v1
// @Param q query request.SearchDomainNames true "Similar domain name query"
// @Produce  json
// @Router /v1/search [get]
// @Success 200 {array} entity.DomainName "domain names"
func (q QueryController) SearchDomainNames() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := request.SearchDomainNames{}
		if err := c.Bind(&query); err != nil {
			c.Error(err)
			return
		}

		domainNames, err := q.queryUseCase.Handle(query)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, domainNames)
	}
}
