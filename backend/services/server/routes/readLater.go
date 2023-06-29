package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type ReadLaterRoutes struct {
	handler handlers.ReadLaterHandlerInterface
}

func NewReadLaterRoutes(handler handlers.ReadLaterHandlerInterface) *ReadLaterRoutes {
	route := &ReadLaterRoutes{
		handler: handler,
	}
	return route
}

func (route *ReadLaterRoutes) Setup(r *gin.Engine) {
	routes := r.Group("read-later")
	{
		routes.POST("add", middlewares.CheckAccessToken(), route.handler.AddToReadLaterList)
		routes.POST("remove", middlewares.CheckAccessToken(), route.handler.RemoveFromReadLaterList)
	}
}
