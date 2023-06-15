package repository

import (
	"fmt"
	"server/entities"

	"gorm.io/gorm"
)

type ArticlesSourcesRepository interface {
	CreateIfNotExist(articlesSource entities.ArticlesSource) (error)
}	

type ArticlesSourcesRepo struct {
	DB *gorm.DB
}

func NewArticlesSourcesRepo(db *gorm.DB) *ArticlesSourcesRepo {
	return &ArticlesSourcesRepo{
		DB: db,
	}
}

func (repo *ArticlesSourcesRepo) CreateIfNotExist(articlesSource entities.ArticlesSource) (error) {
	result := repo.DB.Where(entities.ArticlesSource{Link: articlesSource.Link}).FirstOrCreate(&articlesSource)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("article source already exist")
	}

	return nil
}

// func (repo *ArticlesSourcesRepo) Upsert(articlesSource entities.ArticlesSource) (error) {
// 	result := repo.DB.Model(&articlesSource).Where("link = ?", articlesSource.Link).
// 	Updates(entities.ArticlesSource{
// 		Title: articlesSource.Title,
// 		Description: articlesSource.Description,
// 		Link: articlesSource.Link,
// 		FeedLink: articlesSource.FeedLink,
// 		Image: articlesSource.Image,
// 	})

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	if result.RowsAffected == 0 {
// 		err := repo.DB.Create(&articlesSource).Error
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }