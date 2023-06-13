package services

import (
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
	Create(articlesSource *entities.ArticlesSource) error
}

type CrawlerServices interface {
	TestCrawler(crawler *entities.Crawler) (*ArticlesSourceResponse, []*ArticleResponse, error)
	FirstCrawl(crawler *entities.Crawler) (error)
	ScheduledCrawl(crawlerID uint) (error)
}
