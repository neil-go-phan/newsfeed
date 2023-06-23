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
		routes.GET("get-by-topicid", middlewares.CheckAccessToken(), route.handler.GetByTopicIDPaginate)
	}
}
