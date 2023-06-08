// create - get one (get full) - get 6 article (get image, title, published)
package repository

import (
	"crawler/entities"

	"gorm.io/gorm"
)



type ArticleRepository interface {
	CreateIfNotExist(article *entities.Article) (error) 
}	

type ArticleRepo struct {
	DB *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *ArticleRepo {
	return &ArticleRepo{
		DB: db,
	}
}

func (repo *ArticleRepo) CreateIfNotExist(article *entities.Article) (error) {
	err := repo.DB.FirstOrCreate(article, entities.Article{Title: article.Title, Link: article.Link}).Error
	if err != nil {
		return err
	}
	return nil
}