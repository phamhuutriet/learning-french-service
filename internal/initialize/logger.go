package initialize

import (
	"learning-french-service/global"
	"learning-french-service/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(&global.Config.Log)
}
