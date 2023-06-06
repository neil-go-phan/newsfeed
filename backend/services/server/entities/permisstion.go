package entities

import (
	"gorm.io/gorm"
)

type Permisstion struct {
	gorm.Model
	Entity      string
	Method      string
	Description string
	Name        string
	RoleID      uint `gorm:"foreignKey:RoleID"`
}
