package services

import "crawler/entities"

type ArticleServices interface {
	CreateIfNotExist(article *entities.Article) error
}

type ArticlesSourceServices interface {
	Create(articlesSource *entities.ArticlesSource) error
}

type CrawlerServices interface {
	TestRSSCrawler(crawler entities.Crawler) (*entities.ArticlesSource, []*entities.Article, error)
	TestCustomCrawler(crawler entities.Crawler) (entities.ArticlesSource, []*entities.Article, error)
	FirstCrawl(crawler entities.Crawler) (error)
	ScheduledCrawl(crawlerID uint) (error)
}
