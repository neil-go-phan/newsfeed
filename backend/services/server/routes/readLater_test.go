package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockReadLaterHandler struct{}

func (h *MockReadLaterHandler) AddToReadLaterList(c *gin.Context)      {}
func (h *MockReadLaterHandler) RemoveFromReadLaterList(c *gin.Context) {}

func TestReadLaterSetup(t *testing.T) {
	r := gin.Default()

	MockReadLaterHandler := &MockReadLaterHandler{}

	testroutes := NewReadLaterRoutes(MockReadLaterHandler)

	testroutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 2)
}
