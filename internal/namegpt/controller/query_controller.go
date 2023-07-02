package controller

import (
	"errors"
	"fmt"
	_ "github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/service"
	"github.com/copolio/namegpt/pkg/client/chatgpt"
	"github.com/copolio/namegpt/pkg/dto"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
	"strings"
)

type QueryController struct {
	queryUseCase service.QueryUseCase
}

func NewQueryController() *QueryController {
	return &QueryController{queryUseCase: service.NewQueryUseCase()}
}

// GenerateSimilarDomains
// @Summary Get similar domain names given input
// @MetaData Creates a user in database.
// @Tags v1
// @Param q body request.SimilarDomainNames true "Generate similar domain name"
// @Produce  json
// @Router /api/v1/domains/similar-names [post]
// @Success 200 {array} entity.DomainName "domain names"
func (q QueryController) GenerateSimilarDomains() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := dto.SimilarDomainNames{}
		if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
			c.Error(err)
			return
		}

		domainNames, err := q.queryUseCase.Handle(request)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, domainNames)
	}
}

// GenerateRecommendedDomains
// @Summary Generates domain names in server sent event.
// @MetaData Generates domain recommendations
// @Tags v1
// @Param user body request.RecommendDomainNames true "Generate domain name recommendation"
// @Accept  json
// @Produce  json
// @Router /api/v1/domains/recommendations [post]
// @Success 201 {object} []string "domain"
func (q QueryController) GenerateRecommendedDomains() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := dto.RecommendDomainNames{}
		if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
			c.Error(err)
			return
		}
		c.Stream(func(w io.Writer) bool {
			stream, err := chatgpt.GetRecommendedDomainStream(request.Description)
			if err != nil {
				c.Error(err)
				return false
			}
			defer stream.Close()

			fmt.Println("Stream start")
			var domains = make([]string, 50)
			var domain = ""
			for {
				recv, err := stream.Recv()
				if errors.Is(err, io.EOF) {
					fmt.Println("\nStream finished")
					return false
				}
				if err != nil {
					fmt.Errorf("Stream Recv failed: %w", err)
					return false
				}
				fmt.Printf("Recv: %s\n", recv.Choices[0].Delta.Content)
				if strings.Contains(recv.Choices[0].Delta.Content, "\"") {
					if domain != "" {
						domains = append(domains, domain)
						c.SSEvent("domain", domain)
						fmt.Printf("Domain: %s\n", domain)
					}
					domain = ""
				} else {
					domain += recv.Choices[0].Delta.Content
				}
			}
			fmt.Println("Stream end")
			stream.Close()
			return true
		})
	}
}
