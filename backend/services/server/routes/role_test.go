package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockRoleHandler struct{}

func (h *MockRoleHandler)Get(c *gin.Context){}
func (h *MockRoleHandler)List(c *gin.Context){}
func (h *MockRoleHandler)Total(c *gin.Context){}
func (h *MockRoleHandler)Create(c *gin.Context){}
func (h *MockRoleHandler)Delete(c *gin.Context){}
func (h *MockRoleHandler)Update(c *gin.Context){}
func (h *MockRoleHandler)ListNames(c *gin.Context) {}

func TestRoleSetup(t *testing.T) {
	r := gin.Default()

	MockRoleHandler := &MockRoleHandler{}

	testroutes := NewRoleRoutes(MockRoleHandler)

	testroutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 7)
}
