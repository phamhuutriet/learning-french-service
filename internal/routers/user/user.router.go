package user

import (
	"learning-french-service/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	// userController controller.UserController
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, err := wire.InitUserRouterHandler()
	if err != nil {
		panic(err)
	}

	// public router
	userRouterPublic := Router.Group("user")
	{
		userRouterPublic.GET("/register", userController.Register)
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
