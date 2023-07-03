package handlers

import (
	"net/http"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type FollowHandler struct {
	service services.FollowServices
}

type FollowHandlerInterface interface {
	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
	GetArticleSourceFollowed(c *gin.Context)
	GetNewestSourceUpdatedID(c *gin.Context) 
}

func NewFollowHandler(service services.FollowServices) *FollowHandler {
	return &FollowHandler{
		service: service,
	}
}

func (h *FollowHandler) Follow(c *gin.Context) {
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	articlesSourceID, err := strconv.Atoi(c.Query("articles_source_id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	err = h.service.Follow(username.(string), uint(articlesSourceID))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "follow success"})
}

func (h *FollowHandler) Unfollow(c *gin.Context) {
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	articlesSourceID, err := strconv.Atoi(c.Query("articles_source_id"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	err = h.service.Unfollow(username.(string), uint(articlesSourceID))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "unfollow success"})
}

func (h *FollowHandler) GetArticleSourceFollowed(c *gin.Context) {
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	articlesSourcesFollowed, err := h.service.GetUserFollowedSources(username.(string))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles_sources": articlesSourcesFollowed})
}

func (h *FollowHandler) GetNewestSourceUpdatedID(c *gin.Context) {
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}

	ids, err := h.service.GetNewestSourceUpdatedID(username.(string))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles_source_ids": ids})
}