package entities

import "time"

type ReadLater struct {
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username" gorm:"foreignKey:Username;constraint:OnDelete:CASCADE"`
	ArticleID uint      `json:"article_id" gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE"`
}
