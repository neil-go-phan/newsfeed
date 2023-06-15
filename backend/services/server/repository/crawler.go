package repository

import (
	"fmt"
	"server/entities"

	"gorm.io/gorm"
)

type CrawlerRepository interface {
	Get(id uint) (*entities.Crawler, error)
	CreateIfNotExist(crawler entities.Crawler) (error) 
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

func (repo *CrawlerRepo) CreateIfNotExist(crawler entities.Crawler) (error) {
	crawler.Schedule = "@daily"
	result := repo.DB.Where(entities.Crawler{SourceLink: crawler.SourceLink}).FirstOrCreate(&crawler)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("crawler already exist")
	}

	return nil
}