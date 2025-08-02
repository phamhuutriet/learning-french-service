//go:build wireinject

package wire

import (
	"learning-french-service/internal/controller"
	"learning-french-service/internal/repo"
	"learning-french-service/internal/services"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		services.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
