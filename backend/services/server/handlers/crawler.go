package handlers

import (
	"net/http"
	"server/entities"
	"server/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CrawlerHandler struct {
	service services.CrawlerServices
}

type CrawlerHandlerInterface interface {
	TestCrawler(c *gin.Context)
}

func NewCrawlerHandler(service services.CrawlerServices) *CrawlerHandler {
	return &CrawlerHandler{
		service: service,
	}
}

func (h *CrawlerHandler) TestCrawler(c *gin.Context) {
	var inputCrawler entities.Crawler
	err := c.BindJSON(&inputCrawler)
	if err != nil {
		log.Error("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	articlesSource, articles, err := h.service.TestCrawler(&inputCrawler)
	if err != nil {
		log.Error("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "articles_source": articlesSource, "articles": articles})
}
