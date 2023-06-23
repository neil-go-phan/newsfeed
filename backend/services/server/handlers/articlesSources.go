package handlers

import (
	"net/http"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ArticlesSourceHandler struct {
	service services.ArticlesSourceServices
}

type ArticlesSourceHandlerInterface interface {
	GetByTopicIDPaginate(c *gin.Context)
	// ListAll(c *gin.Context)
	// GetPagination(c *gin.Context)
	// Count(c *gin.Context)

	// Create(c *gin.Context)
	// Update(c *gin.Context)
}

func NewArticlesSourceHandler(service services.ArticlesSourceServices) *ArticlesSourceHandler {
	return &ArticlesSourceHandler{
		service: service,
	}
}

func (h *ArticlesSourceHandler) GetByTopicIDPaginate(c *gin.Context) {
	topicID, err := strconv.Atoi(c.Query("topic_id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	articlesSources, found, err := h.service.GetByTopicIDPaginate(uint(topicID), page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles_sources": articlesSources, "found": found})
}
