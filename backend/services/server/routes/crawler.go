package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type CrawlerRoutes struct {
	handler handlers.CrawlerHandlerInterface
}

func NewCrawlerRoutes(handler handlers.CrawlerHandlerInterface) *CrawlerRoutes {
	return &CrawlerRoutes{
		handler: handler,
	}
}

func (crawlerRoutes *CrawlerRoutes) Setup(r *gin.Engine) {
	route := r.Group("crawler")
	{
		route.POST("test", middlewares.CheckAccessToken(), crawlerRoutes.handler.TestCrawler)
	}
}
