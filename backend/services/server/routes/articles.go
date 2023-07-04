package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type ArticleRoutes struct {
	handler handlers.ArticleHandlerInterface
}

func NewArticleRoutes(handler handlers.ArticleHandlerInterface) *ArticleRoutes {
	route := &ArticleRoutes{
		handler: handler,
	}
	return route
}

func (route *ArticleRoutes) Setup(r *gin.Engine) {
	routes := r.Group("articles")
	{
		routes.GET("get/sourceid", middlewares.CheckAccessToken(), route.handler.GetArticlesPaginationByArticlesSourceID)
		routes.GET("get/all/followed", middlewares.CheckAccessToken(), route.handler.GetArticlesPaginationByUserFollowedSources)

		routes.GET("get/sourceid/unread", middlewares.CheckAccessToken(), route.handler.GetUnreadArticlesPaginationByArticlesSourceID)
		routes.GET("get/all/followed/unread", middlewares.CheckAccessToken(), route.handler.GetUnreadArticlesByUserFollowedSource)

		routes.GET("get/sourceid/readlater", middlewares.CheckAccessToken(), route.handler.GetReadLaterListPaginationByArticlesSourceID)
		routes.GET("get/all/followed/readlater", middlewares.CheckAccessToken(), route.handler.GetReadLaterListPaginationByUserFollowedSource)

		routes.GET("get/all/recently", middlewares.CheckAccessToken(), route.handler.GetRecentlyReadArticle)

		routes.GET("get/all/treding", middlewares.CheckAccessToken(), route.handler.GetTredingArticle)

		routes.GET("list/all", middlewares.CheckAccessToken(), route.handler.ListAll)
		routes.POST("delete/id", middlewares.CheckAccessToken(), route.handler.Delete)

		routes.GET("search/across-source", middlewares.CheckAccessToken(), route.handler.SearchArticlesAcrossUserFollowedSources)
		routes.GET("search/admin", middlewares.CheckAccessToken(), route.handler.AdminSearchArticlesWithFilter)

		routes.GET("count/total", middlewares.CheckAccessToken(), route.handler.Count)
		routes.GET("count/previous-week", middlewares.CheckAccessToken(), route.handler.CountArticleCreateAWeekAgoByArticlesSourceID)

	}
}
