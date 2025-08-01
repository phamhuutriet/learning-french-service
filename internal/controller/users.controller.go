package controller

import (
	"learning-french-service/internal/services"
	"learning-french-service/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usersService *services.UsersService
}

func NewUserController() *UserController {
	return &UserController{
		usersService: services.NewUsersService(),
	}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users := uc.usersService.GetUsers()
	response.SuccessResponse(c, response.ErrCodeSuccess, users)
}
