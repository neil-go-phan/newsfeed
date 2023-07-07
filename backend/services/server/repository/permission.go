package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	List() ([]entities.Permission, error)
}

type PermissionRepo struct {
	DB *gorm.DB
}

func NewPermission(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{
		DB: db,
	}
}

func (repo *PermissionRepo)	List() ([]entities.Permission, error) {
	permissions := make([]entities.Permission, 0)
	err := repo.DB.Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
