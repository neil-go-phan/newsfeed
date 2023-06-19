package entities

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name       string `json:"name" validate:"required"`
	CategoryID uint   `json:"category_id" gorm:"foreignKey:CategoryID" validate:"required"`
}
