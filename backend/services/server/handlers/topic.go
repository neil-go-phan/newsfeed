package handlers

import (
	"net/http"
	"server/entities"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TopicHandler struct {
	service services.TopicServices
}

type TopicHandlerInterface interface {
	List(c *gin.Context)
	GetPagination(c *gin.Context)
	GetByCategory(c *gin.Context)
	SearchTopicAndArticlesSource(c *gin.Context)
	Count(c *gin.Context)

	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func NewTopicHandler(service services.TopicServices) *TopicHandler {
	return &TopicHandler{
		service: service,
	}
}

func (h *TopicHandler) Create(c *gin.Context) {
	var topic entities.Topic
	err := c.BindJSON(&topic)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.CreateIfNotExist(topic)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "create success"})
}

func (h *TopicHandler) List(c *gin.Context) {
	topics, err := h.service.List()
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "topics": topics})
}

func (h *TopicHandler) Delete(c *gin.Context) {
	var topic entities.Topic
	err := c.BindJSON(&topic)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.Delete(topic)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete success"})
}

func (h *TopicHandler) Update(c *gin.Context) {
	var topic entities.Topic
	err := c.BindJSON(&topic)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.Update(topic)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update success"})
}

func (h *TopicHandler) GetPagination(c *gin.Context) {
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

	topics, err := h.service.GetPagination(page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "topics": topics})
}

func (h *TopicHandler) Count(c *gin.Context) {
	total, err := h.service.Count()
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "total": total})
}

func (h *TopicHandler) GetByCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	

	topics, err := h.service.GetByCategory(uint(categoryID))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "topics": topics})
}

func (h *TopicHandler) SearchTopicAndArticlesSource(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	topics, articlesSources, found, err := h.service.SearchTopicAndArticlesSourcePaginate(keyword, page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "topics": topics, "articles_sources": articlesSources, "found": found})
}
