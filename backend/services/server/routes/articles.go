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
		routes.GET("get-page-by-articles-source-id", middlewares.CheckAccessToken(), route.handler.GetArticlesPaginationByArticlesSourceID)
		routes.GET("get-page-by-all-user-followed-sources", middlewares.CheckAccessToken(), route.handler.GetArticlesPaginationByUserFollowedSources)

		routes.GET("get-page-by-articles-source-id-unread", middlewares.CheckAccessToken(), route.handler.GetUnreadArticlesPaginationByArticlesSourceID)
		routes.GET("get-page-by-all-user-followed-sources-unread", middlewares.CheckAccessToken(), route.handler.GetUnreadArticlesByUserFollowedSource)

		routes.GET("get-page-by-articles-source-id-readlater", middlewares.CheckAccessToken(), route.handler.GetReadLaterListPaginationByArticlesSourceID)
		routes.GET("get-page-by-all-user-followed-sources-readlater", middlewares.CheckAccessToken(), route.handler.GetReadLaterListPaginationByUserFollowedSource)

		routes.GET("search-articles-across-source", middlewares.CheckAccessToken(), route.handler.SearchArticlesAcrossUserFollowedSources)
		routes.GET("count-article-previous-week", middlewares.CheckAccessToken(), route.handler.CountArticleCreateAWeekAgoByArticlesSourceID)

	}
}
