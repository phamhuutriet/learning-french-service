package deck

import (
	"learning-french-service/internal/wire"

	"github.com/gin-gonic/gin"
)

type DeckRouter struct {
	// deckController controller.DeckController
}

func (dr *DeckRouter) InitDeckRouter(Router *gin.RouterGroup) {
	deckController, err := wire.InitDeckRouterHandler()
	if err != nil {
		panic(err)
	}

	// private router - requires authentication
	deckRouterPrivate := Router.Group("decks")
	// deckRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		deckRouterPrivate.GET("", deckController.GetDecks)
		deckRouterPrivate.POST("", deckController.CreateDeck)
		deckRouterPrivate.GET("/:id", deckController.GetDeck)
		deckRouterPrivate.PUT("/:id", deckController.UpdateDeck)
		deckRouterPrivate.DELETE("/:id", deckController.DeleteDeck)
	}
}
