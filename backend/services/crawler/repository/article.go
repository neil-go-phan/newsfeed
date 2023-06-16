package repository

import (
	"crawler/entities"
	"fmt"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateIfNotExist(article entities.Article) (error) 
}	

type ArticleRepo struct {
	DB *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *ArticleRepo {
	return &ArticleRepo{
		DB: db,
	}
}

func (repo *ArticleRepo) CreateIfNotExist(article entities.Article) (error) {
	result := repo.DB.FirstOrCreate(&article, entities.Article{Title: article.Title, Link: article.Link})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("article already exist")
	}
	return nil
}