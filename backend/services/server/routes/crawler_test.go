package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockCrawlerHandler struct{}

func (h *MockCrawlerHandler) TestRSSCrawler(c *gin.Context)     {}
func (h *MockCrawlerHandler) TestCustomCrawler(c *gin.Context)  {}
func (h *MockCrawlerHandler) CreateCrawler(c *gin.Context)      {}
func (h *MockCrawlerHandler) GetHtmlPage(c *gin.Context)        {}
func (h *MockCrawlerHandler) CreateCrawlerCronjobFromDB() error { return nil }
func (h *MockCrawlerHandler) ListAllPaging(c *gin.Context)      {}
func (h *MockCrawlerHandler) UpdateSchedule(c *gin.Context)     {}
func (h *MockCrawlerHandler) Update(c *gin.Context)             {}
func (h *MockCrawlerHandler) Get(c *gin.Context)                {}
func (h *MockCrawlerHandler) GetCronjobOnHour(c *gin.Context)   {}
func (h *MockCrawlerHandler) GetCronjobOnDay(c *gin.Context)    {}

func TestCrawlerSetup(t *testing.T) {
	r := gin.Default()

	MockCrawlerHandler := &MockCrawlerHandler{}

	testroutes := NewCrawlerRoutes(MockCrawlerHandler)

	testroutes.Setup(r)

	routes := r.Routes()
	assert.Len(t, routes, 10)
}
