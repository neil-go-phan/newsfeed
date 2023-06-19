package repository

import (
	"fmt"
	"server/entities"

	"gorm.io/gorm"
)

type CrawlerRepository interface {
	Get(id uint) (*entities.Crawler, error)
	CreateIfNotExist(crawler entities.Crawler) (entities.Crawler, error)
	List() ([]entities.Crawler, error)
}

type CrawlerRepo struct {
	DB *gorm.DB
}

func NewCrawlerRepo(db *gorm.DB) *CrawlerRepo {
	return &CrawlerRepo{
		DB: db,
	}
}

func (repo *CrawlerRepo) Get(id uint) (*entities.Crawler, error) {
	crawler := new(entities.Crawler)
	crawler.ID = id
	err := repo.DB.First(crawler).Error
	if err != nil {
		return nil, err
	}
	return crawler, nil
}

func (repo *CrawlerRepo) CreateIfNotExist(crawler entities.Crawler) (entities.Crawler, error) {
	result := repo.DB.Where(entities.Crawler{SourceLink: crawler.SourceLink, ArticlesSourceID: crawler.ArticlesSourceID}).FirstOrCreate(&crawler)
	if result.Error != nil {
		return crawler, result.Error
	}
	if result.RowsAffected == 0 {
		return crawler, fmt.Errorf("crawler already exist")
	}

	return crawler, nil
}

func (repo *CrawlerRepo) List() ([]entities.Crawler, error) {
	crawlers := make([]entities.Crawler, 0)
	err := repo.DB.Find(&crawlers).Error
	if err != nil {
		return nil, err
	}
	return crawlers, nil
}
