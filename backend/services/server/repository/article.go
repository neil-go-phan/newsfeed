// get one (get full) - get 6 article (get image, title, published)
package repository

import (
	"fmt"
	"server/entities"
	"server/helpers"
	"strings"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	SearchArticlesAcrossSources(keyword string, page int, pageSize int) ([]entities.Article, int64, error)
	GetPaginationByArticlesSourceID(articlesSourceID uint, page int, pageSize int) ([]entities.Article, error)
	GetPaginationByUserFollowedSource(username string, page int, pageSize int) ([]entities.Article, error)

	CreateIfNotExist(article *entities.Article) error
}

type ArticleRepo struct {
	DB *gorm.DB
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

	err := repo.DB.Where("articles_source_id = ?", articlesSourceID).Scopes(helpers.Paginate(page, pageSize)).Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) GetPaginationByUserFollowedSource(username string, page int, pageSize int) ([]entities.Article, error) {
	articles := make([]entities.Article, 0)

	err := repo.DB.
		Distinct("title", "description", "link", "published", "authors", "follows.articles_source_id", "created_at").
		Joins("LEFT JOIN follows ON articles.articles_source_id = follows.articles_source_id", repo.DB.Where(&entities.Follow{Username: username})).
		Scopes(helpers.Paginate(page, pageSize)).
		Order("created_at desc").
		Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (repo *ArticleRepo) SearchArticlesAcrossSources(keyword string, page int, pageSize int) ([]entities.Article, int64, error) {
	articles := make([]entities.Article, 0)
	searchKeyword := fmt.Sprint("%" + strings.ToLower(keyword) + "%")

	var found int64

	err := repo.DB.
		Scopes(helpers.Paginate(page, pageSize)).
		Where("LOWER(title) LIKE ? or LOWER(description) LIKE ?", searchKeyword, searchKeyword).
		Find(&articles).
		Count(&found).Error
	if err != nil {
		return articles, found, err
	}
	return articles, found, nil
}
