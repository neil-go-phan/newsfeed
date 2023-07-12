package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockFollowHandler struct{}

func (h *MockFollowHandler)Follow(c *gin.Context) {}
func (h *MockFollowHandler)Unfollow(c *gin.Context) {}
func (h *MockFollowHandler)GetArticleSourceFollowed(c *gin.Context) {}
func (h *MockFollowHandler)GetNewestSourceUpdatedID(c *gin.Context)  {}

func TestFollowSetup(t *testing.T) {
	r := gin.Default()

	MockFollowHandler := &MockFollowHandler{}

	testroutes := NewFollowRoutes(MockFollowHandler)

	testroutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 4)
}
