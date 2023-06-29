package handlers

import (
	"net/http"
	"server/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ReadHandler struct {
	service services.ReadServices
}

type ReadHandlerInterface interface {
	Read(c *gin.Context)
	Unread(c *gin.Context)
}

type readAndUnreadPayload struct {
	ArticlesSourceID uint `json:"articles_source_id"`
	ArticleID        uint `json:"article_id"`
}

func NewReadHandler(service services.ReadServices) *ReadHandler {
	return &ReadHandler{
		service: service,
	}
}

func (h *ReadHandler) Read(c *gin.Context) {
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	var readPayload readAndUnreadPayload
	err := c.BindJSON(&readPayload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.MarkArticleAsRead(username.(string), readPayload.ArticleID, readPayload.ArticlesSourceID)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "mark article as read success"})
}

func (h *ReadHandler) Unread(c *gin.Context) {
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	var unreadPayload readAndUnreadPayload
	err := c.BindJSON(&unreadPayload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.MarkArticleAsUnRead(username.(string), unreadPayload.ArticleID, unreadPayload.ArticlesSourceID)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "mark article as unread success"})
}