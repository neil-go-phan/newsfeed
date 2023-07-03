package entities

import "time"

type Follow struct {
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Username         string    `json:"username" gorm:"foreignKey:Username"`
	ArticlesSourceID uint      `json:"articles_source_id" gorm:"foreignKey:ArticlesSourceID"`
	Unread           int       `json:"unread"`
	ArticlesSource   ArticlesSource
}
