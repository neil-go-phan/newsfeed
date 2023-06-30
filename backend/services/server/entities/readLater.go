package entities

import "time"

type ReadLater struct {
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username" gorm:"foreignKey:Username"`
	ArticleID uint      `json:"article_id" gorm:"foreignKey:ArticleID"`
}
