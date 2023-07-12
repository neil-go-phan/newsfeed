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

func TestReadReadArticle_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	handler.Read(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestReadReadArticle_PayloadInvalid(t *testing.T) {
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

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	handler.Read(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestReadReadArticle_ServiceError(t *testing.T) {
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

	payload := readAndUnreadPayload{
		ArticlesSourceID: 1,
		ArticleID: 10,
	}

	mockJsonPost(c, payload)

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	articleID := uint(10)
	articleSourceID := uint(1)

	mockService.On("MarkArticleAsRead", TEST_USERNAME, articleID, articleSourceID).Return(fmt.Errorf("service failed"))

	handler.Read(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestReadUnreadReadArticle_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	handler.Unread(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestReadUnreadArticle_PayloadInvalid(t *testing.T) {
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

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	handler.Unread(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestReadUnreadArticle_ServiceError(t *testing.T) {
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

	payload := readAndUnreadPayload{
		ArticlesSourceID: 1,
		ArticleID: 10,
	}

	mockJsonPost(c, payload)

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	articleID := uint(10)
	articleSourceID := uint(1)

	mockService.On("MarkArticleAsUnRead", TEST_USERNAME, articleID, articleSourceID).Return(fmt.Errorf("service failed"))

	handler.Unread(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestReadMarkAllAsReadRead_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	handler.MarkAllAsRead(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestReadMarkAllAsRead_ServiceError(t *testing.T) {
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

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	mockService.On("MarkAllAsReadByUserFollowedSource", TEST_USERNAME).Return(fmt.Errorf("service failed"))

	handler.MarkAllAsRead(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestReadMarkAllAsReadBySourceID_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	handler.MarkAllAsReadBySourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestReadMarkAllAsReadBySourceID_PayloadInvalid(t *testing.T) {
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

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	handler.MarkAllAsReadBySourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestReadMarkAllAsReadBySourceID_ServiceError(t *testing.T) {
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

	payload := MarkAllAsReadBySourceIDPayload{
		ArticlesSourceID: 1,
	}

	mockJsonPost(c, payload)

	mockService := new(mocks.ReadServices)
	handler := NewReadHandler(mockService)
	assert := assert.New(t)

	articleSourceID := uint(1)

	mockService.On("MarkAllAsReadBySourceID", TEST_USERNAME, articleSourceID).Return(fmt.Errorf("service failed"))

	handler.MarkAllAsReadBySourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}