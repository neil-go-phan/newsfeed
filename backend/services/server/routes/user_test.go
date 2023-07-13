package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockUserHandler struct{}

func (h *MockUserHandler) CheckAuth(c *gin.Context)       {}
func (h *MockUserHandler) Token(c *gin.Context)           {}
func (h *MockUserHandler) Register(c *gin.Context)        {}
func (h *MockUserHandler) Login(c *gin.Context)           {}
func (h *MockUserHandler) GoogleOAuth(c *gin.Context)     {}
func (h *MockUserHandler) AccessAdminPage(c *gin.Context) {}
func (h *MockUserHandler) ChangeRole(c *gin.Context)      {}
func (h *MockUserHandler) List(c *gin.Context)            {}
func (h *MockUserHandler) Delete(c *gin.Context)          {}
func (h *MockUserHandler) Total(c *gin.Context)           {}
func (h *MockUserHandler) UserUpgrateRole(c *gin.Context) {}

func TestUserSetup(t *testing.T) {
	r := gin.Default()

	MockUserHandler := &MockUserHandler{}

	testroutes := NewUserRoutes(MockUserHandler)

	testroutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 11)
}
