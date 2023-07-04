package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type ArticlesSourceRoutes struct {
	handler handlers.ArticlesSourceHandlerInterface
}

func NewArticlesSourceRoutes(handler handlers.ArticlesSourceHandlerInterface) *ArticlesSourceRoutes {
	route := &ArticlesSourceRoutes{
		handler: handler,
	}
	return route
}

func (route *ArticlesSourceRoutes) Setup(r *gin.Engine) {
	routes := r.Group("articles-sources")
	{
		routes.GET("get/topicid", middlewares.CheckAccessToken(), route.handler.GetByTopicIDPaginate)
		routes.GET("get/id", middlewares.CheckAccessToken(), route.handler.GetWithID)
		routes.GET("get/most-active", middlewares.CheckAccessToken(), route.handler.GetMostActiveSources)
		routes.GET("list/all", middlewares.CheckAccessToken(), route.handler.ListAll)

	}
}
