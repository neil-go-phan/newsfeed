package handlers

import (
	"net/http"
	"server/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ReadLaterHandler struct {
	service services.ReadLaterServices
}

//go:generate mockery --name ReadLaterHandlerInterface
type ReadLaterHandlerInterface interface {
	AddToReadLaterList(c *gin.Context)
	RemoveFromReadLaterList(c *gin.Context)
}

type addAndRemoveItemInReadLaterListPayload struct {
	ArticleID uint `json:"article_id"`
}

func NewReadLaterHandler(service services.ReadLaterServices) *ReadLaterHandler {
	return &ReadLaterHandler{
		service: service,
	}
}

func (h *ReadLaterHandler) AddToReadLaterList(c *gin.Context) {
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	var payload addAndRemoveItemInReadLaterListPayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.AddToReadLaterList(username.(string), payload.ArticleID)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "add articles to read later list success"})
}

func (h *ReadLaterHandler) RemoveFromReadLaterList(c *gin.Context) {
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	var payload addAndRemoveItemInReadLaterListPayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.RemoveFromReadLaterList(username.(string), payload.ArticleID)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "remove articles from read later list success"})
}
