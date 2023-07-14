package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

//go:generate mockery --name FcmNotificationRepository
type FcmNotificationRepository interface {
	Create(FcmNotification entities.FcmNotification) error
	List() ([]entities.FcmNotification, error)
}

type FcmNotificationRepo struct {
	DB *gorm.DB
}

func NewFcmNotification(db *gorm.DB) *FcmNotificationRepo {
	return &FcmNotificationRepo{
		DB: db,
	}
}

func (repo *FcmNotificationRepo) Create(FcmNotification entities.FcmNotification) error {
	err := repo.DB.Create(FcmNotification).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *FcmNotificationRepo) List() ([]entities.FcmNotification, error) {
	FcmNotifications := make([]entities.FcmNotification, 0)
	err := repo.DB.
		Find(&FcmNotifications).Error
	if err != nil {
		return FcmNotifications, err
	}
	return FcmNotifications, nil
}
