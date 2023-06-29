package repository

import (
	"server/entities"
	"time"

	"gorm.io/gorm"
)

type ReadRepository interface {
	SelectByUsernameOnDay(username string, day time.Time) ([]entities.Read, error)
	SelectByUsernameAndSourceIDOnDay(username string, articlesSourceID uint, day time.Time) ([]entities.Read, error)
	CountByUsernameAndSourceID(read entities.Read) (int64, error)

	MarkAllAsReadBySourceID(username string, articlesSourceID uint) error
	MarkAllAsReadByUserFollowedSource(username string) error 
	Create(read entities.Read) error
	Delete(read entities.Read) error
}

type ReadRepo struct {
	DB *gorm.DB
}

func NewRead(db *gorm.DB) *ReadRepo {
	return &ReadRepo{
		DB: db,
	}
}

func (repo *ReadRepo) Create(read entities.Read) error {
	err := repo.DB.Create(&read).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ReadRepo) Delete(read entities.Read) error {
	err := repo.DB.
		Where(entities.Read{Username: read.Username, ArticlesSourceID: read.ArticlesSourceID, ArticleID: read.ArticleID}).
		Delete(&read).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ReadRepo) CountByUsernameAndSourceID(read entities.Read) (int64, error) {
	var count int64
	err := repo.DB.
		Where(entities.Read{Username: read.Username, ArticlesSourceID: read.ArticlesSourceID}).
		Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (repo *ReadRepo) SelectByUsernameOnDay(username string, day time.Time) ([]entities.Read, error) {
	reads := make([]entities.Read, 0)
	dayString := day.Format("2006-01-02")
	err := repo.DB.
		Where("username = ? AND created_at BETWEEN ? AND ?", username, dayString+" 00:00:00", dayString+" 23:59:59").
		Find(&reads).Error
	if err != nil {
		return reads, err
	}
	return reads, nil
}

func (repo *ReadRepo) SelectByUsernameAndSourceIDOnDay(username string, articlesSourceID uint, day time.Time) ([]entities.Read, error) {
	reads := make([]entities.Read, 0)
	dayString := day.Format("2006-01-02")
	err := repo.DB.
		Where("articles_source_id = ? username = ? AND created_at BETWEEN ? AND ?", articlesSourceID, username, dayString+" 00:00:00", dayString+" 23:59:59").
		Find(&reads).Error
	if err != nil {
		return reads, err
	}
	return reads, nil
}

func (repo *ReadRepo) MarkAllAsReadBySourceID(username string, articlesSourceID uint) error {
	query := `INSERT INTO reads (username, article_id, articles_source_id)
	SELECT ?, id, articles_source_id
	FROM articles
	WHERE NOT EXISTS (
		SELECT 1
		FROM reads
		WHERE reads.article_id = articles.id
		AND reads.username =  ?
		AND articles_source_id = ?
	) AND articles_source_id = ?;`
	err := repo.DB.Exec(query, username, username, articlesSourceID, articlesSourceID).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ReadRepo) MarkAllAsReadByUserFollowedSource(username string) error {
	query := `INSERT INTO reads (username, article_id, articles_source_id)
		SELECT ?, id, articles.articles_source_id
		FROM articles JOIN (
			SELECT articles_source_id 
			FROM "follows" 
			WHERE username = ?) 
		q on q.articles_source_id = articles.articles_source_id 
		WHERE NOT EXISTS (
			SELECT 1
			FROM reads
			WHERE reads.article_id = articles.id
			AND reads.username =  ?
	);`
	err := repo.DB.Exec(query, username, username, username).Error
	if err != nil {
		return err
	}
	return nil
}


