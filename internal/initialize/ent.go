package initialize

import (
	"context"
	"fmt"
	"learning-french-service/global"
	"learning-french-service/internal/ent"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func InitEnt() {
	m := global.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.Dbname)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		global.Logger.Error("failed to open mysql", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("mysql connected", zap.String("dsn", dsn))

	global.EntClient = client

	// Run auto migration
	if err := client.Schema.Create(context.Background()); err != nil {
		global.Logger.Error("failed to create schema", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("database schema created successfully")

	// Create initial data if needed
	createInitialData(client)
}

func createInitialData(client *ent.Client) {
	ctx := context.Background()

	// Check if we already have users
	count, err := client.User.Query().Count(ctx)
	if err != nil {
		global.Logger.Error("failed to count users", zap.Error(err))
		return
	}

	// If no users exist, create a default admin user
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			global.Logger.Error("failed to hash password", zap.Error(err))
			return
		}

		_, err = client.User.Create().
			SetEmail("admin@example.com").
			SetPasswordHash(string(hashedPassword)).
			SetUsername("admin").
			SetFirstName("Admin").
			SetLastName("User").
			SetCurrentLevel("A1").
			SetTargetLevel("C2").
			SetDailyGoal(20).
			SetTimezone("UTC").
			SetIsActive(true).
			Save(ctx)

		if err != nil {
			global.Logger.Error("failed to create admin user", zap.Error(err))
			return
		}

		global.Logger.Info("created default admin user")
	}
}
