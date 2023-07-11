package services

import (
	"net/url"
	"server/entities"
)

type UserServices interface {
	GetUser(username string) (u *entities.User, err error)
	List(page int, pageSize int) ([]UserResponse, error)
	Delete(role string, id uint) error
	ChangeRole(role string, id uint, newRole string) error
	UserUpgrateRole(role string, username string) (string, string, error)
	Count() (int, error)

	CreateUser(registerUserInput *RegisterUserInput) (*entities.User, error)

	LoginWithUsername(inputUser *LoginUserInput) (accessToken string, refreshToken string, err error)
	LoginWithEmail(inputUser *LoginUserInput) (accessToken string, refreshToken string, err error)
	// UpdateUser(registerUserInput *RegisterUserInput) error
	GoogleOAuth(googleUser *GoogleUserResult) (accessToken string, refreshToken string, err error)

	AccessAdminPage(role string) error
}

type RoleServices interface {
	List(page int, pageSize int) ([]RoleResponse, error)
	Validate(roleName string) (err error)
	Get(roleName string) (RoleResponse, error)
	Delete(id uint) error
	Create(rolePayload RoleResponse) error
	Count() (int, error)
	Update(rolePayload RoleResponse) error
	GrantPermission(userRole string, entity string, method string) bool
	ListRoleName() ([]string, error) 
}

type PermissionServices interface {
	List() ([]PermissionResponse, error)
}

type ArticleServices interface {
	SearchArticlesAcrossUserFollowedSources(username string, keyword string, page int, pageSize int) ([]ArticleResponse, int64, error)

	GetArticlesPaginationByUserFollowedSource(username string, page int, pageSize int) ([]ArticleForReadResponse, error)
	GetArticlesPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleForReadResponse, error)

	GetUnreadArticlesPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleForReadResponse, error)
	GetUnreadArticlesByUserFollowedSource(username string, page int, pageSize int) ([]ArticleForReadResponse, error)

	GetReadLaterListPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleForReadResponse, error)
	GetReadLaterListPaginationByUserFollowedSource(username string, page int, pageSize int) ([]ArticleForReadResponse, error)

	GetRecentlyReadArticle(username string, page int, pageSize int) ([]ArticleForReadResponse, error)

	GetTredingArticle(username string) ([]TredingArticleResponse, error)

	ListAll(page int, pageSize int) ([]ArticleResponse, error)
	Count() (int, error)
	Delete(articleID uint) error
	AdminSearchArticlesWithFilter(keyword string, page int, pageSize int, articlesSourceID uint) ([]ArticleResponse, int64, error)
	CountArticleCreateAWeekAgoByArticlesSourceID(articlesSourceID uint) (int64, error)

	CreateIfNotExist(article *entities.Article) error
}

type ArticlesSourceServices interface {
	GetByTopicIDPaginate(topicID uint, page int, pageSize int) ([]ArticlesSourceResponseRender, int64, error)
	Search(keyword string, page int, pageSize int) ([]ArticlesSourceResponseRender, int64, error)
	GetMostActiveSources() ([]ArticlesSourceRecommended, error)
	GetWithID(id uint) (ArticlesSourceResponseRender, error)
	ListAll() ([]ArticlesSourceResponseRender, error)
	SearchWithFilter(keyword string, page int, pageSize int, topicID uint) ([]ArticlesSourceResponseRender, int64, error)

	CreateIfNotExist(articlesSource entities.ArticlesSource) (entities.ArticlesSource, error)
	UpdateTopicOneSource(articlesSource entities.ArticlesSource, newTopicId uint) error
	UpdateTopicAllSource(oldTopicId uint, newTopicId uint) error
	UserFollow(articlesSourceID uint) error
	UserUnfollow(articlesSourceID uint) error

	Count() (int, error)
	ListAllPaging(page int, pageSize int) ([]ArticlesSourceResponseRender, error)
	Delete(role string, sourceID uint) error
	Update(role string, articlesSource entities.ArticlesSource) error
}

type CrawlerServices interface {
	TestRSSCrawler(crawler entities.Crawler) (*ArticlesSourceResponseCrawl, []*ArticleResponse, error)
	TestCustomCrawler(crawler entities.Crawler) (*ArticlesSourceResponseCrawl, []*ArticleResponse, error)

	CreateCrawlerWithCorrespondingArticlesSource(role string, payload CreateCrawlerPayload) error
	GetHtmlPage(url *url.URL) error
	Get(id uint) (*entities.Crawler, error)

	CreateCrawlerCronjobFromDB() error
	UpdateSchedule(role string, id uint, newSchedule string) error
	Update(role string, crawler entities.Crawler) error
	ListAllPaging(page int, pageSize int) ([]CrawlerResponse, int64, error)

	CronjobOnDay(timeString string) (*[24]ChartDay, error)
	CronjobOnHour(timeString string) (*[60]ChartHour, error)
}

type CronjobServices interface {
	CreateCrawlerCronjob(crawler entities.Crawler)
	RemoveCronjob(crawler entities.Crawler) error

	GetCronjobRuntime() []CronjobResponse
	CronjobOnHour(timeString string) (*[60]ChartHour, error)
	CronjobOnDay(timeString string) (*[24]ChartDay, error)
}

type CategoryServices interface {
	ListName() ([]CategoryResponse, error)
	ListAll() ([]CategoryResponse, error)
	GetPagination(page int, pageSize int) ([]CategoryResponse, error)
	Count() (int, error)

	CreateIfNotExist(role string, category entities.Category) error
	Delete(role string, category entities.Category) error
	Update(role string, payload UpdateNameCategoryPayload) error
}

type TopicServices interface {
	List() ([]TopicResponse, error)
	GetPagination(page int, pageSize int) ([]TopicResponse, error)
	GetByCategory(categoryID uint) ([]TopicResponse, error)
	SearchByName(keyword string) ([]TopicResponse, error)
	SearchTopicAndArticlesSourcePaginate(keyword string, page int, pageSize int) ([]TopicResponse, []ArticlesSourceResponseRender, int64, error)
	Count() (int, error)

	CreateIfNotExist(role string, topic entities.Topic) error
	Delete(role string, topic entities.Topic) error
	Update(role string, topic entities.Topic) error
	UpdateWhenDeteleCategory(oldCategoryID uint, newCategoryID uint) error
}

type FollowServices interface {
	Follow(username string, articlesSourceID uint) error
	Unfollow(username string, articlesSourceID uint) error
	GetUserFollowedSources(username string) ([]ArticlesSourceUserFollow, error)
	GetNewestSourceUpdatedID(username string) ([]uint, error)
}

type ReadServices interface {
	MarkAllAsReadBySourceID(username string, articlesSourceID uint) error
	MarkAllAsReadByUserFollowedSource(username string) error
	MarkArticleAsRead(username string, articleID uint, articlesSourceID uint) error
	MarkArticleAsUnRead(username string, articleID uint, articlesSourceID uint) error
}

type ReadLaterServices interface {
	AddToReadLaterList(username string, articleID uint) error
	RemoveFromReadLaterList(username string, articleID uint) error
}
