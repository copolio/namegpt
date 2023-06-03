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
// @Description Creates a user in database.
// @Tags v0
// @Param user body request.CreateUser true "Create user request"
// @Accept  json
// @Produce  json
// @Router /v0/users [post]
// @Success 200 {object} entity.User "user"
func (controller UserController) CreateUser(c *gin.Context) {
	createUser := request.CreateUser{}
	if err := c.ShouldBindBodyWith(&createUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	user, db := controller.userUseCase.CreateUser(createUser.Name)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, db.Error)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUser
// @Summary Gets a user info by name.
// @Description Gets a user info from database.
// @Tags v0
// @Param name path string true "Username"
// @Accept  json
// @Produce  json
// @Router /v0/users/{name} [get]
// @Success 200 {object} entity.User "user"
func (controller UserController) GetUser(c *gin.Context) {
	name := c.Param("name")
	user, db := controller.userUseCase.GetUser(name)
	if db.Error != nil {
		c.JSON(http.StatusNoContent, struct{}{})
	} else {
		c.JSON(http.StatusOK, user)
	}
}
