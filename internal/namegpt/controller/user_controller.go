package controller

import (
	_ "github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/service"
	"github.com/copolio/namegpt/pkg/dto/request"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type UserController struct {
	userUseCase service.UserUseCase
}

func NewUserController() *UserController {
	return &UserController{
		userUseCase: service.NewUserUseCase(),
	}
}

// CreateUser
// @Summary Creates a user.
// @MetaData Creates a user in database.
// @Tags v0
// @Param user body request.CreateUser true "Create user request"
// @Accept  json
// @Produce  json
// @Router /v0/users [post]
// @Success 200 {object} entity.User "user"
func (controller UserController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		createUser := request.CreateUser{}
		if err := c.ShouldBindBodyWith(&createUser, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		user, err := controller.userUseCase.CreateUser(createUser.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

// GetUser
// @Summary Gets a user info by name.
// @MetaData Gets a user info from database.
// @Tags v0
// @Param name path string true "Username"
// @Accept  json
// @Produce  json
// @Router /v0/users/{name} [get]
// @Success 200 {object} entity.User "user"
func (controller UserController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		user, err := controller.userUseCase.GetUser(name)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
