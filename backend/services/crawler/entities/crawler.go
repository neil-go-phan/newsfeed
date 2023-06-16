package entities

import "gorm.io/gorm"

type Crawler struct {
	gorm.Model
	SourceLink         string `json:"source_link"`
	FeedLink           string `json:"feed_link"`
	CrawlType          string `json:"crawl_type"` // two type: "RSS link", "custom crawler"
	ArticleDiv         string `json:"article_div"`
	ArticleTitle       string `json:"article_title"`
	ArticleDescription string `json:"article_description"`
	ArticleLink        string `json:"article_link"`
	ArticlePublished   string `json:"article_published"`
	ArticleAuthors     string `json:"article_authors"`
	Schedule           string `json:"schedule"`
	ArticlesSourceID   uint   `json:"articles_source_id" gorm:"foreignKey:ArticlesSourceID"`
}