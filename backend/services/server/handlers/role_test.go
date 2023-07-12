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

func TestRoleGet_NotfoundRole(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	handler.Get(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleGet_ServiceError(t *testing.T) {
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

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	roleResponse := services.RoleResponse{}

	mockService.On("Get", SUPER_ADMIN_ROLE).Return(roleResponse, fmt.Errorf("service failed"))

	handler.Get(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleList_InvalidPage(t *testing.T) {
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

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	handler.List(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleList_InvalidPageSize(t *testing.T) {
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

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	handler.List(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleList_ServiceError(t *testing.T) {
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

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	page := 1
	pageSize := 10

	roleResponse := []services.RoleResponse{}

	mockService.On("List", page, pageSize).Return(roleResponse, fmt.Errorf("error"))

	handler.List(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleTotal_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	
	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	total := 1

	mockService.On("Count").Return(total, fmt.Errorf("error"))

	handler.Total(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleListNames_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	
	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	names := []string{}

	mockService.On("ListRoleName").Return(names, fmt.Errorf("error"))

	handler.ListNames(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleDelete_PayloadInvalid(t *testing.T) {
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

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	handler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleDelete_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := deleteRolePayload{
		ID: 1,
	}

	mockJsonPost(c, payload)

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	mockService.On("Delete", payload.ID).Return(fmt.Errorf("service failed"))

	handler.Delete(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleCreate_PayloadInvalid(t *testing.T) {
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

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	handler.Create(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleCreate_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := services.RoleResponse{}

	mockJsonPost(c, payload)

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	mockService.On("Create", payload).Return(fmt.Errorf("service failed"))

	handler.Create(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleCreate_ValidateError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := services.RoleResponse{}

	mockJsonPost(c, payload)

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	mockService.On("Create", payload).Return(fmt.Errorf("validate failed"))

	handler.Create(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleUpdate_PayloadInvalid(t *testing.T) {
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

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleUpdate_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := services.RoleResponse{}

	mockJsonPost(c, payload)

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	mockService.On("Update", payload).Return(fmt.Errorf("service failed"))

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestRoleUpdate_ValidateError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := services.RoleResponse{}

	mockJsonPost(c, payload)

	mockService := new(mocks.RoleServices)
	handler := NewRoleHandler(mockService)
	assert := assert.New(t)

	mockService.On("Update", payload).Return(fmt.Errorf("validate failed"))

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}