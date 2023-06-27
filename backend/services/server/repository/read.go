package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type ReadRepository interface {
	CreateIfNotExist(Follow entities.Follow) error
}

type ReadRepo struct {
	DB *gorm.DB
}

func NewRead(db *gorm.DB) *ReadRepo {
	return &ReadRepo{
		DB: db,
	}
}

