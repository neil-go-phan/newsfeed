package repository

import (
	"fmt"
	"server/entities"
	"server/helpers"

	"gorm.io/gorm"
)

//go:generate mockery --name CrawlerRepository
type CrawlerRepository interface {
	Get(id uint) (*entities.Crawler, error)
	List() ([]entities.Crawler, error) 
	ListAllPaging(page int, pageSize int) ([]entities.Crawler, int64, error)

	UpdateSchedule(id uint, newSchedule string) error 
	Update(crawler entities.Crawler) error
	CreateIfNotExist(crawler entities.Crawler) (entities.Crawler, error)
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

func (repo *CrawlerRepo) GetBySourceLink(sourceLink string) (*entities.Crawler, error) {
	crawler := new(entities.Crawler)
	err := repo.DB.Where("source_link = ?", sourceLink).First(crawler).Error
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


func (repo *CrawlerRepo) ListAllPaging(page int, pageSize int) ([]entities.Crawler, int64, error) {
	crawlers := make([]entities.Crawler, 0)
	var found int64

	err := repo.DB.
		Scopes(helpers.Paginate(page, pageSize)).
		Order("articles_source_id asc").
		Find(&crawlers).
		Count(&found).Error
	if err != nil {
		return crawlers, found, err
	}
	return crawlers, found, nil
}

func (repo *CrawlerRepo) UpdateSchedule(id uint, newSchedule string) error {
	err := repo.DB.Model(&entities.Crawler{}).
		Where("id = ?", id).
		Update("schedule", newSchedule).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CrawlerRepo) Update(crawler entities.Crawler) error {
	err := repo.DB.Model(&crawler).
		Updates(entities.Crawler{
				FeedLink: crawler.FeedLink,
				CrawlType: crawler.CrawlType,
				ArticleDiv: crawler.ArticleDiv,
				ArticleTitle: crawler.ArticleTitle,
				ArticleDescription: crawler.ArticleDescription,
				ArticleLink: crawler.ArticleLink,
				ArticleAuthors: crawler.ArticleAuthors,
			}).Error
	if err != nil {
		return err
	}
	return nil
}
