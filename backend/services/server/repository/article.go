// get one (get full) - get 6 article (get image, title, published)
package repository

import (
	"server/entities"
	"server/helpers"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetPaginationByArticlesSourceID(articlesSourceID uint, page int, pageSize int) ([]entities.Article, error)
	
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
