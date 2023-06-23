package repository

import (
	"fmt"
	"server/entities"
	"server/helpers"
	"strings"

	"gorm.io/gorm"
)

type ArticlesSourcesRepository interface {
	GetWithTopicPaginate(topicID uint, page int, pageSize int) ([]entities.ArticlesSource, int64, error)
	SearchByTitleAndDescriptionPaginate(keyword string, page int, pageSize int) ([]entities.ArticlesSource, int64, error)

	CreateIfNotExist(articlesSource entities.ArticlesSource) (entities.ArticlesSource, error)
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

func (repo *ArticlesSourcesRepo) GetWithTopicPaginate(topicID uint, page int, pageSize int) ([]entities.ArticlesSource, int64, error) {
	articlesSources := make([]entities.ArticlesSource, 0)
	var found int64

	err := repo.DB.Scopes(helpers.Paginate(page, pageSize)).Where("topic_id = ?", topicID).Find(&articlesSources).Count(&found).Error
	if err != nil {
		return articlesSources, found, err
	}
	return articlesSources, found, nil
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

func (repo *ArticlesSourcesRepo) SearchByTitleAndDescriptionPaginate(keyword string, page int, pageSize int) ([]entities.ArticlesSource, int64, error) {
	articlesSources := make([]entities.ArticlesSource, 0)
	searchKeyword := fmt.Sprint(strings.ToLower(keyword) + "%")
	var found int64

	err := repo.DB.Scopes(helpers.Paginate(page, pageSize)).Where("LOWER(title) LIKE ? or LOWER(description) LIKE ?", searchKeyword, searchKeyword).Find(&articlesSources).Count(&found).Error
	if err != nil {
		return articlesSources, found, err
	}
	return articlesSources, found, nil
}
