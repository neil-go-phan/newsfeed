package handlers

import (
	"net/http"
	"server/entities"
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
	GetMostActiveSources(c *gin.Context)
	SearchWithFilter(c *gin.Context)
	GetWithID(c *gin.Context)
	ListAll(c *gin.Context)
	ListAllPaging(c *gin.Context)
	Count(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type deleteArticleSourcePayload struct {
	ID uint `json:"id"`
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

func (h *ArticlesSourceHandler) GetMostActiveSources(c *gin.Context) {
	articlesSources, err := h.service.GetMostActiveSources()
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles_sources": articlesSources})
}

func (h *ArticlesSourceHandler) ListAll(c *gin.Context) {
	articlesSources, err := h.service.ListAll()
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles_sources": articlesSources})
}

func (h *ArticlesSourceHandler) GetWithID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	articlesSource, err := h.service.GetWithID(uint(id))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles_source": articlesSource})
}

func (h *ArticlesSourceHandler) Count(c *gin.Context) {
	total, err := h.service.Count()
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "total": total})
}

func (h *ArticlesSourceHandler) ListAllPaging(c *gin.Context) {
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

	articlesSources, err := h.service.ListAllPaging(page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles_sources": articlesSources})
}

func (h *ArticlesSourceHandler) SearchWithFilter(c *gin.Context) {
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
	topicID, err := strconv.Atoi(c.Query("topic_id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	articlesSources, found, err := h.service.SearchWithFilter(keyword, page, pageSize, uint(topicID))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles_sources": articlesSources, "found": found})
}

func (h *ArticlesSourceHandler) Delete(c *gin.Context) {
	var payload deleteArticleSourcePayload
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

	err = h.service.Delete(role.(string), payload.ID)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete success"})
}

func (h *ArticlesSourceHandler) Update(c *gin.Context) {
	var payload entities.ArticlesSource
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
	err = h.service.Update(role.(string),payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update success"})
}
