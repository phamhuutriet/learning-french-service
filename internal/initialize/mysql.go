package initialize

import (
	"fmt"
	"learning-french-service/global"
	"learning-french-service/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errStr string) {
	if err != nil {
		global.Logger.Error(errStr, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	checkErrorPanic(err, "failed to open mysql")
	global.Logger.Info("mysql connected", zap.String("dsn", dsn))
	global.MDb = db

	// Set pool
	SetPool()

	// Migrate tables
	migrateTables()
}

func SetPool() {
	sqlDB, err := global.MDb.DB()
	checkErrorPanic(err, "failed to get sql db")

	sqlDB.SetConnMaxLifetime(time.Duration(global.Config.MySQL.ConnMaxLifetime) * time.Second)
	sqlDB.SetMaxIdleConns(global.Config.MySQL.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.Config.MySQL.MaxOpenConns)
}

func migrateTables() {
	err := global.MDb.AutoMigrate(&po.User{}, &po.Role{})
	checkErrorPanic(err, "failed to migrate tables")
}
