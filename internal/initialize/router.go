package initialize

import (
	"learning-french-service/global"
	"learning-french-service/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	// r.Use()
	// r.Use()
	// r.Use()
	userRouter := routers.RouterGroupApp.User
	deckRouter := routers.RouterGroupApp.Deck

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
		deckRouter.InitDeckRouter(MainGroup)
	}

	return r
}
