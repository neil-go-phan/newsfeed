package handlers

import (
	"net/http"
	"server/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type PermissionHandler struct {
	service services.PermissionServices
}

//go:generate mockery --name PermissionHandlerInterface
type PermissionHandlerInterface interface {
	List(c *gin.Context)
}

func NewPermissionHandler(service services.PermissionServices) *PermissionHandler {
	return &PermissionHandler{
		service: service,
	}
}

func (h *PermissionHandler) List(c *gin.Context) {
	permissions, err := h.service.List()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "permissions": permissions})
}
