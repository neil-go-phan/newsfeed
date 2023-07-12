package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockArticlesHandler struct{}

func (h *MockArticlesHandler) GetArticlesPaginationByArticlesSourceID(c *gin.Context) {}
func (h *MockArticlesHandler) GetArticlesPaginationByUserFollowedSources(c *gin.Context) {}
func (h *MockArticlesHandler) GetUnreadArticlesPaginationByArticlesSourceID(c *gin.Context) {}
func (h *MockArticlesHandler) GetUnreadArticlesByUserFollowedSource(c *gin.Context) {}
func (h *MockArticlesHandler) GetReadLaterListPaginationByArticlesSourceID(c *gin.Context) {}
func (h *MockArticlesHandler) GetReadLaterListPaginationByUserFollowedSource(c *gin.Context) {}
func (h *MockArticlesHandler) GetRecentlyReadArticle(c *gin.Context) {}
func (h *MockArticlesHandler) GetTredingArticle(c *gin.Context) {}
func (h *MockArticlesHandler) ListAll(c *gin.Context) {}
func (h *MockArticlesHandler) Delete(c *gin.Context) {}
func (h *MockArticlesHandler) SearchArticlesAcrossUserFollowedSources(c *gin.Context) {}
func (h *MockArticlesHandler) AdminSearchArticlesWithFilter(c *gin.Context) {}
func (h *MockArticlesHandler) Count(c *gin.Context) {}
func (h *MockArticlesHandler) CountArticleCreateAWeekAgoByArticlesSourceID(c *gin.Context) {}

func TestArticleSetup(t *testing.T) {
	r := gin.Default()

	mockArticlesHandler := &MockArticlesHandler{}

	articleRoutes := NewArticleRoutes(mockArticlesHandler)

	articleRoutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 14)
}