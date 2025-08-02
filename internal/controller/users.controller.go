package controller

import (
	"learning-french-service/internal/services"
	"learning-french-service/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	email := c.Query("email")
	purpose := c.Query("purpose")
	code := uc.userService.Register(email, purpose)
	response.SuccessResponse(c, code, nil)
}
