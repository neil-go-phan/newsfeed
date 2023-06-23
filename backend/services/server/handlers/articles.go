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
	GetPaginationByArticlesSourceID(c *gin.Context)
	SearchArticlesAcrossSources(c *gin.Context)
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

func (h *ArticleHandler) GetPaginationByArticlesSourceID(c *gin.Context) {
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

	articles, err := h.service.GetPaginationByArticlesSourceID(uint(articlesSourceID), page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles})
}

func (h *ArticleHandler) SearchArticlesAcrossSources(c *gin.Context) {
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

	articles, found, err := h.service.SearchArticlesAcrossSources(keyword, page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles, "found": found})
}
