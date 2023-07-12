package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"server/handlers/mocks"
	"server/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPermissionList_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.PermissionServices)
	handler := NewPermissionHandler(mockService)
	assert := assert.New(t)

	permissions := []services.PermissionResponse{}

	mockService.On("List").Return(permissions, fmt.Errorf("service failed"))

	handler.List(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}