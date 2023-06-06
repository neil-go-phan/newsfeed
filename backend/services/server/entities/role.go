package entities

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name         string `gorm:"uniqueIndex"`
	Description  string
	Permisstions []Permisstion `gorm:"foreignKey:RoleID"`
}
