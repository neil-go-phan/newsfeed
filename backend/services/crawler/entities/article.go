package entities

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Link             string    `json:"link"`
	Published        time.Time `json:"published"`
	Authors          string    `json:"authors"`
	ArticlesSourceID uint      `gorm:"foreignKey:ArticlesSourceID"`
}

