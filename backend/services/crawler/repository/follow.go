package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type FollowRepository interface {
	UpdateUnreadBySourceID(sourceID uint, countNewArticleFound int32) error
}

type FollowRepo struct {
	DB *gorm.DB
}

func NewFollow(db *gorm.DB) *FollowRepo {
	return &FollowRepo{
		DB: db,
	}
}

func (repo *FollowRepo) UpdateUnreadBySourceID(sourceID uint, countNewArticleFound int32) error {
	err := repo.DB.
		Model(entities.Follow{}).Where("articles_source_id = ?", sourceID).
		Update("unread", gorm.Expr("unread + ?", countNewArticleFound)).Error
	if err != nil {
		return err
	}

	return nil
}
