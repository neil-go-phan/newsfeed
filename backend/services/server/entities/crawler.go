package entities

import "gorm.io/gorm"

type Crawler struct {
	gorm.Model
	SourceLink         string `json:"source_link" validate:"required,url"`
	FeedLink           string `json:"feed_link"`
	CrawlType          string `json:"crawl_type" validate:"required"`
	ArticleDiv         string `json:"article_div"`
	ArticleTitle       string `json:"article_title"`
	ArticleDescription string `json:"article_description"`
	ArticleLink        string `json:"article_link"`
	ArticlePublished   string `json:"article_published"`
	ArticleAuthors     string `json:"article_authors"`
	Schedule           string `json:"schedule"`
}
