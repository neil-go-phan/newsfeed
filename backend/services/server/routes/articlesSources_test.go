package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockArticlesSourceHandler struct{}

func (h *MockArticlesSourceHandler)GetByTopicIDPaginate(c *gin.Context) {}
func (h *MockArticlesSourceHandler)GetMostActiveSources(c *gin.Context){}
func (h *MockArticlesSourceHandler)SearchWithFilter(c *gin.Context){}
func (h *MockArticlesSourceHandler)GetWithID(c *gin.Context){}
func (h *MockArticlesSourceHandler)ListAll(c *gin.Context){}
func (h *MockArticlesSourceHandler)ListAllPaging(c *gin.Context){}
func (h *MockArticlesSourceHandler)Count(c *gin.Context){}
func (h *MockArticlesSourceHandler)Delete(c *gin.Context){}
func (h *MockArticlesSourceHandler)Update(c *gin.Context){}
func TestArticlesSourcesSetup(t *testing.T) {
	r := gin.Default()

	mockArticlesSourceHandler := &MockArticlesSourceHandler{}

	articleRoutes := NewArticlesSourceRoutes(mockArticlesSourceHandler)

	articleRoutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 9)
}