package global

import (
	"learning-french-service/pkg/logger"
	"learning-french-service/settings"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *settings.Config
	Logger *logger.LoggerZap
	MDb    *gorm.DB
	Rdb    *redis.Client
)
