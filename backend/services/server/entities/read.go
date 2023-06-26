package entities

type Read struct {
	Username         string `json:"username" gorm:"foreignKey:Username"`
	ArticleID        uint   `json:"article_id" gorm:"foreignKey:ArticleID"`
	ArticlesSourceID uint   `json:"articles_source_id" gorm:"foreignKey:ArticlesSourceID"`
}
