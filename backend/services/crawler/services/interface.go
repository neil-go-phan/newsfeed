package services

import "crawler/entities"

type ArticleServices interface {
	CreateIfNotExist(article *entities.Article) error
}

type ArticlesSourceServices interface {
	Create(articlesSource *entities.ArticlesSource) error
}

type CrawlerServices interface {
	TestCrawler(crawler entities.Crawler) (*entities.ArticlesSource, []*entities.Article, error)
	// first crawl (when admin click submit a new crawler - crawl and create a new articleSource)
	FirstCrawl(crawler entities.Crawler) (error)
	ScheduledCrawl(crawlerID uint) (error)
}
