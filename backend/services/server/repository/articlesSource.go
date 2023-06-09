package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type ArticlesSourcesRepository interface {
	Create(articlesSource *entities.ArticlesSource) (error)
}	

type ArticlesSourcesRepo struct {
	DB *gorm.DB
}

func NewArticlesSourcesRepo(db *gorm.DB) *ArticlesSourcesRepo {
	return &ArticlesSourcesRepo{
		DB: db,
	}
}

func (repo *ArticlesSourcesRepo) Create(articlesSource *entities.ArticlesSource) (error) {
	err := repo.DB.Create(articlesSource).Error
	if err != nil {
		return err
	}
	return nil
}