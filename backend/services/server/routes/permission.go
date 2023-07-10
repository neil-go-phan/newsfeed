package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type PermissionRoutes struct {
	handler handlers.PermissionHandlerInterface
}

func NewPermissionRoutes(handler handlers.PermissionHandlerInterface) *PermissionRoutes {
	route := &PermissionRoutes{
		handler: handler,
	}
	return route
}

func (route *PermissionRoutes) Setup(r *gin.Engine) {
	routes := r.Group("permission")
	{
		routes.GET("list", middlewares.CheckAccessToken(), route.handler.List)
	}
}
