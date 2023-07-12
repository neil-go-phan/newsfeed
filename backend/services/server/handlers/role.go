package handlers

import (
	"net/http"
	"server/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type RoleHandler struct {
	service services.RoleServices
}

//go:generate mockery --name RoleHandlerInterface
type RoleHandlerInterface interface {
	Get(c *gin.Context)
	List(c *gin.Context)
	Total(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	ListNames(c *gin.Context) 
}

type deleteRolePayload struct {
	ID uint `json:"id"`
}

func NewRoleHandler(service services.RoleServices) *RoleHandler {
	return &RoleHandler{
		service: service,
	}
}

func (h *RoleHandler) Get(c *gin.Context) {
	role, exsit := c.Get("role")
	if !exsit {
		log.Error("Not found role in token string")
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	roleResponse, err := h.service.Get(role.(string))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "role": roleResponse})
}

func (h *RoleHandler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}
	roleResponse, err := h.service.List(page, pageSize)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "roles": roleResponse})
}

func (h *RoleHandler) Total(c *gin.Context) {
	total, err := h.service.Count()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "total": total})
}

func (h *RoleHandler) ListNames(c *gin.Context) {
	names, err := h.service.ListRoleName()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "names": names})
}

func (h *RoleHandler) Delete(c *gin.Context) {
	var payload deleteRolePayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}
	err = h.service.Delete(payload.ID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete success"})
}

func (h *RoleHandler) Create(c *gin.Context) {
	var payload services.RoleResponse
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.Create(payload)
	if err != nil {
		log.Error("error occrus:", err)
		if strings.Contains(err.Error(), "validate") {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "create success"})
}

func (h *RoleHandler) Update(c *gin.Context) {
	var payload services.RoleResponse
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.Update(payload)
	if err != nil {
		log.Error("error occrus:", err)
		if strings.Contains(err.Error(), "validate") {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update success"})
}
