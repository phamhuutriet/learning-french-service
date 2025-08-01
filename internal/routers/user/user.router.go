package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	// userController controller.UserController
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userRouterPublic := Router.Group("user")
	{
		userRouterPublic.GET("/register")
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("user")
	// userRouterPrivate.Use(middlewares.LimiterMiddleware())
	// userRouterPrivate.Use(middlewares.AuthMiddleware())
	// userRouterPrivate.Use(middlewares.PermissionMiddleware())

	{
		userRouterPrivate.GET("/info")
	}
}
