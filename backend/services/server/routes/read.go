package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type ReadRoutes struct {
	handler handlers.ReadHandlerInterface
}

func NewReadRoutes(handler handlers.ReadHandlerInterface) *ReadRoutes {
	route := &ReadRoutes{
		handler: handler,
	}
	return route
}

func (route *ReadRoutes) Setup(r *gin.Engine) {
	routes := r.Group("read")
	{
		routes.POST("read", middlewares.CheckAccessToken(), route.handler.Read)
		routes.POST("unread", middlewares.CheckAccessToken(), route.handler.Unread)
	}
}
