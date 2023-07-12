package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockPermissionHandler struct{}

func (h *MockPermissionHandler) List(c *gin.Context) {}

func TestPermissionSetup(t *testing.T) {
	r := gin.Default()

	MockPermissionHandler := &MockPermissionHandler{}

	testroutes := NewPermissionRoutes(MockPermissionHandler)

	testroutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 1)
}
