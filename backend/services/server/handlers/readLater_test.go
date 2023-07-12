package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"server/handlers/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddToReadLaterList_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.ReadLaterServices)
	handler := NewReadLaterHandler(mockService)
	assert := assert.New(t)

	handler.AddToReadLaterList(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestAddToReadLaterList_PayloadInvalid(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", TEST_USERNAME)

	invalidJson := "invalid"

	mockJsonPost(c, invalidJson)

	mockService := new(mocks.ReadLaterServices)
	handler := NewReadLaterHandler(mockService)
	assert := assert.New(t)

	handler.AddToReadLaterList(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestAddToReadLaterList_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", TEST_USERNAME)

	payload := addAndRemoveItemInReadLaterListPayload{
		ArticleID: 10,
	}

	mockJsonPost(c, payload)

	mockService := new(mocks.ReadLaterServices)
	handler := NewReadLaterHandler(mockService)
	assert := assert.New(t)

	articleID := uint(10)

	mockService.On("AddToReadLaterList", TEST_USERNAME, articleID).Return(fmt.Errorf("service failed"))

	handler.AddToReadLaterList(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRemoveFromReadLaterList_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.ReadLaterServices)
	handler := NewReadLaterHandler(mockService)
	assert := assert.New(t)

	handler.RemoveFromReadLaterList(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRemoveFromReadLaterList_PayloadInvalid(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", TEST_USERNAME)

	invalidJson := "invalid"

	mockJsonPost(c, invalidJson)

	mockService := new(mocks.ReadLaterServices)
	handler := NewReadLaterHandler(mockService)
	assert := assert.New(t)

	handler.RemoveFromReadLaterList(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRemoveFromReadLaterList_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", TEST_USERNAME)

	payload := addAndRemoveItemInReadLaterListPayload{
		ArticleID: 10,
	}

	mockJsonPost(c, payload)

	mockService := new(mocks.ReadLaterServices)
	handler := NewReadLaterHandler(mockService)
	assert := assert.New(t)

	articleID := uint(10)

	mockService.On("RemoveFromReadLaterList", TEST_USERNAME, articleID).Return(fmt.Errorf("service failed"))

	handler.RemoveFromReadLaterList(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}