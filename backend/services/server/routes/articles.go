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
		routes.GET("get-page-by-articles-source-id", middlewares.CheckAccessToken(), route.handler.GetPaginationByArticlesSourceID)
	}
}
