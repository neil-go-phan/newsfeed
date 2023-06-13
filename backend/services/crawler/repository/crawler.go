package repository

import (
	"crawler/entities"

	"gorm.io/gorm"
)

type CrawlerRepository interface {
	Create(crawler *entities.Crawler) (*entities.Crawler, error) 
	Get(id uint) (*entities.Crawler, error)
}	

type CrawlerRepo struct {
	DB *gorm.DB
}

func NewCrawlerRepo(db *gorm.DB) *CrawlerRepo {
	return &CrawlerRepo{
		DB: db,
	}
}

func (repo *CrawlerRepo) Create(crawler *entities.Crawler) (*entities.Crawler, error) {
	err := repo.DB.Create(crawler).Error
	if err != nil {
		return nil, err
	}
	return crawler, nil
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