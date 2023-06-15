package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"server/entities"
	"server/services"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CrawlerHandler struct {
	service services.CrawlerServices
}

type CrawlerHandlerInterface interface {
	TestRSSCrawler(c *gin.Context)
	TestCustomCrawler(c *gin.Context)
	CreateCrawler(c *gin.Context)
	GetHtmlPage(c *gin.Context)
}

func NewCrawlerHandler(service services.CrawlerServices) *CrawlerHandler {
	return &CrawlerHandler{
		service: service,
	}
}

func (h *CrawlerHandler) TestRSSCrawler(c *gin.Context) {
	var inputCrawler entities.Crawler
	err := c.BindJSON(&inputCrawler)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	articlesSource, articles, err := h.service.TestRSSCrawler(inputCrawler)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "articles_source": articlesSource, "articles": articles})
}

func (h *CrawlerHandler) TestCustomCrawler(c *gin.Context) {
	var inputCrawler entities.Crawler
	err := c.BindJSON(&inputCrawler)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	articlesSource, articles, err := h.service.TestCustomCrawler(inputCrawler)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "articles_source": articlesSource, "articles": articles})
}

func (h *CrawlerHandler) CreateCrawler(c *gin.Context) {
	var payload services.CreateCrawlerPayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.CreateCrawlerWithCorrespondingArticlesSource(payload)
	if err != nil {
		log.Error("error occrus:", err)
		if strings.Contains(err.Error(), "validate") {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "create success"})
}

func (h *CrawlerHandler) GetHtmlPage(c *gin.Context) {
	urlInput := c.Query("url")
	url, err := url.Parse(urlInput)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Bad request"})
		return
	}

	err = h.service.GetHtmlPage(url)
	if err != nil {
		log.Errorln("error occurs when response page content to frontend: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal server error"})
		return
	}
	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	filePath := fmt.Sprintf("page%s.html", hostname)

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	c.Header("Content-Type", "text/html")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = os.Remove(filePath)
	if err != nil {
		log.Errorln("error occurs when delete html file: ", err)
	}
}
