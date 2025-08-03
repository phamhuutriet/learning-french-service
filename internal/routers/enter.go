package routers

import (
	"learning-french-service/internal/routers/deck"
	"learning-french-service/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Deck deck.DeckRouterGroup
}

var RouterGroupApp = new(RouterGroup)
