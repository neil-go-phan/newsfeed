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

func TestCrawlerTestRSSCrawler_PayloadInvalid(t *testing.T) {
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

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.TestRSSCrawler(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerTestRSSCrawler_ServicesError(t *testing.T) {
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

	crawler := entities.Crawler{
		SourceLink: "test.com",
		FeedLink:   "test.com",
		CrawlType:  "feed",
	}

	mockJsonPost(c, crawler)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	articlesSource := &services.ArticlesSourceResponseCrawl{}
	articles := []*services.ArticleResponse{}

	service.On("TestRSSCrawler", crawler).Return(articlesSource, articles, fmt.Errorf("service failed"))

	handler.TestRSSCrawler(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerTestCustomCrawler_PayloadInvalid(t *testing.T) {
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

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.TestCustomCrawler(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerTestCustomCrawler_ServicesError(t *testing.T) {
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

	crawler := entities.Crawler{
		SourceLink: "test.com",
		FeedLink:   "test.com",
		CrawlType:  "feed",
	}

	mockJsonPost(c, crawler)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	articlesSource := &services.ArticlesSourceResponseCrawl{}
	articles := []*services.ArticleResponse{}

	service.On("TestCustomCrawler", crawler).Return(articlesSource, articles, fmt.Errorf("service failed"))

	handler.TestCustomCrawler(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerCreateCrawler_PayloadInvalid(t *testing.T) {
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

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.CreateCrawler(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerCreateCrawler_NotFoundRole(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := services.CreateCrawlerPayload{
		ArticlesSource: services.ArticlesSourceFromFrontend{},
		Crawler: services.CrawlerFromFrontend{},
	}

	mockJsonPost(c, payload)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.CreateCrawler(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'bad request' message to be returned")
}

func TestCrawlerCreateCrawler_ServicesError(t *testing.T) {
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

	payload := services.CreateCrawlerPayload{
		ArticlesSource: services.ArticlesSourceFromFrontend{},
		Crawler: services.CrawlerFromFrontend{},
	}

	mockJsonPost(c, payload)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	service.On("CreateCrawlerWithCorrespondingArticlesSource",SUPER_ADMIN_ROLE, payload).Return(fmt.Errorf("service failed"))

	handler.CreateCrawler(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerCreateCrawler_ValidateError(t *testing.T) {
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

	payload := services.CreateCrawlerPayload{
		ArticlesSource: services.ArticlesSourceFromFrontend{},
		Crawler: services.CrawlerFromFrontend{},
	}

	mockJsonPost(c, payload)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	service.On("CreateCrawlerWithCorrespondingArticlesSource",SUPER_ADMIN_ROLE, payload).Return(fmt.Errorf("validate failed"))

	handler.CreateCrawler(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerCreateCrawlerCronjobFromDB(t *testing.T) {
	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	service.On("CreateCrawlerCronjobFromDB").Return(fmt.Errorf("validate failed"))

	err := handler.CreateCrawlerCronjobFromDB()

	assert.Equal(fmt.Errorf("validate failed"), err, "Expected Err")
}

func TestCrawlerListAllPaging_InvalidPage(t *testing.T) {
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

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.ListAllPaging(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestCrawlerListAllPaging_InvalidPageSize(t *testing.T) {
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

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.ListAllPaging(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestCrawlerListAllPaging_ServicesError(t *testing.T) {
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

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	crawlers := []services.CrawlerResponse{}
	found := int64(1)

	page := 1
	pageSize := 10

	service.On("ListAllPaging", page, pageSize).Return(crawlers, found, fmt.Errorf("service error"))

	handler.ListAllPaging(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestCrawlerUpdateSchedule_PayloadInvalid(t *testing.T) {
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

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.UpdateSchedule(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerUpdateSchedule_NotFoundRole(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := updateSchedulePayload{
		ID: 1,
		Schedule: "@every 0h5m",
	}
	mockJsonPost(c, payload)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.UpdateSchedule(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'bad request' message to be returned")
}

func TestCrawlerUpdateSchedule_ServicesError(t *testing.T) {
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

	payload := updateSchedulePayload{
		ID: 1,
		Schedule: "@every 0h5m",
	}
	mockJsonPost(c, payload)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	service.On("UpdateSchedule",SUPER_ADMIN_ROLE, payload.ID, payload.Schedule).Return(fmt.Errorf("service failed"))

	handler.UpdateSchedule(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerUpdateSchedule_ValidateError(t *testing.T) {
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

	payload := updateSchedulePayload{
		ID: 1,
		Schedule: "@every 0h5m",
	}
	mockJsonPost(c, payload)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	service.On("UpdateSchedule",SUPER_ADMIN_ROLE, payload.ID, payload.Schedule).Return(fmt.Errorf("validate failed"))

	handler.UpdateSchedule(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerGet_InvalidID(t *testing.T) {
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

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.Get(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestCrawlerGet_ServicesError(t *testing.T) {
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

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	id := uint(1)

	crawler := &entities.Crawler{}

	service.On("Get",id).Return(crawler, fmt.Errorf("service error"))

	handler.Get(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(INTERNAL_SERVER_ERROR_RESPONSE, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestCrawlerUpdate_PayloadInvalid(t *testing.T) {
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

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerUpdate_NotFoundRole(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	payload := updatePayload{}
	mockJsonPost(c, payload)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(BAD_REQUEST_RESPONSE, w.Body.String(), "Expected 'bad request' message to be returned")
}

func TestCrawlerUpdate_ServicesError(t *testing.T) {
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

	payload := updatePayload{}
	mockJsonPost(c, payload)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	crawler := entities.Crawler{}

	service.On("Update",SUPER_ADMIN_ROLE, crawler).Return(fmt.Errorf("service failed"))

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(SERVICE_FAIL_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerUpdate_ValidateError(t *testing.T) {
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

	payload := updatePayload{}
	mockJsonPost(c, payload)

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	crawler := entities.Crawler{}

	service.On("Update",SUPER_ADMIN_ROLE, crawler).Return(fmt.Errorf("validate failed"))

	handler.Update(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(INVALID_INPUT_RESPONSE, w.Body.String(), "Expected 'internal server error' message to be returned")
}

func TestCrawlerGetCronjobOnHour_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("time", "1/1/0001")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	cronjobs := &[60]services.ChartHour{}
	time := "1/1/0001"

	service.On("CronjobOnHour", time).Return(cronjobs, fmt.Errorf("validate failed"))

	handler.GetCronjobOnHour(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
}

func TestCrawlerGetCronjobOnDay_ServiceError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("time", "1/1/0001")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	cronjobs := &[24]services.ChartDay{}
	time := "1/1/0001"

	service.On("CronjobOnDay", time).Return(cronjobs, fmt.Errorf("validate failed"))

	handler.GetCronjobOnDay(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
}

func TestCrawlerGetCronjobOnDay_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("time", "1/1/0001")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	service := new(mocks.CrawlerServices)
	handler := NewCrawlerHandler(service)
	assert := assert.New(t)

	cronjobs := &[24]services.ChartDay{}
	time := "1/1/0001"

	service.On("CronjobOnDay", time).Return(cronjobs, nil)

	handler.GetCronjobOnDay(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
}