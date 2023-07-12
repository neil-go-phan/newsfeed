package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockReadHandler struct{}

func (h *MockReadHandler)Read(c *gin.Context) {}
func (h *MockReadHandler)Unread(c *gin.Context) {}
func (h *MockReadHandler)MarkAllAsRead(c *gin.Context) {}
func (h *MockReadHandler)MarkAllAsReadBySourceID(c *gin.Context) {}
func TestReadSetup(t *testing.T) {
	r := gin.Default()

	MockReadHandler := &MockReadHandler{}

	testroutes := NewReadRoutes(MockReadHandler)

	testroutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 4)
}
