package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HttpStatus int

type ResponseStatusError struct {
	Code     HttpStatus `json:"code"`
	Message  string     `json:"message"`
	MetaData string     `json:"metaData"`
}

func (e *ResponseStatusError) Error() string {
	return e.Message
}

var ErrorHandler gin.HandlerFunc

func init() {
	ErrorHandler = func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			log.Default().Printf("Error: %v\n", err)
			target := &ResponseStatusError{}
			if errors.As(err, &target) {
				c.JSON(int(target.Code), target)
				return
			}
		}
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, &ResponseStatusError{
				Code:     http.StatusInternalServerError,
				Message:  "Internal Server Error",
				MetaData: c.Errors.Last().Error(),
			})
		}
	}
}
