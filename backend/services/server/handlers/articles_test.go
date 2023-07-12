package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"server/handlers/mocks"
	"server/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const BAD_REQUEST_RESPONSE = `{"success":false,"message":"bad request"}`
const INTERNAL_SERVER_ERROR_RESPONSE = `{"success":false,"message":"internal server error"}`
const INVALID_INPUT_RESPONSE = `{"success":false,"message":"input invalid"}`
const SERVICE_FAIL_RESPONSE = `{"success":false,"message":"service failed"}`
const TEST_USERNAME = "testuser"

func TestGetArticlesPaginationByArticlesSourceID_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetArticlesPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetArticlesPaginationByArticlesSourceID_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetArticlesPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetArticlesPaginationByArticlesSourceID_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetArticlesPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetArticlesPaginationByArticlesSourceID_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetArticlesPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetArticlesPaginationByArticlesSourceID_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", "testuser")

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	sourceID := uint(1)
	page := 1
	pageSize := 10

	articles := []services.ArticleForReadResponse{}

	mockArticleServices.On("GetArticlesPaginationByArticlesSourceID", TEST_USERNAME, sourceID, page, pageSize).Return(articles, fmt.Errorf("service failed"))

	articleHandler.GetArticlesPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetArticlesPaginationByUserFollowedSources_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetArticlesPaginationByUserFollowedSources(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetArticlesPaginationByUserFollowedSources_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetArticlesPaginationByUserFollowedSources(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetArticlesPaginationByUserFollowedSources_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetArticlesPaginationByUserFollowedSources(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetArticlesPaginationByUserFollowedSources_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetArticlesPaginationByUserFollowedSources(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetArticlesPaginationByUserFollowedSources_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", "testuser")

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	page := 1
	pageSize := 10

	articles := []services.ArticleForReadResponse{}

	mockArticleServices.On("GetArticlesPaginationByUserFollowedSource", TEST_USERNAME, page, pageSize).Return(articles, fmt.Errorf("service failed"))

	articleHandler.GetArticlesPaginationByUserFollowedSources(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetUnreadArticlesPaginationByArticlesSourceID_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetUnreadArticlesPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetUnreadArticlesPaginationByArticlesSourceID_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetUnreadArticlesPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetUnreadArticlesPaginationByArticlesSourceID_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetUnreadArticlesPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetUnreadArticlesPaginationByArticlesSourceID_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetUnreadArticlesPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetUnreadArticlesPaginationByArticlesSourceID_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", "testuser")

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	sourceID := uint(1)
	page := 1
	pageSize := 10

	articles := []services.ArticleForReadResponse{}

	mockArticleServices.On("GetUnreadArticlesPaginationByArticlesSourceID", TEST_USERNAME, sourceID, page, pageSize).Return(articles, fmt.Errorf("service failed"))

	articleHandler.GetUnreadArticlesPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetUnreadArticlesByUserFollowedSource_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetUnreadArticlesByUserFollowedSource(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetUnreadArticlesByUserFollowedSource_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetUnreadArticlesByUserFollowedSource(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetUnreadArticlesByUserFollowedSource_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetUnreadArticlesByUserFollowedSource(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetUnreadArticlesByUserFollowedSource_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetUnreadArticlesByUserFollowedSource(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetUnreadArticlesByUserFollowedSource_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", "testuser")

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	page := 1
	pageSize := 10

	articles := []services.ArticleForReadResponse{}

	mockArticleServices.On("GetUnreadArticlesByUserFollowedSource", TEST_USERNAME, page, pageSize).Return(articles, fmt.Errorf("service failed"))

	articleHandler.GetUnreadArticlesByUserFollowedSource(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetReadLaterListPaginationByArticlesSourceID_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetReadLaterListPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetReadLaterListPaginationByArticlesSourceID_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetReadLaterListPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetReadLaterListPaginationByArticlesSourceID_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetReadLaterListPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetReadLaterListPaginationByArticlesSourceID_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetReadLaterListPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetReadLaterListPaginationByArticlesSourceID_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("articles_source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", "testuser")

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	sourceID := uint(1)
	page := 1
	pageSize := 10

	articles := []services.ArticleForReadResponse{}

	mockArticleServices.On("GetReadLaterListPaginationByArticlesSourceID", TEST_USERNAME, sourceID, page, pageSize).Return(articles, fmt.Errorf("service failed"))

	articleHandler.GetReadLaterListPaginationByArticlesSourceID(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetReadLaterListPaginationByUserFollowedSource_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetReadLaterListPaginationByUserFollowedSource(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetReadLaterListPaginationByUserFollowedSource_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetReadLaterListPaginationByUserFollowedSource(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetReadLaterListPaginationByUserFollowedSource_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetReadLaterListPaginationByUserFollowedSource(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetReadLaterListPaginationByUserFollowedSource_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetReadLaterListPaginationByUserFollowedSource(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetReadLaterListPaginationByUserFollowedSource_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", "testuser")

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	page := 1
	pageSize := 10

	articles := []services.ArticleForReadResponse{}

	mockArticleServices.On("GetReadLaterListPaginationByUserFollowedSource", TEST_USERNAME, page, pageSize).Return(articles, fmt.Errorf("service failed"))

	articleHandler.GetReadLaterListPaginationByUserFollowedSource(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetRecentlyReadArticle_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetRecentlyReadArticle(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetRecentlyReadArticle_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetRecentlyReadArticle(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetRecentlyReadArticle_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetRecentlyReadArticle(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetRecentlyReadArticle_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetRecentlyReadArticle(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetRecentlyReadArticle_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", "testuser")

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	page := 1
	pageSize := 10

	articles := []services.ArticleForReadResponse{}

	mockArticleServices.On("GetRecentlyReadArticle", TEST_USERNAME, page, pageSize).Return(articles, fmt.Errorf("service failed"))

	articleHandler.GetRecentlyReadArticle(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCountArticleCreateAWeekAgoByArticlesSourceID_InvalidSourceID(t *testing.T) {
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

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.CountArticleCreateAWeekAgoByArticlesSourceID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestCountArticleCreateAWeekAgoByArticlesSourceID_ServiceError(t *testing.T) {
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
	c.Set("username", "testuser")

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	sourceID := uint(1)

	mockArticleServices.On("CountArticleCreateAWeekAgoByArticlesSourceID", sourceID).Return(int64(1), fmt.Errorf("service failed"))

	articleHandler.CountArticleCreateAWeekAgoByArticlesSourceID(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestSearchArticlesAcrossUserFollowedSources_NotFoundUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("q", "keyword")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.SearchArticlesAcrossUserFollowedSources(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestSearchArticlesAcrossUserFollowedSources_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")
	q.Add("q", "keyword")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.SearchArticlesAcrossUserFollowedSources(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestSearchArticlesAcrossUserFollowedSources_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")
	q.Add("q", "keyword")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.SearchArticlesAcrossUserFollowedSources(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestSearchArticlesAcrossUserFollowedSources_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("q", "keyword")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.SearchArticlesAcrossUserFollowedSources(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestSearchArticlesAcrossUserFollowedSources_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("q", "keyword")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("username", TEST_USERNAME)
	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	page := 1
	pageSize := 10
	keyword := "keyword"

	articles := []services.ArticleResponse{}

	mockArticleServices.On("SearchArticlesAcrossUserFollowedSources", TEST_USERNAME, keyword, page, pageSize).Return(articles, int64(1), fmt.Errorf("service failed"))

	articleHandler.SearchArticlesAcrossUserFollowedSources(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetTredingArticle_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.GetTredingArticle(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetTredingArticle_ServiceError(t *testing.T) {
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
	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articles := []services.TredingArticleResponse{}

	mockArticleServices.On("GetTredingArticle", TEST_USERNAME).Return(articles, fmt.Errorf("service failed"))

	articleHandler.GetTredingArticle(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestListAll_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.ListAll(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestListAll_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.ListAll(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestListAll_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	page := 1
	pageSize := 10

	articles := []services.ArticleResponse{}

	mockArticleServices.On("ListAll", page, pageSize).Return(articles, fmt.Errorf("service failed"))

	articleHandler.ListAll(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCount_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	mockArticleServices.On("Count").Return(1, fmt.Errorf("service failed"))

	articleHandler.Count(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestDelete_PayloadInvalid(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	invalidJson := "invalid"

	mockJsonPost(c, invalidJson)

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestDelete_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	invalidJson := map[string]interface{}{"id": 1}

	mockJsonPost(c, invalidJson)

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	mockArticleServices.On("Delete", uint(1)).Return(fmt.Errorf("service failed"))

	articleHandler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'input invalid' message to be returned")
}

func TestAdminSearchArticlesWithFilter_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")
	q.Add("source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.AdminSearchArticlesWithFilter(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestAdminSearchArticlesWithFilter_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")
	q.Add("source_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.AdminSearchArticlesWithFilter(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestAdminSearchArticlesWithFilter_InvalidSourceID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("source_id", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.AdminSearchArticlesWithFilter(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestAdminSearchArticlesWithFilter_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("source_id", "1")
	q.Add("q", "keyword")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mocks.ArticleServices)
	articleHandler := NewArticlesHandler(mockArticleServices)
	assert := assert.New(t)

	keyword := "keyword"
	sourceID := uint(1)
	page := 1
	pageSize := 10

	articles := []services.ArticleResponse{}

	mockArticleServices.On("AdminSearchArticlesWithFilter", keyword, page, pageSize, sourceID).Return(articles, int64(1), fmt.Errorf("service failed"))

	articleHandler.AdminSearchArticlesWithFilter(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func mockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}
