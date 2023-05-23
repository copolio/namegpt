package controller

import (
	"github.com/copolio/namegpt/internal/namegpt/service"
	"github.com/copolio/namegpt/pkg/dto/request"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (controller UserController) CreateUser(c *gin.Context) {
	createUser := request.CreateUser{}
	if err := c.ShouldBindBodyWith(&createUser, binding.JSON); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	user := controller.userService.CreateUser(createUser.Name)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (controller UserController) GetUser(c *gin.Context) {
	name := c.Param("name")
	user := controller.userService.GetUser(name)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
