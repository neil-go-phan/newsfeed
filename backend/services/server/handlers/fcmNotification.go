package handlers

import (
	"net/http"
	"server/entities"
	"server/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type FcmNotificationHandler struct {
	service services.FcmNotificationServices
}

//go:generate mockery --name FcmNotificationHandlerInterface
type FcmNotificationHandlerInterface interface {
	Create(c *gin.Context)
	CreateCronjob() 
}

func NewFcmNotificationHandler(service services.FcmNotificationServices) *FcmNotificationHandler {
	return &FcmNotificationHandler{
		service: service,
	}
}

func (h *FcmNotificationHandler) Create(c *gin.Context) {
	token := c.Query("token")
	username, exsit := c.Get("username")
	if !exsit {
		log.Error("Not found username in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	log.Println("token",token)
	entity := entities.FcmNotification{
		Username: username.(string),
		FirebaseToken: token,
	}

	err := h.service.Create(entity)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "create success"})
}

func (h *FcmNotificationHandler) CreateCronjob() {
	h.service.CronjobPushNotification()
}