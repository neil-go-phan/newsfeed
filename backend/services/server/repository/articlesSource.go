package repository

import (
	"fmt"
	"server/entities"
	"server/helpers"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ArticlesSourcesRepository interface {
	GetWithTopicPaginate(topicID uint, page int, pageSize int) ([]entities.ArticlesSource, int64, error)

	SearchByTitleAndDescriptionPaginate(keyword string, page int, pageSize int) ([]entities.ArticlesSource, int64, error)

	CreateIfNotExist(articlesSource entities.ArticlesSource) (entities.ArticlesSource, error)
	UpdateTopicOneSource(articlesSource entities.ArticlesSource, newTopicId uint) error
	UpdateTopicAllSource(oldTopicId uint, newTopicId uint) error
	IncreaseFollowByOne(articlesSource entities.ArticlesSource) error
	DecreaseFollowByOne(articlesSource entities.ArticlesSource) error

	GetMostActiveSources() ([]MostActiveSource, error)
}

type ArticlesSourcesRepo struct {
	DB *gorm.DB
}

type MostActiveSource struct {
	entities.ArticlesSource
	ArticlesPreviousWeek int `json:"articles_previous_week"`
}

const NEWEST_SOURCE_USER_DASHBOARD_DISPLAY = 3
const NEWEST_ARTICLES_EACH_SOURCE_USER_DASHBOARD_DISPLAY = 4

func NewArticlesSourcesRepo(db *gorm.DB) *ArticlesSourcesRepo {
	return &ArticlesSourcesRepo{
		DB: db,
	}
}

func (repo *ArticlesSourcesRepo) CreateIfNotExist(articlesSource entities.ArticlesSource) (entities.ArticlesSource, error) {
	result := repo.DB.
		Where(entities.ArticlesSource{Link: articlesSource.Link}).
		FirstOrCreate(&articlesSource)
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

	err := repo.DB.
		Scopes(helpers.Paginate(page, pageSize)).
		Where("topic_id = ?", topicID).
		Find(&articlesSources).
		Count(&found).Error
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
	searchKeyword := fmt.Sprint("%" + strings.ToLower(keyword) + "%")
	var found int64

	err := repo.DB.
		Scopes(helpers.Paginate(page, pageSize)).
		Where("LOWER(title) LIKE ? or LOWER(description) LIKE ?", searchKeyword, searchKeyword).
		Find(&articlesSources).
		Count(&found).Error
	if err != nil {
		return articlesSources, found, err
	}
	return articlesSources, found, nil
}

func (repo *ArticlesSourcesRepo) IncreaseFollowByOne(articlesSource entities.ArticlesSource) error {
	err := repo.DB.
		Model(&articlesSource).
		Update("follower", gorm.Expr("follower + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ArticlesSourcesRepo) DecreaseFollowByOne(articlesSource entities.ArticlesSource) error {
	err := repo.DB.
		Model(&articlesSource).
		Update("follower", gorm.Expr("follower - ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ArticlesSourcesRepo) GetMostActiveSources() ([]MostActiveSource, error) {
	articlesSource := make([]MostActiveSource, 0)
	today := time.Now()
	aWeekAgo := today.AddDate(0, 0, -7)
	todayString := today.Format("2006-01-02")
	aWeekAgoString := aWeekAgo.Format("2006-01-02")

	PAGE := 1
	PAGE_SIZE := 5

	subQuery := repo.DB.
		Model(&entities.Article{}).
		Select("articles_source_id", "count(id) as articles_previous_week").
		Where("created_at BETWEEN ? AND ?", aWeekAgoString+" 00:00:00", todayString+" 23:59:59").
		Scopes(helpers.Paginate(PAGE, PAGE_SIZE)).
		Group("articles_source_id")

	err := repo.DB.
		Model(&entities.ArticlesSource{}).
		Distinct("id", "title", "description", "link", "follower", "image", "articles_previous_week", "feed_link").
		Joins("JOIN (?) q on q.articles_source_id = articles_sources.id", subQuery).
		Order("articles_previous_week desc").
		Find(&articlesSource).Error
	if err != nil {
		return articlesSource, err
	}
	return articlesSource, nil
}