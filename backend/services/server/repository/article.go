package repository

import (
	"fmt"
	"server/entities"
	"server/helpers"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	SearchArticlesAcrossUserFollowedSources(username string, keyword string, page int, pageSize int) ([]entities.Article, int64, error)
	GetArticlesPaginationByArticlesSourceID(articlesSourceID uint, page int, pageSize int) ([]entities.Article, int64, error)
	GetArticlesPaginationByUserFollowedSource(username string, page int, pageSize int) ([]ArticleLeftJoinRead, error)
	GetArticlesPaginationByArticlesSourceIDWithReadStatus(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleLeftJoinRead, error)
	GetUnreadArticlesPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleLeftJoinRead, error)
	GetReadLaterListPaginationByUserFollowedSource(username string, page int, pageSize int) ([]ArticleLeftJoinRead, error)
	GetReadLaterListPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleLeftJoinRead, error)
	GetUnreadArticlesByUserFollowedSource(username string, page int, pageSize int) ([]ArticleLeftJoinRead, error)
	GetRecentlyReadArticle(username string, page int, pageSize int) ([]ArticleLeftJoinRead, error)
	GetTredingArticle(username string) ([]TredingArticle, error)

	GetMostRead(from time.Time, to time.Time) (entities.Article, error) 
	
	ListAll(page int, pageSize int) ([]entities.Article, error)
	Count() (int, error)
	Delete(article entities.Article) error

	AdminSearchArticles(keyword string, page int, pageSize int) ([]entities.Article, int64, error)
	AdminSearchArticlesWithFilter(keyword string, page int, pageSize int, articlesSourceID uint) ([]entities.Article, int64, error)

	CountArticleCreateAWeekAgoByArticlesSourceID(articlesSourceID uint) (int64, error)
	CreateIfNotExist(article *entities.Article) error
}

type ArticleRepo struct {
	DB *gorm.DB
}

type ArticleLeftJoinRead struct {
	entities.Article
	IsRead      bool `json:"is_read"`
	IsReadLater bool `json:"is_read_later"`
}

type TredingArticle struct {
	ArticleLeftJoinRead
	ArticlesSource entities.ArticlesSource `json:"articles_source" gorm:"foreignKey:ArticlesSourceID"`
}

func NewArticleRepo(db *gorm.DB) *ArticleRepo {
	return &ArticleRepo{
		DB: db,
	}
}

func (repo *ArticleRepo) CreateIfNotExist(article *entities.Article) error {
	err := repo.DB.FirstOrCreate(article, entities.Article{Title: article.Title, Link: article.Link}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ArticleRepo) GetArticlesPaginationByArticlesSourceID(articlesSourceID uint, page int, pageSize int) ([]entities.Article, int64, error) {
	articles := make([]entities.Article, 0)
	var found int64

	err := repo.DB.
		Where("articles_source_id = ?", articlesSourceID).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("created_at desc").
		Find(&articles).
		Count(&found).Error
	if err != nil {
		return articles, found, err
	}
	return articles, found, nil
}

func (repo *ArticleRepo) GetArticlesPaginationByArticlesSourceIDWithReadStatus(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleLeftJoinRead, error) {
	articles := make([]ArticleLeftJoinRead, 0)
	subQuery1 := repo.DB.
		Model(&entities.Read{}).
		Select("*").
		Where("username = ? AND articles_source_id = ?", username, articlesSourceID)

	subQuery2 := repo.DB.
		Model(&entities.ReadLater{}).
		Select("*").
		Where("username = ?", username)

	err := repo.DB.
		Model(&entities.Article{}).
		Distinct("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "articles.created_at", "CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read", "CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later").
		Joins("LEFT JOIN (?) r on articles.id = r.article_id", subQuery1).
		Joins("LEFT JOIN (?) rl on articles.id = rl.article_id ", subQuery2).
		Where("articles.articles_source_id = ?", articlesSourceID).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("articles.created_at desc").
		Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) GetArticlesPaginationByUserFollowedSource(username string, page int, pageSize int) ([]ArticleLeftJoinRead, error) {
	articles := make([]ArticleLeftJoinRead, 0)
	subQuery1 := repo.DB.
		Model(&entities.Follow{}).
		Select("articles_source_id").
		Where("username = ?", username)

	subQuery2 := repo.DB.
		Model(&entities.ReadLater{}).
		Select("username, article_id").
		Where("username = ?", username)

	err := repo.DB.
		Model(&entities.Article{}).
		Distinct("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "articles.created_at", "CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read", "CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later").
		Joins("JOIN (?) f on f.articles_source_id = articles.articles_source_id", subQuery1).
		Joins("LEFT JOIN reads r on articles.id = r.article_id").
		Joins("LEFT JOIN (?) rl on articles.id = rl.article_id", subQuery2).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("articles.created_at desc").
		Find(&articles).Error

	if err != nil {
		return articles, err
	}

	return articles, nil
}

func (repo *ArticleRepo) GetReadLaterListPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleLeftJoinRead, error) {
	articles := make([]ArticleLeftJoinRead, 0)
	subQuery1 := repo.DB.
		Model(&entities.Read{}).
		Select("username, article_id").
		Where("username = ? AND articles_source_id = ?", username, articlesSourceID)

	subQuery2 := repo.DB.
		Model(&entities.ReadLater{}).
		Select("username, article_id").
		Where("username = ?", username)

	err := repo.DB.
		Model(&entities.Article{}).
		Distinct("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "articles.created_at", "CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read", "CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later").
		Joins("LEFT JOIN (?) r on articles.id = r.article_id", subQuery1).
		Joins("JOIN (?) rl on articles.id = rl.article_id ", subQuery2).
		Where("articles.articles_source_id = ?", articlesSourceID).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("articles.created_at desc").
		Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) GetReadLaterListPaginationByUserFollowedSource(username string, page int, pageSize int) ([]ArticleLeftJoinRead, error) {
	articles := make([]ArticleLeftJoinRead, 0)

	subQuery2 := repo.DB.
		Model(&entities.ReadLater{}).
		Select("username, article_id").
		Where("username = ?", username)

	err := repo.DB.
		Model(&entities.Article{}).
		Distinct("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "articles.created_at", "CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read", "CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later").
		Joins("LEFT JOIN reads r on articles.id = r.article_id").
		Joins("JOIN (?) rl on articles.id = rl.article_id", subQuery2).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("articles.created_at desc").
		Find(&articles).Error

	if err != nil {
		return articles, err
	}

	return articles, nil
}

func (repo *ArticleRepo) GetUnreadArticlesPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleLeftJoinRead, error) {
	articles := make([]ArticleLeftJoinRead, 0)
	subQuery := repo.DB.
		Model(&entities.Read{}).
		Select("*").
		Where("username = ? AND articles_source_id = ?", username, articlesSourceID)

	err := repo.DB.
		Model(&entities.Article{}).
		Distinct("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "articles.created_at").
		Joins("LEFT OUTER JOIN (?) q on articles.id = q.article_id", subQuery).
		Where("articles.articles_source_id = ? AND q.article_id IS NULL", articlesSourceID).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("articles.created_at desc").
		Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) GetUnreadArticlesByUserFollowedSource(username string, page int, pageSize int) ([]ArticleLeftJoinRead, error) {
	articles := make([]ArticleLeftJoinRead, 0)
	subQuery := repo.DB.
		Model(&entities.Follow{}).
		Select("articles_source_id").
		Where("username = ?", username)

	err := repo.DB.
		Model(&entities.Article{}).
		Distinct("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "articles.created_at").
		Joins("JOIN (?) q on q.articles_source_id = articles.articles_source_id", subQuery).
		Joins("LEFT OUTER JOIN reads r on articles.id = r.article_id").
		Where("r.article_id IS NULL").
		Scopes(helpers.Paginate(page, pageSize)).
		Order("articles.created_at desc").
		Find(&articles).Error

	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) GetRecentlyReadArticle(username string, page int, pageSize int) ([]ArticleLeftJoinRead, error) {
	articles := make([]ArticleLeftJoinRead, 0)
	subQuery := repo.DB.
		Model(&entities.Follow{}).
		Select("articles_source_id").
		Where("username = ?", username)

	err := repo.DB.
		Model(&entities.Article{}).
		Distinct("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "r.created_at", "CASE WHEN r.article_id IS NULL THEN false ELSE true END AS is_read").
		Joins("JOIN (?) q on q.articles_source_id = articles.articles_source_id", subQuery).
		Joins("JOIN reads r on articles.id = r.article_id").
		Scopes(helpers.Paginate(page, pageSize)).
		Order("r.created_at desc").
		Find(&articles).Error

	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) CountArticleCreateAWeekAgoByArticlesSourceID(articlesSourceID uint) (int64, error) {
	var count int64
	today := time.Now()
	aWeekAgo := today.AddDate(0, 0, -7)
	todayString := today.Format("2006-01-02")
	aWeekAgoString := aWeekAgo.Format("2006-01-02")
	err := repo.DB.
		Model(&entities.Article{}).
		Where("articles_source_id = ? AND created_at BETWEEN ? AND ?", articlesSourceID, aWeekAgoString+" 00:00:00", todayString+" 23:59:59").
		Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (repo *ArticleRepo) SearchArticlesAcrossUserFollowedSources(username string, keyword string, page int, pageSize int) ([]entities.Article, int64, error) {
	articles := make([]entities.Article, 0)
	searchKeyword := fmt.Sprint("%" + strings.ToLower(keyword) + "%")

	var found int64
	subQuery := repo.DB.Model(&entities.Follow{}).Select("articles_source_id").Where("username = ?", username)

	err := repo.DB.
		Distinct("title", "description", "link", "published", "authors", "q.articles_source_id", "created_at").
		Joins("JOIN (?) q on articles.articles_source_id = q.articles_source_id", subQuery).
		Where("LOWER(title) LIKE ? or LOWER(description) LIKE ?", searchKeyword, searchKeyword).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("created_at desc").
		Find(&articles).
		Count(&found).Error
	if err != nil {
		return articles, found, err
	}
	return articles, found, nil
}

func (repo *ArticleRepo) AdminSearchArticles(keyword string, page int, pageSize int) ([]entities.Article, int64, error) {
	articles := make([]entities.Article, 0)
	searchKeyword := fmt.Sprint("%" + strings.ToLower(keyword) + "%")

	var found int64

	err := repo.DB.
		Distinct("title", "description", "link", "published", "authors", "articles_source_id", "created_at").
		Where("LOWER(title) LIKE ? or LOWER(description) LIKE ?", searchKeyword, searchKeyword).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("created_at desc").
		Find(&articles).
		Count(&found).Error
	if err != nil {
		return articles, found, err
	}
	return articles, found, nil
}

func (repo *ArticleRepo) AdminSearchArticlesWithFilter(keyword string, page int, pageSize int, articlesSourceID uint) ([]entities.Article, int64, error) {
	articles := make([]entities.Article, 0)
	searchKeyword := fmt.Sprint("%" + strings.ToLower(keyword) + "%")

	var found int64

	err := repo.DB.
		Distinct("title", "description", "link", "published", "authors", "articles_source_id", "created_at").
		Where("articles_source_id = ? AND (LOWER(title) LIKE ? or LOWER(description) LIKE ?)", articlesSourceID, searchKeyword, searchKeyword).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("created_at desc").
		Find(&articles).
		Count(&found).Error
	if err != nil {
		return articles, found, err
	}
	return articles, found, nil
}

func (repo *ArticleRepo) GetTredingArticle(username string) ([]TredingArticle, error) {
	articles := make([]TredingArticle, 0)
	AMOUNT_OF_ARTICLES := 10
	today := time.Now()
	todayString := today.Format("2006-01-02")
	subQuery1 := repo.DB.
		Model(&entities.Read{}).
		Select("article_id", "count(article_id) as read").
		Where("created_at between ? AND ?", todayString+" 00:00:00", todayString+" 23:59:59").
		Group("article_id").
		Order("read desc").
		Limit(AMOUNT_OF_ARTICLES)

	subQuery2 := repo.DB.
		Model(&entities.ReadLater{}).
		Select("*").
		Where("username = ?", username)

	err := repo.DB.
		Model(&entities.Article{}).
		Select("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "articles.created_at", "CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later").
		Joins("JOIN (?) r on articles.id = r.article_id", subQuery1).
		Joins("LEFT JOIN (?) rl on articles.id = rl.article_id ", subQuery2).
		Preload("ArticlesSource").
		Order("r.read desc").
		Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) ListAll(page int, pageSize int) ([]entities.Article, error) {
	articles := make([]entities.Article, 0)
	err := repo.DB.
		Scopes(helpers.Paginate(page, pageSize)).
		Order("created_at desc").
		Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) Count() (int, error) {
	var count int64
	err := repo.DB.Table("articles").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (repo *ArticleRepo) Delete(article entities.Article) error {
	err := repo.DB.
		Where("id = ?", article.ID).
		Unscoped().
		Delete(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ArticleRepo) GetMostRead(from time.Time, to time.Time) (entities.Article, error) {
	fromString := from.Format("2006-01-02 15:04:05")
	toString := to.Format("2006-01-02 15:04:05")
	article := entities.Article{}
	subQuery1 := repo.DB.
		Model(&entities.Read{}).
		Select("article_id", "count(article_id) as read").
		Where("created_at between ? AND ?", fromString, toString).
		Group("article_id").
		Order("read desc").
		Limit(1)

		err := repo.DB.
		Select("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "articles.created_at").
		Joins("JOIN (?) r on articles.id = r.article_id", subQuery1).
		Order("r.read desc").
		First(&article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}
