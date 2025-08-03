//go:build wireinject

package wire

import (
	"learning-french-service/internal/controller"
	"learning-french-service/internal/repo"
	"learning-french-service/internal/services"

	"github.com/google/wire"
)

func InitDeckRouterHandler() (*controller.DeckController, error) {
	wire.Build(
		repo.NewDeckRepository,
		services.NewDeckService,
		controller.NewDeckController,
	)

	return new(controller.DeckController), nil
}
