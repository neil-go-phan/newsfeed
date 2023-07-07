package entities

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Entity      string `json:"entity"`
	Method      string `json:"method"`
	Description string `json:"description"`
}
