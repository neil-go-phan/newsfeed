package entities

type ReadLater struct {
	Username         string `json:"username" gorm:"foreignKey:Username"`
	ArticleID        uint   `json:"article_id" gorm:"foreignKey:ArticleID"`
}
