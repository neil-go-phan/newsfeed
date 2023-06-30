package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
		route.POST("test/rss", middlewares.CheckAccessToken(), crawlerRoutes.handler.TestRSSCrawler)
		route.POST("test/custom", middlewares.CheckAccessToken(), crawlerRoutes.handler.TestCustomCrawler)
		route.POST("create", middlewares.CheckAccessToken(), crawlerRoutes.handler.CreateCrawler)
		route.GET("get/html-page", middlewares.CheckAccessToken(), crawlerRoutes.handler.GetHtmlPage)
	}
}

func (crawlerRoutes *CrawlerRoutes) CreateCrawlerCronjobFromDB() {
	err := crawlerRoutes.handler.CreateCrawlerCronjobFromDB()
	if err != nil {
		log.Error(err)
	} else {
		log.Println("create crawler cronjob from database success")
	}
}
