package initialize

import (
	"fmt"
	"learning-french-service/global"

	"go.uber.org/zap"
)

func Run() {
	// load config
	LoadConfig()
	fmt.Println("Loading config", global.Config.MySQL)

	// init logger
	InitLogger()
	global.Logger.Info("Logger initialized", zap.String("config", global.Config.MySQL.Host))

	// init mysql
	InitMysql()

	// init redis
	InitRedis()

	// init router
	router := InitRouter()
	router.Run(":8080")
}
