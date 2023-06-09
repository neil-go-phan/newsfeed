// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entities "server/entities"
	services "server/services"

	mock "github.com/stretchr/testify/mock"
)

// ArticleServices is an autogenerated mock type for the ArticleServices type
type ArticleServices struct {
	mock.Mock
}

// AdminSearchArticlesWithFilter provides a mock function with given fields: keyword, page, pageSize, articlesSourceID
func (_m *ArticleServices) AdminSearchArticlesWithFilter(keyword string, page int, pageSize int, articlesSourceID uint) ([]services.ArticleResponse, int64, error) {
	ret := _m.Called(keyword, page, pageSize, articlesSourceID)

	var r0 []services.ArticleResponse
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(string, int, int, uint) ([]services.ArticleResponse, int64, error)); ok {
		return rf(keyword, page, pageSize, articlesSourceID)
	}
	if rf, ok := ret.Get(0).(func(string, int, int, uint) []services.ArticleResponse); ok {
		r0 = rf(keyword, page, pageSize, articlesSourceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.ArticleResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int, uint) int64); ok {
		r1 = rf(keyword, page, pageSize, articlesSourceID)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(string, int, int, uint) error); ok {
		r2 = rf(keyword, page, pageSize, articlesSourceID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Count provides a mock function with given fields:
func (_m *ArticleServices) Count() (int, error) {
	ret := _m.Called()

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountArticleCreateAWeekAgoByArticlesSourceID provides a mock function with given fields: articlesSourceID
func (_m *ArticleServices) CountArticleCreateAWeekAgoByArticlesSourceID(articlesSourceID uint) (int64, error) {
	ret := _m.Called(articlesSourceID)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int64, error)); ok {
		return rf(articlesSourceID)
	}
	if rf, ok := ret.Get(0).(func(uint) int64); ok {
		r0 = rf(articlesSourceID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(articlesSourceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateIfNotExist provides a mock function with given fields: article
func (_m *ArticleServices) CreateIfNotExist(article *entities.Article) error {
	ret := _m.Called(article)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Article) error); ok {
		r0 = rf(article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: articleID
func (_m *ArticleServices) Delete(articleID uint) error {
	ret := _m.Called(articleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(articleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetArticlesPaginationByArticlesSourceID provides a mock function with given fields: username, articlesSourceID, page, pageSize
func (_m *ArticleServices) GetArticlesPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	ret := _m.Called(username, articlesSourceID, page, pageSize)

	var r0 []services.ArticleForReadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, uint, int, int) ([]services.ArticleForReadResponse, error)); ok {
		return rf(username, articlesSourceID, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, uint, int, int) []services.ArticleForReadResponse); ok {
		r0 = rf(username, articlesSourceID, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.ArticleForReadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, uint, int, int) error); ok {
		r1 = rf(username, articlesSourceID, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticlesPaginationByUserFollowedSource provides a mock function with given fields: username, page, pageSize
func (_m *ArticleServices) GetArticlesPaginationByUserFollowedSource(username string, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	ret := _m.Called(username, page, pageSize)

	var r0 []services.ArticleForReadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]services.ArticleForReadResponse, error)); ok {
		return rf(username, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []services.ArticleForReadResponse); ok {
		r0 = rf(username, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.ArticleForReadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(username, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMostReadInDay provides a mock function with given fields:
func (_m *ArticleServices) GetMostReadInDay() (entities.Article, error) {
	ret := _m.Called()

	var r0 entities.Article
	var r1 error
	if rf, ok := ret.Get(0).(func() (entities.Article, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() entities.Article); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entities.Article)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReadLaterListPaginationByArticlesSourceID provides a mock function with given fields: username, articlesSourceID, page, pageSize
func (_m *ArticleServices) GetReadLaterListPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	ret := _m.Called(username, articlesSourceID, page, pageSize)

	var r0 []services.ArticleForReadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, uint, int, int) ([]services.ArticleForReadResponse, error)); ok {
		return rf(username, articlesSourceID, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, uint, int, int) []services.ArticleForReadResponse); ok {
		r0 = rf(username, articlesSourceID, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.ArticleForReadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, uint, int, int) error); ok {
		r1 = rf(username, articlesSourceID, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReadLaterListPaginationByUserFollowedSource provides a mock function with given fields: username, page, pageSize
func (_m *ArticleServices) GetReadLaterListPaginationByUserFollowedSource(username string, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	ret := _m.Called(username, page, pageSize)

	var r0 []services.ArticleForReadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]services.ArticleForReadResponse, error)); ok {
		return rf(username, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []services.ArticleForReadResponse); ok {
		r0 = rf(username, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.ArticleForReadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(username, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecentlyReadArticle provides a mock function with given fields: username, page, pageSize
func (_m *ArticleServices) GetRecentlyReadArticle(username string, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	ret := _m.Called(username, page, pageSize)

	var r0 []services.ArticleForReadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]services.ArticleForReadResponse, error)); ok {
		return rf(username, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []services.ArticleForReadResponse); ok {
		r0 = rf(username, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.ArticleForReadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(username, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTredingArticle provides a mock function with given fields: username
func (_m *ArticleServices) GetTredingArticle(username string) ([]services.TredingArticleResponse, error) {
	ret := _m.Called(username)

	var r0 []services.TredingArticleResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]services.TredingArticleResponse, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) []services.TredingArticleResponse); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.TredingArticleResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnreadArticlesByUserFollowedSource provides a mock function with given fields: username, page, pageSize
func (_m *ArticleServices) GetUnreadArticlesByUserFollowedSource(username string, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	ret := _m.Called(username, page, pageSize)

	var r0 []services.ArticleForReadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]services.ArticleForReadResponse, error)); ok {
		return rf(username, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []services.ArticleForReadResponse); ok {
		r0 = rf(username, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.ArticleForReadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(username, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnreadArticlesPaginationByArticlesSourceID provides a mock function with given fields: username, articlesSourceID, page, pageSize
func (_m *ArticleServices) GetUnreadArticlesPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	ret := _m.Called(username, articlesSourceID, page, pageSize)

	var r0 []services.ArticleForReadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, uint, int, int) ([]services.ArticleForReadResponse, error)); ok {
		return rf(username, articlesSourceID, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, uint, int, int) []services.ArticleForReadResponse); ok {
		r0 = rf(username, articlesSourceID, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.ArticleForReadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, uint, int, int) error); ok {
		r1 = rf(username, articlesSourceID, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAll provides a mock function with given fields: page, pageSize
func (_m *ArticleServices) ListAll(page int, pageSize int) ([]services.ArticleResponse, error) {
	ret := _m.Called(page, pageSize)

	var r0 []services.ArticleResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]services.ArticleResponse, error)); ok {
		return rf(page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(int, int) []services.ArticleResponse); ok {
		r0 = rf(page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.ArticleResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchArticlesAcrossUserFollowedSources provides a mock function with given fields: username, keyword, page, pageSize
func (_m *ArticleServices) SearchArticlesAcrossUserFollowedSources(username string, keyword string, page int, pageSize int) ([]services.ArticleResponse, int64, error) {
	ret := _m.Called(username, keyword, page, pageSize)

	var r0 []services.ArticleResponse
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]services.ArticleResponse, int64, error)); ok {
		return rf(username, keyword, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []services.ArticleResponse); ok {
		r0 = rf(username, keyword, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.ArticleResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) int64); ok {
		r1 = rf(username, keyword, page, pageSize)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(string, string, int, int) error); ok {
		r2 = rf(username, keyword, page, pageSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewArticleServices interface {
	mock.TestingT
	Cleanup(func())
}

// NewArticleServices creates a new instance of ArticleServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewArticleServices(t mockConstructorTestingTNewArticleServices) *ArticleServices {
	mock := &ArticleServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
