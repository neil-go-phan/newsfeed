package repository

import (
	"fmt"
	"server/entities"

	"gorm.io/gorm"
)

type FollowRepository interface {
	GetByUsername(username string) ([]entities.Follow, error)

	CreateIfNotExist(Follow entities.Follow) error
	Delete(Follow entities.Follow) error
}

type FollowRepo struct {
	DB *gorm.DB
}

func NewFollow(db *gorm.DB) *FollowRepo {
	return &FollowRepo{
		DB: db,
	}
}

func (repo *FollowRepo) CreateIfNotExist(follow entities.Follow) error {
	result := repo.DB.
		Where(entities.Follow{Username: follow.Username, ArticlesSourceID: follow.ArticlesSourceID}).
		FirstOrCreate(&follow)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user already follow this articles source")
	}

	return nil
}

func (repo *FollowRepo) Delete(follow entities.Follow) error {
	err := repo.DB.
		Where(entities.Follow{Username: follow.Username, ArticlesSourceID: follow.ArticlesSourceID}).
		Delete(&follow).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *FollowRepo) GetByUsername(username string) ([]entities.Follow, error) {
	follows := make([]entities.Follow, 0)
	err := repo.DB.
		Where(&entities.Follow{Username: username}).
		Preload("ArticlesSource").
		Find(&follows).
		Error
	if err != nil {
		return follows, err
	}
	return follows, nil
}
