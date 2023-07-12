package middlewares

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExpiredAccessTokenHandler(t *testing.T) {
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodGet, "/token", nil)
	c.Request.Header.Set("X-Refresh-Token", "expired-token")

	handler := ExpiredAccessTokenHandler()
	handler(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.True(t, c.IsAborted())
}