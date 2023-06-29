package handlers

import (
	"net/http"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ArticleHandler struct {
	service services.ArticleServices
}

type ArticleHandlerInterface interface {
	GetArticlesPaginationByArticlesSourceID(c *gin.Context)
	GetArticlesPaginationByUserFollowedSources(c *gin.Context)
	GetUnreadArticlesPaginationByArticlesSourceID(c *gin.Context) 
	GetUnreadArticlesByUserFollowedSource(c *gin.Context)

	SearchArticlesAcrossUserFollowedSources(c *gin.Context)
	CountArticleCreateAWeekAgoByArticlesSourceID(c *gin.Context)
	// ListAll(c *gin.Context)
	// GetPagination(c *gin.Context)
	// Count(c *gin.Context)

	// Create(c *gin.Context)
	// Update(c *gin.Context)
}

func NewArticlesHandler(service services.ArticleServices) *ArticleHandler {
	return &ArticleHandler{
		service: service,
	}
}

func (h *ArticleHandler) GetArticlesPaginationByArticlesSourceID(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	articlesSourceID, err := strconv.Atoi(c.Query("articles_source_id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	articles, err := h.service.GetArticlesPaginationByArticlesSourceID(username.(string), uint(articlesSourceID), page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles})
}

func (h *ArticleHandler) GetArticlesPaginationByUserFollowedSources(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	articles, err := h.service.GetArticlesPaginationByUserFollowedSource(username.(string), page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles})
}

func (h *ArticleHandler) GetUnreadArticlesPaginationByArticlesSourceID(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	articlesSourceID, err := strconv.Atoi(c.Query("articles_source_id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	articles, err := h.service.GetUnreadArticlesPaginationByArticlesSourceID(username.(string), uint(articlesSourceID), page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles})
}

func (h *ArticleHandler) GetUnreadArticlesByUserFollowedSource(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	articles, err := h.service.GetUnreadArticlesByUserFollowedSource(username.(string), page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles})
}

func (h *ArticleHandler) CountArticleCreateAWeekAgoByArticlesSourceID(c *gin.Context) {
	articlesSourceID, err := strconv.Atoi(c.Query("articles_source_id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	count, err := h.service.CountArticleCreateAWeekAgoByArticlesSourceID(uint(articlesSourceID))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": count})
}

func (h *ArticleHandler) SearchArticlesAcrossUserFollowedSources(c *gin.Context) {
	keyword := c.Query("q")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	articles, found, err := h.service.SearchArticlesAcrossUserFollowedSources(username.(string), keyword, page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles, "found": found})
}
