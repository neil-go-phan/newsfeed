package entities

import (
	"time"

	"gorm.io/gorm"
)

type Cronjob struct {
	gorm.Model
	Name             string
	StartedAt          time.Time
	EndedAt            time.Time
	NewArticlesCount int32
	CrawlerID        uint    `gorm:"foreignKey:CrawlerID"`
	Crawler          Crawler `gorm:"foreignKey:CrawlerID"`
}
