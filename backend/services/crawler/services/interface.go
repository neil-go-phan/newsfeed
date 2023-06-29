package services

import "crawler/entities"

type ArticleServices interface {
	CreateIfNotExist(article entities.Article) error
	StoreArticles(articles []entities.Article, articlesSourceID uint) (count int32)
}

type ArticlesSourceServices interface {
	Create(articlesSource *entities.ArticlesSource) error
}

type CrawlerServices interface {
	TestRSSCrawler(crawler entities.Crawler) (entities.ArticlesSource, []entities.Article, error)
	TestCustomCrawler(crawler entities.Crawler) (entities.ArticlesSource, []entities.Article, error)
	Crawl(crawler entities.Crawler) (newArticleCount int32, err error)
}

type FollowServices interface {
	UpdateNewUnreadArticles(sourceID uint, countNewArticleFound int32) error
}
