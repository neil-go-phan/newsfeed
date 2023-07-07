package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"server/entities"
	"server/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CrawlerHandler struct {
	service services.CrawlerServices
}

type CrawlerHandlerInterface interface {
	TestRSSCrawler(c *gin.Context)
	TestCustomCrawler(c *gin.Context)
	CreateCrawler(c *gin.Context)
	GetHtmlPage(c *gin.Context)
	CreateCrawlerCronjobFromDB() error
	ListAllPaging(c *gin.Context)
	UpdateSchedule(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	GetCronjobOnHour(c *gin.Context) 
	GetCronjobOnDay(c *gin.Context)
}

type updateSchedulePayload struct {
	ID       uint   `json:"id"`
	Schedule string `json:"schedule"`
}

type updatePayload struct {
	ID      uint          `json:"id"`
	Crawler crawlerUpdate `json:"crawler"`
}

type crawlerUpdate struct {
	FeedLink           string `json:"feed_link"`
	CrawlType          string `json:"crawl_type" validate:"required"`
	ArticleDiv         string `json:"article_div"`
	ArticleTitle       string `json:"article_title"`
	ArticleDescription string `json:"article_description"`
	ArticleLink        string `json:"article_link"`
	ArticleAuthors     string `json:"article_authors"`
}

type chartDayResponse struct {
	Hour        int               `json:"hour"`
	AmountOfJob int               `json:"amount_of_jobs"`
	Cronjobs    []cronjobRunTimes `json:"cronjobs"`
}

type cronjobRunTimes struct {
	Name  string `json:"name"`
	Times int    `json:"times"`
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
	role, exsit := c.Get("role")
	if !exsit {
		log.Error("Not found role in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	err = h.service.CreateCrawlerWithCorrespondingArticlesSource(role.(string),payload)
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

func (h *CrawlerHandler) CreateCrawlerCronjobFromDB() error {
	return h.service.CreateCrawlerCronjobFromDB()
}

func (h *CrawlerHandler) ListAllPaging(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	crawlers, found, err := h.service.ListAllPaging(page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "crawlers": crawlers, "found": found})
}

func (h *CrawlerHandler) UpdateSchedule(c *gin.Context) {
	var payload updateSchedulePayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}
	role, exsit := c.Get("role")
	if !exsit {
		log.Error("Not found role in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	err = h.service.UpdateSchedule(role.(string),payload.ID, payload.Schedule)
	if err != nil {
		log.Error("error occrus:", err)
		if strings.Contains(err.Error(), "validate") {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update schedule success"})
}

func (h *CrawlerHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	crawler, err := h.service.Get(uint(id))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "crawler": crawler})
}

func (h *CrawlerHandler) Update(c *gin.Context) {
	var payload updatePayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	crawler := entities.Crawler{
		Model:              gorm.Model{ID: payload.ID},
		FeedLink:           payload.Crawler.FeedLink,
		CrawlType:          payload.Crawler.CrawlType,
		ArticleDiv:         payload.Crawler.ArticleDiv,
		ArticleTitle:       payload.Crawler.ArticleTitle,
		ArticleDescription: payload.Crawler.ArticleDescription,
		ArticleLink:        payload.Crawler.ArticleLink,
		ArticleAuthors:     payload.Crawler.ArticleAuthors,
	}
	role, exsit := c.Get("role")
	if !exsit {
		log.Error("Not found role in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	err = h.service.Update(role.(string),crawler)
	if err != nil {
		log.Error("error occrus:", err)
		if strings.Contains(err.Error(), "validate") {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update success"})
}

func (h *CrawlerHandler) GetCronjobOnHour(c *gin.Context) {
	time := c.Query("time")
	cronjobs, err := h.service.CronjobOnHour(time)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Get cronjob failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "cronjobs": cronjobs})
}

func (h *CrawlerHandler) GetCronjobOnDay(c *gin.Context) {
	time := c.Query("time")
	cronjobs, err := h.service.CronjobOnDay(time)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Get Cronjob failed"})
		return
	}

	reponse := parseChartDayToResponse(cronjobs)

	c.JSON(http.StatusOK, gin.H{"success": true, "cronjobs": reponse})
}

func parseChartDayToResponse(cronjobs *[24]services.ChartDay) [24]chartDayResponse {
	response := [24]chartDayResponse{}
	for index, cronjob := range cronjobs {
		response[index].Hour = cronjob.Hour
		response[index].AmountOfJob = cronjob.AmountOfJob
		cronjobResponese := make([]cronjobRunTimes, 0)
		for key, value := range cronjob.Cronjobs {
			cronjobResponese = append(cronjobResponese, cronjobRunTimes{
				Name:  key,
				Times: value,
			})
		}
		response[index].Cronjobs = cronjobResponese
	}
	return response
}
