package services

import (
	"net/url"
	"server/entities"
)

type UserServices interface {
	GetUser(username string) (u *entities.User, err error)
	ListUsers() (user *[]entities.User, err error)
	CreateUser(registerUserInput *RegisterUserInput) (*entities.User, error)
	// VerifyUser(username string, registerUserInput RegisterUserInput) (bool, error)
	DeleteUser(username string) error
	LoginWithUsername(inputUser *LoginUserInput) (accessToken string, refreshToken string, err error)
	LoginWithEmail(inputUser *LoginUserInput) (accessToken string, refreshToken string, err error)
	// UpdateUser(registerUserInput *RegisterUserInput) error
	GoogleOAuth(googleUser *GoogleUserResult) (accessToken string, refreshToken string, err error)
}

type RoleServices interface {
	ListRole() (role *[]entities.Role, err error) 
	Validate(roleName string) (err error)
	GetRole(roleName string) (*entities.Role, error)
}

type ArticleServices interface {
	CreateIfNotExist(article *entities.Article) error
}

type ArticlesSourceServices interface {
	CreateIfNotExist(articlesSource entities.ArticlesSource) (entities.ArticlesSource, error)
}

type CrawlerServices interface {
	TestRSSCrawler(crawler entities.Crawler) (*ArticlesSourceResponse, []*ArticleResponse, error)
	TestCustomCrawler(crawler entities.Crawler) (*ArticlesSourceResponse, []*ArticleResponse, error)
	
	CreateCrawlerWithCorrespondingArticlesSource(payload CreateCrawlerPayload) error
	GetHtmlPage(url *url.URL) error

	CreateCrawlerCronjobFromDB() error
}

type CronjobServices interface {
	CreateCrawlerCronjob(crawler entities.Crawler)
	GetCronjobRuntime() []CronjobResponse
	CronjobOnHour(timeString string) (*[60]ChartHour, error)
	CronjobOnDay(timeString string) (*[24]ChartDay, error)
}
