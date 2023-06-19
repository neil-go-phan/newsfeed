package repository

import (
	"fmt"
	"server/entities"

	"gorm.io/gorm"
)

type ArticlesSourcesRepository interface {
	CreateIfNotExist(articlesSource entities.ArticlesSource) (entities.ArticlesSource, error)
	GetWithTopic(topicID uint) ([]entities.ArticlesSource, error)
	UpdateTopicOneSource(articlesSource entities.ArticlesSource, newTopicId uint) error
	UpdateTopicAllSource(oldTopicId uint, newTopicId uint) error
}

type ArticlesSourcesRepo struct {
	DB *gorm.DB
}

func NewArticlesSourcesRepo(db *gorm.DB) *ArticlesSourcesRepo {
	return &ArticlesSourcesRepo{
		DB: db,
	}
}

func (repo *ArticlesSourcesRepo) CreateIfNotExist(articlesSource entities.ArticlesSource) (entities.ArticlesSource, error) {
	result := repo.DB.Where(entities.ArticlesSource{Link: articlesSource.Link}).FirstOrCreate(&articlesSource)
	if result.Error != nil {
		return articlesSource, result.Error
	}
	if result.RowsAffected == 0 {
		return articlesSource, fmt.Errorf("article source already exist")
	}

	return articlesSource, nil
}

func (repo *ArticlesSourcesRepo) GetWithTopic(topicID uint) ([]entities.ArticlesSource, error) {
	articlesSources := make([]entities.ArticlesSource, 0)
	err := repo.DB.Where("topic_id = ?", topicID).Find(&articlesSources).Error
	if err != nil {
		return articlesSources, err
	}
	return articlesSources, nil
}

func (repo *ArticlesSourcesRepo) UpdateTopicOneSource(articlesSource entities.ArticlesSource, newTopicId uint) error {
	err := repo.DB.Model(&articlesSource).Update("topic_id", newTopicId).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ArticlesSourcesRepo) UpdateTopicAllSource(oldTopicId uint, newTopicId uint) error {
	err := repo.DB.Model(&entities.ArticlesSource{}).
		Where("topic_id = ?", oldTopicId).
		Update("topic_id", newTopicId).Error
	if err != nil {
		return err
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
