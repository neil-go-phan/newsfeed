// get one (get full) - get 6 article (get image, title, published)
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
	GetPaginationByArticlesSourceID(articlesSourceID uint, page int, pageSize int) ([]entities.Article, error)
	GetPaginationByUserFollowedSource(username string, page int, pageSize int) ([]ArticleLeftJoinRead, error) 
	GetPaginationByArticlesSourceIDWithReadStatus(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleLeftJoinRead, error)
	CountArticleCreateAWeekAgoByArticlesSourceID(articlesSourceID uint) (int64, error)

	CreateIfNotExist(article *entities.Article) error
}

type ArticleRepo struct {
	DB *gorm.DB
}

type ArticleLeftJoinRead struct {
	entities.Article
	Username  string `json:"username"`
	ArticleID uint   `json:"article_id"`
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

func (repo *ArticleRepo) GetPaginationByArticlesSourceID(articlesSourceID uint, page int, pageSize int) ([]entities.Article, error) {
	articles := make([]entities.Article, 0)

	err := repo.DB.
		Where("articles_source_id = ?", articlesSourceID).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("created_at desc").
		Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) GetPaginationByArticlesSourceIDWithReadStatus(username string, articlesSourceID uint, page int, pageSize int) ([]ArticleLeftJoinRead, error) {
	articles := make([]ArticleLeftJoinRead, 0)
	subQuery := repo.DB.
		Model(&entities.Read{}).
		Select("*").
		Where("username = ? AND articles_source_id = ?", username, articlesSourceID)

	err := repo.DB.
		Model(&entities.Article{}).
		Distinct("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "created_at", "username", "article_id").
		Joins("LEFT JOIN (?) q on articles.id = q.article_id", subQuery).
		Where("articles.articles_source_id = ?", articlesSourceID).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("created_at desc").
		Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) GetPaginationByUserFollowedSource(username string, page int, pageSize int) ([]ArticleLeftJoinRead, error) {
	articles := make([]ArticleLeftJoinRead, 0)
	// subQuery := repo.DB.
	// 	Model(&entities.Follow{}).
	// 	Select("follows.articles_source_id", "reads.article_id", "reads.username").
	// 	Joins("JOIN reads on (follows.username = ? AND follows.articles_source_id = reads.articles_source_id AND follows.username = reads.username)", username)
		subQuery := repo.DB.
		Model(&entities.Follow{}).
		Select("articles_source_id").
		Where("username = ?", username)

	err := repo.DB.
		Model(&entities.Article{}).
		Distinct("id", "title", "description", "link", "published", "authors", "articles.articles_source_id", "created_at", "username", "article_id").
		Joins("JOIN (?) q on q.articles_source_id = articles.articles_source_id", subQuery).
		Joins("LEFT JOIN reads r on articles.id = r.article_id").
		Scopes(helpers.Paginate(page, pageSize)).
		Order("created_at desc").
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
