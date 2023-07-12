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

func TestCreateCategory_NotFoundRole(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	category := entities.Category{
		Name:         "Test",
		Illustration: "Test too",
	}

	mockJsonPost(c, category)

	services := new(mocks.CategoryServices)
	handler := NewCategoryHandler(services)
	assert := assert.New(t)

	handler.Create(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCreateCategory_PayloadInvalid(t *testing.T) {
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

	services := new(mocks.CategoryServices)
	handler := NewCategoryHandler(services)
	assert := assert.New(t)

	handler.Create(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCreateCategory_ServicesError(t *testing.T) {
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

	category := entities.Category{
		Name:         "Test",
		Illustration: "Test too",
	}

	mockJsonPost(c, category)

	services := new(mocks.CategoryServices)
	handler := NewCategoryHandler(services)
	assert := assert.New(t)

	services.On("CreateIfNotExist", SUPER_ADMIN_ROLE, category).Return(fmt.Errorf("service failed"))

	handler.Create(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestListNameCategory_ServicesError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	service := new(mocks.CategoryServices)
	handler := NewCategoryHandler(service)
	assert := assert.New(t)

	categories := []services.CategoryResponse{}

	service.On("ListName").Return(categories, fmt.Errorf("service failed"))

	handler.ListName(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestListAllCategory_ServicesError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	service := new(mocks.CategoryServices)
	handler := NewCategoryHandler(service)
	assert := assert.New(t)

	categories := []services.CategoryResponse{}

	service.On("ListAll").Return(categories, fmt.Errorf("service failed"))

	handler.ListAll(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCountCategory_ServicesError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	service := new(mocks.CategoryServices)
	handler := NewCategoryHandler(service)
	assert := assert.New(t)

	total := 1

	service.On("Count").Return(total, fmt.Errorf("service failed"))

	handler.Count(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestUpdateCategory_NotFoundRole(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	category := entities.Category{
		Name:         "Test",
		Illustration: "Test too",
	}

	mockJsonPost(c, category)

	services := new(mocks.CategoryServices)
	handler := NewCategoryHandler(services)
	assert := assert.New(t)

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestUpdateCategory_PayloadInvalid(t *testing.T) {
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

	category := "invalid"

	mockJsonPost(c, category)

	services := new(mocks.CategoryServices)
	handler := NewCategoryHandler(services)
	assert := assert.New(t)

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestUpdateCategory_ServicesError(t *testing.T) {
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

	payload := services.UpdateNameCategoryPayload{
		Category: entities.Category{
			Name:         "Test",
			Illustration: "Test too",
		},
		NewName:         "New name",
		NewIllustration: "New image",
	}

	mockJsonPost(c, payload)

	services := new(mocks.CategoryServices)
	handler := NewCategoryHandler(services)
	assert := assert.New(t)

	services.On("Update", SUPER_ADMIN_ROLE, payload).Return(fmt.Errorf("service failed"))

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestDeleteCategory_NotFoundRole(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	category := entities.Category{
		Name:         "Test",
		Illustration: "Test too",
	}

	mockJsonPost(c, category)

	services := new(mocks.CategoryServices)
	handler := NewCategoryHandler(services)
	assert := assert.New(t)

	handler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestDeleteCategory_PayloadInvalid(t *testing.T) {
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

	services := new(mocks.CategoryServices)
	handler := NewCategoryHandler(services)
	assert := assert.New(t)

	handler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestDeleteCategory_ServicesError(t *testing.T) {
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

	category := entities.Category{
		Name:         "Test",
		Illustration: "Test too",
	}

	mockJsonPost(c, category)

	services := new(mocks.CategoryServices)
	handler := NewCategoryHandler(services)
	assert := assert.New(t)

	services.On("Delete", SUPER_ADMIN_ROLE, category).Return(fmt.Errorf("service failed"))

	handler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetPaginationCategory_ServicesError(t *testing.T) {
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

	service := new(mocks.CategoryServices)
	handler := NewCategoryHandler(service)
	assert := assert.New(t)

	page := 1
	pageSize := 10

	categories := []services.CategoryResponse{}

	service.On("GetPagination", page, pageSize).Return(categories, fmt.Errorf("service failed"))

	handler.GetPagination(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetPaginationCategory_InvalidPage(t *testing.T) {
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

	service := new(mocks.CategoryServices)
	handler := NewCategoryHandler(service)
	assert := assert.New(t)

	handler.GetPagination(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestGetPaginationCategory_InvalidPageSize(t *testing.T) {
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

	service := new(mocks.CategoryServices)
	handler := NewCategoryHandler(service)
	assert := assert.New(t)

	handler.GetPagination(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}