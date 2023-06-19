package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type TopicRoutes struct {
	handler handlers.TopicHandlerInterface
}

func NewTopicRoutes(handler handlers.TopicHandlerInterface) *TopicRoutes {
	route := &TopicRoutes{
		handler: handler,
	}
	return route
}

func (route *TopicRoutes) Setup(r *gin.Engine) {
	routes := r.Group("topic")
	{
		routes.GET("list", middlewares.CheckAccessToken(), route.handler.List)
		routes.GET("get-page", middlewares.CheckAccessToken(), route.handler.GetPagination)
		routes.GET("count", middlewares.CheckAccessToken(), route.handler.Count)
		routes.POST("create", middlewares.CheckAccessToken(), route.handler.Create)
		routes.POST("update", middlewares.CheckAccessToken(), route.handler.Update)
		routes.POST("delete", middlewares.CheckAccessToken(), route.handler.Delete)
	}
}
