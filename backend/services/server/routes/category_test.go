package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockCategoryHandler struct{}

func (h *MockCategoryHandler) ListName(c *gin.Context)      {}
func (h *MockCategoryHandler) ListAll(c *gin.Context)       {}
func (h *MockCategoryHandler) GetPagination(c *gin.Context) {}
func (h *MockCategoryHandler) Count(c *gin.Context)         {}
func (h *MockCategoryHandler) Create(c *gin.Context)        {}
func (h *MockCategoryHandler) Delete(c *gin.Context)        {}
func (h *MockCategoryHandler) Update(c *gin.Context)        {}

func TestCategorySetup(t *testing.T) {
	r := gin.Default()

	MockCategoryHandler := &MockCategoryHandler{}

	testroutes := NewCategoryRoutes(MockCategoryHandler)

	testroutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 7)
}
