package routers

import (
	"learning-french-service/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)
