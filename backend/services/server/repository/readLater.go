package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type ReadLaterRepository interface {
	Create(readLater entities.ReadLater) error
	Delete(readLater entities.ReadLater) error
}

type ReadLaterRepo struct {
	DB *gorm.DB
}

func NewReadLater(db *gorm.DB) *ReadLaterRepo {
	return &ReadLaterRepo{
		DB: db,
	}
}

func (repo *ReadLaterRepo) Create(readLater entities.ReadLater) error {
	err := repo.DB.Create(&readLater).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ReadLaterRepo) Delete(readLater entities.ReadLater) error {
	err := repo.DB.
		Where(entities.ReadLater{Username: readLater.Username, ArticleID: readLater.ArticleID}).
		Delete(&readLater).Error
	if err != nil {
		return err
	}
	return nil
}

