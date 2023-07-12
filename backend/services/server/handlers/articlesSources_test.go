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

const SUPER_ADMIN_ROLE = "Superamin"

func TestGetByTopicIDPaginate_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")
	q.Add("topic_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.GetByTopicIDPaginate(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetByTopicIDPaginate_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")
	q.Add("topic_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.GetByTopicIDPaginate(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'invalid input' message to be returned")
}

func TestGetByTopicIDPaginate_InvalidTopicID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("topic_id", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.GetByTopicIDPaginate(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestGetByTopicIDPaginate_ServicesError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("topic_id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	topicID := uint(1)
	page := 1
	pageSize := 10

	articlesSources := []services.ArticlesSourceResponseRender{}

	mockArticlesSourceServices.On("GetByTopicIDPaginate", topicID, page, pageSize).Return(articlesSources, int64(1), fmt.Errorf("service failed"))

	articleSourceHandler.GetByTopicIDPaginate(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetMostActiveSources_ServicesError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articlesSources := []services.ArticlesSourceRecommended{}

	mockArticlesSourceServices.On("GetMostActiveSources").Return(articlesSources, fmt.Errorf("service failed"))

	articleSourceHandler.GetMostActiveSources(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestListAll_ServicesError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articlesSources := []services.ArticlesSourceResponseRender{}

	mockArticlesSourceServices.On("ListAll").Return(articlesSources, fmt.Errorf("service failed"))

	articleSourceHandler.ListAll(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetWithID_ServicesError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("id", "1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	id := uint(1)

	articlesSources := services.ArticlesSourceResponseRender{}

	mockArticlesSourceServices.On("GetWithID", id).Return(articlesSources, fmt.Errorf("service failed"))

	articleSourceHandler.GetWithID(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetWithID_InvalidID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("id", "invalid")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.GetWithID(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'bad request' message to be returned")
}

func TestCount_ServicesError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	mockArticlesSourceServices.On("Count").Return(int(1), fmt.Errorf("service failed"))

	articleSourceHandler.Count(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestListAllPaging_InvalidPage(t *testing.T) {
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

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.ListAllPaging(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestListAllPaging_InvalidPageSize(t *testing.T) {
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

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.ListAllPaging(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'invalid input' message to be returned")
}

func TestListAllPaging_ServicesError(t *testing.T) {
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

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	page := 1
	pageSize := 10

	articlesSources := []services.ArticlesSourceResponseRender{}

	mockArticlesSourceServices.On("ListAllPaging", page, pageSize).Return(articlesSources, fmt.Errorf("service failed"))

	articleSourceHandler.ListAllPaging(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestSearchWithFilter_InvalidPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "invalid")
	q.Add("page_size", "10")
	q.Add("topic_id", "1")
	q.Add("q", "keyword")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.SearchWithFilter(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestSearchWithFilter_InvalidPageSize(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "invalid")
	q.Add("topic_id", "1")
	q.Add("q", "keyword")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.SearchWithFilter(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'invalid input' message to be returned")
}

func TestSearchWithFilter_InvalidTopicID(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("topic_id", "invalid")
	q.Add("q", "keyword")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.SearchWithFilter(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestSearchWithFilter_ServicesError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("page_size", "10")
	q.Add("topic_id", "1")
	q.Add("q", "keyword")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	keyword := "keyword"
	topicID := uint(1)
	page := 1
	pageSize := 10

	articlesSources := []services.ArticlesSourceResponseRender{}

	mockArticlesSourceServices.On("SearchWithFilter",keyword, page, pageSize,topicID).Return(articlesSources, int64(1), fmt.Errorf("service failed"))

	articleSourceHandler.SearchWithFilter(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestDelete_ServicesError(t *testing.T) {
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

	invalidJson := map[string]interface{}{"id": 1}

	mockJsonPost(c, invalidJson)

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	id := uint(1)

	mockArticlesSourceServices.On("Delete",SUPER_ADMIN_ROLE, id).Return(fmt.Errorf("service failed"))

	articleSourceHandler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestDeleteSource_PayloadInvalid(t *testing.T) {
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

	invalidJson := "invalid"

	mockJsonPost(c, invalidJson)

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestDeleteSource_NotFoundRole(t *testing.T) {
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

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestUpdateSource_NotFoundRole(t *testing.T) {
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

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestUpdateSource_PayloadInvalid(t *testing.T) {
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

	invalidJson := "invalid"

	mockJsonPost(c, invalidJson)

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	articleSourceHandler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestUpdateSource_ServicesError(t *testing.T) {
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

	mockJson := entities.ArticlesSource{
		Title: "Test",
		Description: "Test",
	}

	mockJsonPost(c, mockJson)

	mockArticlesSourceServices := new(mocks.ArticlesSourceServices)
	articleSourceHandler := NewArticlesSourceHandler(mockArticlesSourceServices)
	assert := assert.New(t)

	mockArticlesSourceServices.On("Update",SUPER_ADMIN_ROLE, mockJson).Return(fmt.Errorf("service failed"))

	articleSourceHandler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}