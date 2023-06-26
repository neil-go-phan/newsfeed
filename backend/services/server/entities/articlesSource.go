package entities

import "gorm.io/gorm"

type ArticlesSource struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Link        string `json:"link" validate:"required,url"`
	FeedLink    string `json:"feed_link" validate:"required,url"`
	Image       string `json:"image" validate:"required"` // base64 image
	Follower    int    `json:"follower"`
	TopicID     uint   `json:"topic_id" gorm:"foreignKey:TopicID" validate:"required"`
}
