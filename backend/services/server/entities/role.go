package entities

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name         string `gorm:"uniqueIndex" json:"name"` 
	Description  string `json:"description"`
	Permissions []*Permission `gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE;" json:"permissions"`
}