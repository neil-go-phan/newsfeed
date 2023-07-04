package entities

type Follow struct {
	Username         string `json:"username" gorm:"foreignKey:Username"`
	ArticlesSourceID uint   `json:"articles_source_id" gorm:"foreignKey:ArticlesSourceID"`
	Unread           int    `json:"unread"`
	ArticlesSource   ArticlesSource
}
