package repository

import (
	"server/entities"

	"gorm.io/gorm"
)



type CrawlerRepository interface {
	Get(id uint) (*entities.Crawler, error)
	// List() (*[]entities.Article, error) 
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