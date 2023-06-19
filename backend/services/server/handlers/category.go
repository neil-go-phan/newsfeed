package handlers

import (
	"net/http"
	"server/entities"
	"server/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CategoryHandler struct {
	service services.CategoryServices
}

type CategoryHandlerInterface interface {
	List(c *gin.Context)
	GetPagination(c *gin.Context)
	Count(c *gin.Context)

	Create(c *gin.Context)
	Delete(c *gin.Context)
	UpdateName(c *gin.Context)
}

func NewCategoryHandler(service services.CategoryServices) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var category entities.Category
	err := c.BindJSON(&category)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.CreateIfNotExist(category)
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

func (h *CategoryHandler) List(c *gin.Context) {
	categories, err := h.service.List()
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "categories": categories})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	var category entities.Category
	err := c.BindJSON(&category)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.Delete(category)
	if err != nil {
		log.Error("error occrus:", err)
		if strings.Contains(err.Error(), "validate") {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete success"})
}

func (h *CategoryHandler) UpdateName(c *gin.Context) {
	var payload services.UpdateNameCategoryPayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = h.service.UpdateName(payload)
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

func (h *CategoryHandler) GetPagination(c *gin.Context) {
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

	categories, err := h.service.GetPagination(page, pageSize)
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "categories": categories})
}

func (h *CategoryHandler) Count(c *gin.Context) {
	total, err := h.service.Count()
	if err != nil {
		log.Error("error occrus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "total": total})
}
