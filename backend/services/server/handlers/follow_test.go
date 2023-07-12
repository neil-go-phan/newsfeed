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

func TestFollowFollowSource_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.FollowServices)
	handler := NewFollowHandler(mockService)
	assert := assert.New(t)

	handler.Follow(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestFollowFollowSource_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	q.Add("articles_source_id", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", TEST_USERNAME)

	mockService := new(mocks.FollowServices)
	handler := NewFollowHandler(mockService)
	assert := assert.New(t)

	handler.Follow(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestFollowFollowSource_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", TEST_USERNAME)

	mockService := new(mocks.FollowServices)
	handler := NewFollowHandler(mockService)
	assert := assert.New(t)

	sourceID := uint(1)

	mockService.On("Follow", TEST_USERNAME, sourceID).Return(fmt.Errorf("service failed"))

	handler.Follow(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestFollowUnfollowSource_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.FollowServices)
	handler := NewFollowHandler(mockService)
	assert := assert.New(t)

	handler.Unfollow(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestFollowUnfollowSource_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	q.Add("articles_source_id", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", TEST_USERNAME)

	mockService := new(mocks.FollowServices)
	handler := NewFollowHandler(mockService)
	assert := assert.New(t)

	handler.Unfollow(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestFollowUnfollowSource_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", TEST_USERNAME)

	mockService := new(mocks.FollowServices)
	handler := NewFollowHandler(mockService)
	assert := assert.New(t)

	sourceID := uint(1)

	mockService.On("Unfollow", TEST_USERNAME, sourceID).Return(fmt.Errorf("service failed"))

	handler.Unfollow(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestFollowGetArticleSourceFollowed_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.FollowServices)
	handler := NewFollowHandler(mockService)
	assert := assert.New(t)

	handler.GetArticleSourceFollowed(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestFollowGetArticleSourceFollowed_ServiceError(t *testing.T) {
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

	mockService := new(mocks.FollowServices)
	handler := NewFollowHandler(mockService)
	assert := assert.New(t)

	articlesSourcesFollowed := []services.ArticlesSourceUserFollow{}

	mockService.On("GetUserFollowedSources", TEST_USERNAME).Return(articlesSourcesFollowed, fmt.Errorf("service failed"))

	handler.GetArticleSourceFollowed(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestFollowGetNewestSourceUpdatedID_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.FollowServices)
	handler := NewFollowHandler(mockService)
	assert := assert.New(t)

	handler.GetNewestSourceUpdatedID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestFollowGetNewestSourceUpdatedID_ServiceError(t *testing.T) {
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

	mockService := new(mocks.FollowServices)
	handler := NewFollowHandler(mockService)
	assert := assert.New(t)

	ids := []uint{}

	mockService.On("GetNewestSourceUpdatedID", TEST_USERNAME).Return(ids, fmt.Errorf("service failed"))

	handler.GetNewestSourceUpdatedID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}
