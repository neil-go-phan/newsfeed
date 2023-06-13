package entities

import "gorm.io/gorm"

type ArticlesSource struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	FeedLink    string `json:"feed_link"`
	Image       string `json:"image"` // link or a base64 image
}
