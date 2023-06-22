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
	GetByTopicID(c *gin.Context)
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

func (h *ArticlesSourceHandler) GetByTopicID(c *gin.Context) {
	topicID, err := strconv.Atoi(c.Query("topic_id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	articlesSources, err := h.service.GetByTopicID(uint(topicID))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles_sources": articlesSources})
}
