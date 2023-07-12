package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"server/entities"
	"server/handlers/mocks"
	"server/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTopicCreate_InvalidPayload(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := "invalid"

	mockJsonPost(c, payload)

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.Create(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicCreate_NotfoundRole(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := entities.Topic{}

	mockJsonPost(c, payload)

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.Create(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicCreate_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("role", SUPER_ADMIN_ROLE)

	payload := entities.Topic{}

	mockJsonPost(c, payload)

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	mockService.On("CreateIfNotExist", SUPER_ADMIN_ROLE, payload).Return(fmt.Errorf("service failed"))

	handler.Create(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicList_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	topics := []services.TopicResponse{}

	mockService.On("List").Return(topics, fmt.Errorf("service failed"))

	handler.List(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicDelete_InvalidPayload(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := "invalid"

	mockJsonPost(c, payload)

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicDelete_NotfoundRole(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := entities.Topic{}

	mockJsonPost(c, payload)

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicDelete_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("role", SUPER_ADMIN_ROLE)

	payload := entities.Topic{}

	mockJsonPost(c, payload)

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	mockService.On("Delete", SUPER_ADMIN_ROLE, payload).Return(fmt.Errorf("service failed"))

	handler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicUpdate_InvalidPayload(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := "invalid"

	mockJsonPost(c, payload)

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicUpdate_NotfoundRole(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := entities.Topic{}

	mockJsonPost(c, payload)

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicUpdate_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()
	c.Set("role", SUPER_ADMIN_ROLE)

	payload := entities.Topic{}

	mockJsonPost(c, payload)

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	mockService.On("Update", SUPER_ADMIN_ROLE, payload).Return(fmt.Errorf("service failed"))

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicGetPagination_InvalidPage(t *testing.T) {
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

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.GetPagination(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicGetPagination_InvalidPageSize(t *testing.T) {
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

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.GetPagination(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicGetPagination_ServiceError(t *testing.T) {
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

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	page := 1
	pageSize := 10

	topics := []services.TopicResponse{}

	mockService.On("GetPagination", page, pageSize).Return(topics, fmt.Errorf("fail"))

	handler.GetPagination(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicCount_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	mockService.On("Count").Return(0, fmt.Errorf("service failed"))

	handler.Count(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicGetByCategory_InvalidCategoryID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("category_id", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.GetByCategory(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicGetByCategory_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("category_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	id := uint(1)

	topics := []services.TopicResponse{}

	mockService.On("GetByCategory", id).Return(topics, fmt.Errorf("fail"))

	handler.GetByCategory(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicSearchTopicAndArticlesSource_InvalidPage(t *testing.T) {
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

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.SearchTopicAndArticlesSource(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicSearchTopicAndArticlesSource_InvalidPageSize(t *testing.T) {
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

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	handler.SearchTopicAndArticlesSource(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestTopicSearchTopicAndArticlesSource_ServiceError(t *testing.T) {
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

	mockService := new(mocks.TopicServices)
	handler := NewTopicHandler(mockService)
	assert := assert.New(t)

	page := 1
	pageSize := 10
	keyword := "keyword"

	topics := []services.TopicResponse{}
	articlesSources := []services.ArticlesSourceResponseRender{}
	found := int64(1)

	mockService.On("SearchTopicAndArticlesSourcePaginate", keyword, page, pageSize).Return(topics, articlesSources, found, fmt.Errorf("fail"))

	handler.SearchTopicAndArticlesSource(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}
