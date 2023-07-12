package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockTopicHandler struct{}

func (h *MockTopicHandler)List(c *gin.Context) {}
func (h *MockTopicHandler)GetPagination(c *gin.Context){}
func (h *MockTopicHandler)GetByCategory(c *gin.Context){}
func (h *MockTopicHandler)SearchTopicAndArticlesSource(c *gin.Context){}
func (h *MockTopicHandler)Count(c *gin.Context){}
func (h *MockTopicHandler)Create(c *gin.Context){}
func (h *MockTopicHandler)Delete(c *gin.Context){}
func (h *MockTopicHandler)Update(c *gin.Context){}

func TestTopicSetup(t *testing.T) {
	r := gin.Default()

	MockTopicHandler := &MockTopicHandler{}

	testroutes := NewTopicRoutes(MockTopicHandler)

	testroutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 8)
}
