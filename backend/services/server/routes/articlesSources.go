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
		routes.GET("list/all/paging", middlewares.CheckAccessToken(), route.handler.ListAllPaging)
		routes.GET("count/total", middlewares.CheckAccessToken(), route.handler.Count)
		routes.GET("search/filter", middlewares.CheckAccessToken(), route.handler.SearchWithFilter)
		routes.POST("delete/id", middlewares.CheckAccessToken(), route.handler.Delete)
		routes.POST("update/id", middlewares.CheckAccessToken(), route.handler.Update)

	}
}
